package main

import (
	"fmt"
	"math/rand"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/sender"
	"splitflap-backend/internal/usb_serial"
	"time"

	"github.com/rs/zerolog/log"
)

var cfg = config.New()

func main() {
	// sender.AddMapping(4)
	send()
}

func manyOf(a string, num int) string {
	newString := ""
	for i := 0; i < num; i++ {
		newString += a
	}
	return newString
}

func teaturtle() string {
	return "suck my ass teaturtle   "
}

func send() {

	connection := usb_serial.NewSerialConnection()
	// connection := usb_serial.NewMockConnection()
	if connection == nil {
		log.Info().Msg("no connection")
		return
	}

	sf := usb_serial.NewSplitflap(connection)
	sf.Start()
	defer sf.Shutdown()

	sf.Calibrate()
	// sf.SetText(sender.MapForSending(manyOf(" ", 24)))
	// // alphabet(sf)
	// alphabetInOrder(sf)

	time.Sleep(30 * time.Second)
}

func alphabet(sf *usb_serial.Splitflap) {
	for i := 0; i < len(sender.IndexToLetterMap); i++ {
		randomVal := rand.Intn(40)
		letter := sender.IndexToLetterMap[randomVal]
		sf.SetText(sender.MapForSending(manyOf(letter, 24)))
		fmt.Println(manyOf(letter, 24))
		time.Sleep(10 * time.Second)
	}
}

func alphabetInOrder(sf *usb_serial.Splitflap) {
	time.Sleep(5 * time.Second)
	for i := 0; i < len(sender.IndexToLetterMap); i++ {
		letter := sender.IndexToLetterMap[i]
		sf.SetText(sender.MapForSending(manyOf(letter, 24)))
		fmt.Println(manyOf(letter, 24))
		time.Sleep(1 * time.Second)
	}
}
