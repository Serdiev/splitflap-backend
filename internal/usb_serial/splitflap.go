package usb_serial

import (
	"errors"
	"math/rand"
	config "splitflap-backend/configs"
	gen "splitflap-backend/internal/generated"
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
	RetryTime                               = time.Millisecond * 500
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
	messageHandlers map[gen.SplitFlapType]func(interface{})
	currentConfig   *gen.SplitflapConfig
	numModules      int
	alphabet        []rune
}

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
		messageHandlers: make(map[gen.SplitFlapType]func(interface{})),
		currentConfig:   nil,
		alphabet:        alphabet,
	}

	// TODO: Remove later
	s.initializeModuleList(cfg.Splitflap.ModuleCount)

	return s
}

func (sf *Splitflap) initializeModuleList(moduleCount int) {
	sf.numModules = moduleCount
	sf.currentConfig = &gen.SplitflapConfig{
		Modules: []*gen.SplitflapConfig_ModuleConfig{},
	}
	for i := 0; i < moduleCount; i++ {
		newModule := gen.SplitflapConfig_ModuleConfig{
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
	for {
		if !sf.run {
			return
		}

		newBytes, err := sf.serial.Read()
		if err != nil {
			log.Info().Msgf("Error reading from serial: %v\n", err)
			return
		}

		if len(newBytes) == 0 {
			continue
		}

		buffer = append(buffer, newBytes...)
		lastByte := buffer[len(buffer)-1]
		if lastByte != 0 {
			continue
		}

		sf.processFrame(buffer[:len(buffer)-1])
		buffer = []byte{}
	}
}

func (sf *Splitflap) processFrame(decoded []byte) {
	payload, validCrc := utils.ParseCRC32EncodedPayload(decoded)
	if !validCrc {
		return
	}

	message := &gen.FromSplitflap{}

	if err := proto.Unmarshal(payload, message); err != nil {
		log.Info().Msgf("Failed to unmarshal message: %v\n", err)
		return
	}
	message.PrintSplitflapState()

	switch message.GetPayload().(type) {
	case *gen.FromSplitflap_Ack:
		nonce := message.GetAck().GetNonce()
		sf.ackQueue <- nonce
	case *gen.FromSplitflap_SplitflapState:
		numModulesReported := len(message.GetSplitflapState().GetModules())

		if sf.numModules == 0 {
			sf.initializeModuleList(numModulesReported)
		} else if sf.numModules != numModulesReported {
			log.Info().Msgf("Number of reported modules changed (was %d, now %d)\n", sf.numModules, numModulesReported)
		}
	}
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
				sf.serial.Write(enqueuedMessage.bytes)
				nextRetry = time.Now().Add(RetryTime)
			}

			if sf.waitingForIncomingMessage() {
				continue
			}

			latestAckNonce := <-sf.ackQueue
			if enqueuedMessage.nonce == latestAckNonce {
				break
			}
		}
	}
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
		return errors.New("cannot set positions before the number of modules is known")
	}

	if len(positions) > sf.numModules {
		return errors.New("more positions specified than modules")
	}

	if forceMovementList != nil && len(positions) != len(forceMovementList) {
		return errors.New("positions and forceMovementList length must match")
	}

	for i := 0; i < len(positions); i++ {
		sf.currentConfig.Modules[i].TargetFlapIndex = positions[i]
		if forceMovementList != nil && forceMovementList[i] {
			sf.currentConfig.Modules[i].MovementNonce = (sf.currentConfig.Modules[i].MovementNonce + 1) % 256
		}
	}

	message := &gen.ToSplitflap{
		Payload: &gen.ToSplitflap_SplitflapConfig{
			SplitflapConfig: sf.currentConfig,
		},
	}

	sf.enqueueMessage(message)
	return nil
}

func (sf *Splitflap) enqueueMessage(message *gen.ToSplitflap) {
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
	message := gen.ToSplitflap{}
	message.Payload = &gen.ToSplitflap_RequestState{
		RequestState: &gen.RequestState{},
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
