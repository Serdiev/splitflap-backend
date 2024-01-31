package sender

type NoopSender struct {
}

func NewNoopSender() *NoopSender {
	return &NoopSender{}
}

// SendMessage sends the given text over the serial usb
func (m *NoopSender) SendMessage(text string, sentBy string) error {
	return nil
}
