package usb_serial

import (
	"errors"
	"math/rand"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/sp"
	"splitflap-backend/internal/utils"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

var cfg = config.New()

const (
	ForceMovementNone         ForceMovement = "none"
	ForceMovementOnlyNonBlank ForceMovement = "only_non_blank"
	ForceMovementAll          ForceMovement = "all"
	RetryTime                               = time.Millisecond * 5000
)

type ForceMovement string

type EnqueuedMessage struct {
	nonce uint32
	bytes []byte // bytes with CRC32 + null ending
}

type Splitflap struct {
	serial          SerialConnection
	outQueue        chan EnqueuedMessage
	ackQueue        chan uint32
	nextNonce       uint32
	run             bool
	lock            sync.Mutex
	messageHandlers map[sp.SplitFlapType]func(interface{})
	currentConfig   *sp.SplitflapConfig
	numModules      int
	alphabet        []rune
}

type forceMovementFunc func(rune) bool

func NewSplitflap(serialInstance SerialConnection) *Splitflap {
	alphabet := []rune{}
	for _, v := range cfg.Splitflap.AlphabetESP32Order {
		alphabet = append(alphabet, v)
	}

	s := &Splitflap{
		serial:          serialInstance,
		outQueue:        make(chan EnqueuedMessage, 100),
		ackQueue:        make(chan uint32, 100),
		nextNonce:       uint32(rand.Intn(256)),
		run:             true,
		messageHandlers: make(map[sp.SplitFlapType]func(interface{})),
		currentConfig:   nil,
		alphabet:        alphabet,
	}

	// TODO: Remove later
	s.initializeModuleList(cfg.Splitflap.ModuleCount)

	return s
}

func (sf *Splitflap) initializeModuleList(moduleCount int) {
	sf.numModules = moduleCount
	sf.currentConfig = &sp.SplitflapConfig{
		Modules: []*sp.SplitflapConfig_ModuleConfig{},
	}
	for i := 0; i < moduleCount; i++ {
		newModule := sp.SplitflapConfig_ModuleConfig{
			TargetFlapIndex: 0,
			MovementNonce:   0,
			ResetNonce:      0,
		}

		sf.currentConfig.Modules = append(sf.currentConfig.Modules, &newModule)
	}
}

func (sf *Splitflap) readLoop() {
	log.Info().Msg("Read loop started")
	buffer := []byte{}
	// next := time.Now()
	for {
		if !sf.run {
			return
		}

		// TODO: remove
		// if time.Now().Before(next) {
		// 	continue
		// }
		// next = time.Now().Add(time.Second * 1)

		newBytes, err := sf.serial.Read()
		if err != nil {
			log.Info().Msgf("Error reading from serial: %v\n", err)
			return
		}

		if len(newBytes) == 0 {
			continue
		}

		// log.Info().Interface("bytes", string(newBytes)).Msg("Got some data!")
		buffer = append(buffer, newBytes...)

		lastByte := buffer[len(buffer)-1]
		if lastByte != 0 {
			// no data or data stream not ready
			// log.Info().Msg("not end of message yet")
			continue
		}

		// log.Info().Msg("got buffer values")
		sf.processFrame(buffer[:len(buffer)-1])
		buffer = []byte{}
	}
}

func (sf *Splitflap) processFrame(decoded []byte) {
	payload, validCrc := utils.ParseCRC32EncodedPayload(decoded)
	if !validCrc {
		return
	}

	// log.Info().Msg("Got valid crc32")
	message := &sp.FromSplitflap{}

	if err := proto.Unmarshal(payload, message); err != nil {
		log.Info().Msgf("Failed to unmarshal message: %v\n", err)
		return
	}
	// log.Info().Msgf("Received message %v\n", message)
	message.PrintSplitflapState()

	switch message.GetPayload().(type) {
	case *sp.FromSplitflap_Ack:
		nonce := message.GetAck().GetNonce()
		sf.ackQueue <- nonce
	case *sp.FromSplitflap_SplitflapState:
		numModulesReported := len(message.GetSplitflapState().GetModules())

		if sf.numModules == 0 {
			sf.initializeModuleList(numModulesReported)
		} else if sf.numModules != numModulesReported {
			log.Info().Msgf("Number of reported modules changed (was %d, now %d)\n", sf.numModules, numModulesReported)
		}
	}

	// sf.lock.Lock()
	// defer sf.lock.Unlock()

	// handler := sf.messageHandlers[message.GetPayloadType()]
	// handler(message.GetPayload())
}

func (sf *Splitflap) waitingForOutgoingMessage() bool {
	return len(sf.outQueue) == 0
}

func (sf *Splitflap) waitingForIncomingMessage() bool {
	return len(sf.ackQueue) == 0
}

func (sf *Splitflap) writeLoop() {
	log.Info().Msg("Write loop started")

	for {
		if !sf.run {
			log.Info().Msg("Stop running, exiting write loop")
			return
		}

		if sf.waitingForOutgoingMessage() {
			// log.Info().Msg("waiting for outgoing messages")
			continue
		}

		enqueuedMessage := <-sf.outQueue

		nextRetry := time.Now()
		writeCount := 0
		for {
			if !sf.run {
				log.Info().Msg("Stop running, exiting write loop")
				return
			}

			if time.Now().After(nextRetry) {
				if writeCount > 0 {
					log.Info().Msg("Write message again")
				}
				writeCount++
				// log.Info().Msgf("Writing message %x", enqueuedMessage.bytes)
				sf.serial.Write(enqueuedMessage.bytes)
				nextRetry = time.Now().Add(RetryTime)
			}

			if sf.waitingForIncomingMessage() {
				continue
			}

			latestAckNonce := <-sf.ackQueue
			// log.Info().Msgf("nonce: %d", latestAckNonce)
			if enqueuedMessage.nonce == latestAckNonce {
				// log.Info().Msg("Correct nonce, breaking")
				break
			}

			// log.Info().Msgf("Got unexpected nonce: %d\n", latestAckNonce)
		}
	}
}

func (sf *Splitflap) Calibrate() {
	sf.serial.Write([]byte("@"))
}

func (sf *Splitflap) SetText(text string) error {
	return sf.setTextWithMovement(text, ForceMovementNone)
}

func (sf *Splitflap) setTextWithMovement(text string, forceMovement ForceMovement) error {
	// Transform text to a list of flap indexes (and pad with blanks so that all modules get updated even if text is shorter)
	var positions []uint32
	for _, c := range text {
		idx := sf.alphabetIndex(c)
		positions = append(positions, idx)
	}

	// Pad with blanks if text is shorter than the number of modules
	for i := len(text); i < sf.numModules; i++ {
		positions = append(positions, sf.alphabetIndex(' '))
	}

	var forceMovementList []bool
	switch forceMovement {
	case ForceMovementNone:
		forceMovementList = nil
	case ForceMovementOnlyNonBlank:
		for _, c := range text {
			forceMovementList = append(forceMovementList, sf.alphabetIndex(c) != 0)
		}
		// Pad with false if text is shorter than the number of modules
		for i := len(text); i < sf.numModules; i++ {
			forceMovementList = append(forceMovementList, false)
		}
	case ForceMovementAll:
		forceMovementList = make([]bool, sf.numModules)
		for i := range forceMovementList {
			forceMovementList[i] = true
		}
	default:
		panic("Bad movement value")
	}

	return sf.setPositions(positions, forceMovementList)
}

func (sf *Splitflap) setPositions(positions []uint32, forceMovementList []bool) error {
	sf.lock.Lock()
	defer sf.lock.Unlock()

	if sf.numModules == 0 {
		return errors.New("Cannot set positions before the number of modules is known")
	}

	if len(positions) > sf.numModules {
		return errors.New("More positions specified than modules")
	}

	if forceMovementList != nil && len(positions) != len(forceMovementList) {
		return errors.New("positions and forceMovementList length must match")
	}

	for i := 0; i < len(positions); i++ {
		if positions[i] != 0 {
			sf.currentConfig.Modules[i].TargetFlapIndex = positions[i]
		}
		if forceMovementList != nil && forceMovementList[i] {
			sf.currentConfig.Modules[i].MovementNonce = (sf.currentConfig.Modules[i].MovementNonce + 1) % 256
		}
	}

	message := &sp.ToSplitflap{
		Payload: &sp.ToSplitflap_SplitflapConfig{
			SplitflapConfig: sf.currentConfig,
		},
	}

	sf.enqueueMessage(message)
	return nil
}

func (sf *Splitflap) enqueueMessage(message *sp.ToSplitflap) {
	message.Nonce = sf.nextNonce
	sf.nextNonce++

	payload, err := proto.Marshal(message)
	if err != nil {
		log.Error().Msgf("Error marshaling message: %v\n", err)
		return
	}

	newMessage := EnqueuedMessage{
		nonce: message.Nonce,
		bytes: utils.CreatePayloadWithCRC32Checksum(payload),
	}

	sf.outQueue <- newMessage

	approxQLength := len(sf.outQueue)
	// TODO: handle error in some way
	// log.Info().Msgf("Out q length: %d\n", approxQLength)
	if approxQLength > 10 {
		log.Info().Msgf("Output queue length is high! (%d) Is the splitflap still connected and functional?\n", approxQLength)
	}
}

func (sf *Splitflap) requestState() {
	message := sp.ToSplitflap{}
	message.Payload = &sp.ToSplitflap_RequestState{
		RequestState: &sp.RequestState{},
	}

	sf.enqueueMessage(&message)
}

func (sf *Splitflap) alphabetIndex(c rune) uint32 {
	for i, char := range sf.alphabet {
		if char == c {
			return uint32(i)
		}
	}

	return 0 // Default to 0 if character not found in alphabet
}

func (sf *Splitflap) Start() {
	go sf.readLoop()
	go sf.writeLoop()
	sf.requestState()
}

func (sf *Splitflap) Shutdown() {
	log.Info().Msg("Shutting down...")
	sf.run = false
	sf.serial.Close()
	close(sf.outQueue)
	close(sf.ackQueue)
}
