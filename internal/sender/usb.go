package sender

type SerialSender struct {
}

func NewSerialSender() *SerialSender {
	return &SerialSender{}
}

// SendMessage sends the given text over the serial usb
func (m *SerialSender) SendMessage(text string) error {

	return nil
}
