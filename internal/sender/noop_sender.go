package sender

import "fmt"

type NoopSerialSender struct {
}

func NewNoopSerialSender() *UsbSerialSender {
	return &UsbSerialSender{}
}

// SendMessage sends the given text over the serial usb
func (m *NoopSerialSender) SendMessage(text string) error {
	fmt.Println(text)
	return nil
}
