package serial

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"log"
	"math/rand"
	"os"
	"splitflap-backend/internal/sp"
	"sync"
	"time"

	"github.com/dim13/cobs"
	"github.com/golang/protobuf/proto"
	"github.com/tarm/serial"
)

const splitflapBaud = 230400
const retryTimeout float32 = 0.25

const (
	ForceMovementNone         ForceMovement = "none"
	ForceMovementOnlyNonBlank ForceMovement = "only_non_blank"
	ForceMovementAll          ForceMovement = "all"
)

type ForceMovement string

type EnqueuedMessage struct {
	nonce   uint32
	message []byte
}

type Splitflap struct {
	serial          *serial.Port
	logger          *log.Logger
	outQ            chan EnqueuedMessage
	ackQ            chan uint32
	nextNonce       uint32
	run             bool
	lock            sync.Mutex
	messageHandlers map[sp.SplitFlapType]func(interface{})
	currentConfig   *sp.SplitflapConfig
	numModules      int
	alphabet        []rune
}

type forceMovementFunc func(rune) bool

var defaultAlphabet = []rune{
	' ', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
	'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'.', ',', '\'',
}

func NewSplitflap(serialInstance *serial.Port) *Splitflap {
	s := &Splitflap{
		serial:          serialInstance,
		logger:          log.New(os.Stdout, "splitflap: ", log.LstdFlags),
		outQ:            make(chan EnqueuedMessage),
		ackQ:            make(chan uint32),
		nextNonce:       uint32(rand.Intn(256)),
		run:             true,
		messageHandlers: make(map[sp.SplitFlapType]func(interface{})),
		currentConfig:   nil,
		alphabet:        defaultAlphabet,
		numModules:      6, // TODO: remove
	}

	// temp code
	s.createModules(6)

	return s
}

func (sf *Splitflap) createModules(moduleCount int) {
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
	sf.logger.Println("Read loop started")
	buffer := make([]byte, 0)
	for {
		tmp := make([]byte, 128)
		n, err := sf.serial.Read(tmp)
		if err != nil {
			sf.logger.Printf("Error reading from serial: %v\n", err)
			return
		}
		buffer = append(buffer, tmp[:n]...)

		nullIndex := bytes.IndexByte(buffer, 0)
		for nullIndex != -1 {
			frame := buffer[:nullIndex]
			buffer = buffer[nullIndex+1:]
			sf.processFrame(frame)
			nullIndex = bytes.IndexByte(buffer, 0)
		}

		if !sf.run {
			return
		}
	}
}

func calculateCRC32(data []byte) uint32 {
	crc32hash := crc32.NewIEEE()
	crc32hash.Write(data)
	return crc32hash.Sum32()
}

func (sf *Splitflap) processFrame(frame []byte) {
	decoded := cobs.Decode(frame)
	// if err != nil {
	// 	sf.logger.Printf("Failed decode (%d bytes)\n", len(frame))
	// 	sf.logger.Printf("%s\n", hex.Dump(frame))
	// 	return
	// }

	if len(decoded) < 4 {
		fmt.Println("empty...")
		return
	}

	payload := decoded[:len(decoded)-4]
	expectedCRC := calculateCRC32(payload)
	providedCRC := binary.LittleEndian.Uint32(decoded[len(decoded)-4:])

	if expectedCRC != providedCRC {
		sf.logger.Printf("Bad CRC. expected=%#x, actual=%#x\n", expectedCRC, providedCRC)
		return
	}

	message := &sp.FromSplitflap{}

	if err := proto.Unmarshal(payload, message); err != nil {
		sf.logger.Printf("Failed to unmarshal message: %v\n", err)
		return
	}
	sf.logger.Printf("%v\n", message)

	switch message.GetPayload().(type) {
	case *sp.FromSplitflap_Ack:
		nonce := message.GetAck().GetNonce()
		sf.ackQ <- nonce
	case *sp.FromSplitflap_SplitflapState:
		numModulesReported := len(message.GetSplitflapState().GetModules())
		sf.lock.Lock()
		defer sf.lock.Unlock()

		if sf.numModules == 0 {
			sf.numModules = numModulesReported
			for i := 0; i < numModulesReported; i++ {
				sf.currentConfig.Modules = append(sf.currentConfig.Modules, &sp.SplitflapConfig_ModuleConfig{})
			}
		} else {
			if sf.numModules != numModulesReported {
				sf.logger.Printf("Number of reported modules changed (was %d, now %d)\n", sf.numModules, numModulesReported)
			}
		}
	}

	sf.lock.Lock()
	defer sf.lock.Unlock()

	handler := sf.messageHandlers[message.GetPayloadType()]
	handler(message.GetPayload())
}

