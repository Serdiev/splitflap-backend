package usb_serial

import (
	"bufio"
	"splitflap-backend/internal/logger"

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
		logger.Error().Err(err).Msg("Failed to get port list")
		return nil
	}

	if len(list) == 0 {
		logger.Error().Err(err).Msg("No ports available")
		return nil
	}

	return NewSerialConnectionOnPort(list[0])
}

func NewSerialConnectionOnPort(port string) *Serial {
	s := Serial{}
	err := s.Open(port)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to port")
		return nil
	}

	logger.Info().Msgf("Connecting to port %s", port)
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
		return err
	}

	s.serial = &port
	return nil
}

func (s *Serial) Write(data []byte) error {
	_, err := s.getSerial().Write(data)
	if err != nil {
		logger.Error().Err(err).Msg("failed writing")
	}
	// logger.Info().Msgf("Bytes written %d", w)
	return err
}

func (s *Serial) Read() ([]byte, error) {
	buffer := []byte{}
	// _, err := s.getSerial().Read(buffer)

	reader := bufio.NewReader(s.getSerial())
	reply, err := reader.ReadBytes(byte(0))
	if err != nil {
		return buffer, err
	}

	return reply, err

	// if err != nil {
	// 	logger.Error().Err(err).Msg("failed reading")
	// }
	// return buffer, err
}

func (s *Serial) Close() error {
	return s.getSerial().Close()
}
