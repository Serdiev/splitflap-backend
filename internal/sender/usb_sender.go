package sender

import (
	"fmt"
	"splitflap-backend/internal/usb_serial"
	"splitflap-backend/internal/utils"

	"github.com/rs/zerolog/log"
)

type UsbSerialSender struct {
	sf          *usb_serial.Splitflap
	CurrentText string
}

func NewUsbSerialSender() *UsbSerialSender {
	connection := usb_serial.NewSerialConnection()
	if connection == nil {
		log.Error().Msg("Could not create USB serial connection")
		return nil
	}

	sf := usb_serial.NewSplitflap(connection)
	sf.Start()

	return &UsbSerialSender{
		sf:          sf,
		CurrentText: "",
	}
}

// SendMessage sends the given text over the serial usb
func (m *UsbSerialSender) SendMessage(text string, sentBy string) error {
	if m.CurrentText == text {
		return nil
	}

	m.CurrentText = text

	fmt.Println("Sent by:", sentBy)
	fmt.Println("upper:", text[0:12])
	fmt.Println("lower:", text[12:])

	mapped := utils.MapForSending(text)
	m.sf.SetText(mapped)
	return nil
}
