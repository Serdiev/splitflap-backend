package sender

import (
	"fmt"
	"math/rand"
	"splitflap-backend/internal/usb_serial"
	"splitflap-backend/internal/utils"
	"time"

	"github.com/rs/zerolog/log"
)

type UsbSerialSender struct {
	sf              *usb_serial.Splitflap
	CurrentText     string
	CurrentRemapped string
}

func NewUsbSerialSender() *UsbSerialSender {
	connection := usb_serial.NewSerialConnection()
	if connection == nil {
		log.Error().Msg("Could not create USB serial connection")
		return nil
	}

	sf := usb_serial.NewSplitflap(connection)
	sf.Start()

	sender := &UsbSerialSender{
		sf:              sf,
		CurrentText:     cfg.Splitflap.AlphabetOffset,
		CurrentRemapped: utils.MapForSending(cfg.Splitflap.AlphabetOffset),
	}
	fmt.Println(sender)
	return sender
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
	// m.matrixReplacement(mapped)
	m.CurrentRemapped = mapped
	return nil
}

func (m *UsbSerialSender) matrixReplacement(newText string) {
	indexesLeft := []int{}
	for i := 0; i < cfg.Splitflap.ModuleCount; i++ {
		indexesLeft = append(indexesLeft, i)
	}

	lettersPerRun := 4
	for i := 0; i < cfg.Splitflap.ModuleCount/lettersPerRun; i++ {
		fmt.Println("new run:")
		// pick 4
		for i := 0; i < lettersPerRun; i++ {
			randomIndex := indexesLeft[rand.Intn(len(indexesLeft))]
			fmt.Println(randomIndex)
			m.CurrentRemapped = replaceAt(m.CurrentRemapped, randomIndex, rune(newText[randomIndex]))
			indexesLeft = removeValue(indexesLeft, randomIndex)
		}

		fmt.Println("curr", m.CurrentRemapped)
		m.sf.SetText(m.CurrentRemapped)
		time.Sleep(50 * time.Millisecond)
	}
}

func replaceAt(s string, i int, c rune) string {
	r := []rune(s)
	r[i] = c
	return string(r)
}

func removeValue(slice []int, valueToRemove int) []int {
	var result []int

	for _, v := range slice {
		if v != valueToRemove {
			result = append(result, v)
		}
	}

	return result
}
