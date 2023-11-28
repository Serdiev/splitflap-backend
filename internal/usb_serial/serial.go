package usb_serial

import (
	"bufio"

	"github.com/rs/zerolog/log"
	"go.bug.st/serial"
)

const DEFAULT_BAUDRATE = 230400
const retryTimeout float32 = 0.25

type SerialConnection interface {
	Open(portName string) error
	Write(data []byte) error
	Read() ([]byte, error)
	Close() error
}

func NewSerialConnection() *Serial {
	list, err := serial.GetPortsList()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get port list")
		return nil
	}

	if len(list) == 0 {
		log.Error().Err(err).Msg("No ports available")
		return nil
	}

	return NewSerialConnectionOnPort(list[0])
}

func NewSerialConnectionOnPort(port string) *Serial {
	s := Serial{}
	err := s.Open(port)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to port")
		return nil
	}
	log.Info().Msgf("Connecting to port %s", port)

	if &s == nil {
		log.Info().Msg("no port?")
	}
	return &s
}

type Serial struct {
	serial *serial.Port
}

func (s *Serial) getSerial() serial.Port {
	return *s.serial
}

func (s *Serial) Open(portName string) error {
	mode := serial.Mode{
		BaudRate: DEFAULT_BAUDRATE,
		DataBits: 8,
	}

	port, err := serial.Open(portName, &mode)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect")
		return err
	}

	s.serial = &port
	return nil
}

func (s *Serial) Write(data []byte) error {
	w, err := s.getSerial().Write(data)
	if err != nil {
		log.Error().Err(err).Msg("failed writing")
	}
	log.Info().Msgf("Bytes written %d", w)
	return err
}

func (s *Serial) Read() ([]byte, error) {
	buffer := []byte{}
	// _, err := s.getSerial().Read(buffer)

	reader := bufio.NewReader(s.getSerial())
	reply, err := reader.ReadBytes('0')
	if err != nil {
		return buffer, err
	}

	return reply, err

	// if err != nil {
	// 	log.Error().Err(err).Msg("failed reading")
	// }
	// return buffer, err
}

func (s *Serial) Close() error {
	return s.getSerial().Close()
}