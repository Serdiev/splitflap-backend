package usb_serial

import (
	gen "splitflap-backend/internal/generated"
	"splitflap-backend/internal/logger"
	"splitflap-backend/internal/utils"

	"google.golang.org/protobuf/proto"
)

var (
	mockOutBuffer = []byte{}
	mockInBuffer  = []byte{}
)

type MockConnection struct{}

func NewMockConnection() *MockConnection {
	return &MockConnection{}
}

func (m *MockConnection) GetAvailablePorts() ([]string, error) {
	return []string{"COM1"}, nil
}

func (m *MockConnection) Open(portName string) error {
	return nil
}

// [4,8,100,34,5,47,210,156,12,0]
// Mock Write will automatically ack the message after 5 milliseconds
func (m *MockConnection) Write(data []byte) error {
	mockOutBuffer = append(mockOutBuffer, data...)
	// time.AfterFunc(time.Millisecond*5, func() {
	fakeOKAckMessage(data)
	// })

	return nil
}

func (m *MockConnection) Read() ([]byte, error) {
	buffer := mockInBuffer
	if len(buffer) > 0 {
		logger.Info().Msg("has data")
	}
	mockInBuffer = []byte{}
	return buffer, nil
}

func (m *MockConnection) Close() error {
	return nil
}

func fakeOKAckMessage(bytes []byte) {
	nonce := getNonceFromWriteMsg(bytes[:len(bytes)-1])
	ackMsg := gen.FromSplitflap{
		Payload: &gen.FromSplitflap_Ack{
			Ack: &gen.Ack{
				Nonce: nonce,
			},
		},
	}

	bytes, err := proto.Marshal(&ackMsg)
	if err != nil {
		panic("err marshal")
	}

	mockInBuffer = append(mockInBuffer, utils.CreatePayloadWithCRC32Checksum(bytes)...)
	logger.Info().Msg("Added ack message")
}

func getNonceFromWriteMsg(bytes []byte) uint32 {
	payload, validCRC32 := utils.ParseCRC32EncodedPayload(bytes)
	if !validCRC32 {
		panic("invalid crc32")
	}

	message := &gen.ToSplitflap{}

	if err := proto.Unmarshal(payload, message); err != nil {
		logger.Info().Msgf("Failed to unmarshal message: %v\n", err)
		return 0
	}

	return message.Nonce
}