func (sf *Splitflap) writeLoop() {
	sf.logger.Println("Write loop started")
	for {
		data := <-sf.outQ
		// Check for shutdown
		if !sf.run {
			sf.logger.Println("Write loop exiting @ _out_q")
			return
		}
		nonce := data.nonce
		encodedMessage := data.message

		nextRetry := time.Now()
		for {
			if time.Now().After(nextRetry) {
				if nextRetry.After(time.Time{}) {
					sf.logger.Println("Retry write...")
				}
				sf.serial.Write(encodedMessage)
				sf.serial.Write([]byte{0})
				nextRetry = time.Now().Add(time.Second)
			}

			select {
			case latestAckNonce := <-sf.ackQ:
				// Check for shutdown
				if !sf.run {
					sf.logger.Println("Write loop exiting @ _ack_q")
					return
				}

				if latestAckNonce == nonce {
					break
				} else {
					sf.logger.Printf("Got unexpected nonce: %d\n", latestAckNonce)
				}
			default:
			}
		}
	}
}

func (sf *Splitflap) SetText(text string, forceMovement ForceMovement) error {
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
		log.Fatalf("bad value %v", forceMovement)
		return errors.New("")
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
	nonce := sf.nextNonce
	sf.nextNonce++

	message.Nonce = 1337 //nonce

	payload, err := proto.Marshal(message)
	if err != nil {
		sf.logger.Printf("Error marshaling message: %v\n", err)
		return
	}

	payload = addCrc32ChecksumToPayload(payload)

	fmt.Println("\x08\xb9\n\x1a\x18\n\x02\x08\x01\n\x02\x08\x02\n\x02\x08\x03\n\x02\x08\x04\n\x02\x08\x05\n\x02\x08\x06\xef\xbcg\xbe")
	fmt.Println("\x08\xb9\n\x1a\x18\n\x02\x08\x01\n\x02\x08\x02\n\x02\x08\x03\n\x02\x08\x04\n\x02\x08\x05\n\x02\x08\x06\xef\xbcg\xbe") // Goal

	newMessage := EnqueuedMessage{
		nonce:   nonce,
		message: payload,
	}

	sf.outQ <- newMessage

	approxQLength := len(sf.outQ)
	// TODO: handle error in some way
	sf.logger.Printf("Out q length: %d\n", approxQLength)
	if approxQLength > 10 {
		sf.logger.Printf("Output queue length is high! (%d) Is the splitflap still connected and functional?\n", approxQLength)
	}
}

func addCrc32ChecksumToPayload(payload []byte) []byte {
	crc32Hash := crc32.NewIEEE()
	crc32Hash.Write(payload)
	checksum := crc32Hash.Sum32()
	byteSlice := make([]byte, 4)
	binary.LittleEndian.PutUint32(byteSlice, checksum)

	for _, b := range byteSlice {
		payload = append(payload, b)
	}
	return payload
}

func (sf *Splitflap) alphabetIndex(c rune) uint32 {
	for i, char := range sf.alphabet {
		if char == c {
			return uint32(i)
		}
	}

	return 0 // Default to 0 if character not found in alphabet
}

func (sf *Splitflap) start() {
	go sf.readLoop()
	go sf.writeLoop()
}

func (sf *Splitflap) shutdown() {
	sf.logger.Println("Shutting down...")
	sf.run = false
	sf.serial.Close()
	close(sf.outQ)
	close(sf.ackQ)
}

// func main() {
// 	port := flag.String("port", "", "Serial port to connect to")
// 	verbose := flag.Bool("verbose", false, "Enable verbose logging")
// 	flag.Parse()

// 	if *port == "" {
// 		fmt.Println("Please provide a serial port using the -port flag.")
// 		os.Exit(1)
// 	}

// 	logLevel := log.Print
// 	if *verbose {
// 		logLevel = log.Println
// 	}
// 	sf := newSplitflap(*port, logLevel)
// 	defer sf.shutdown()

// 	sf.start()

// 	// TODO: Implement the rest of the example logic

// 	sf.wait()
// }
