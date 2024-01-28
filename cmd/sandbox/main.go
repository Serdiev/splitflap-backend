package main

import (
	"fmt"
	"math/rand"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/usb_serial"
	"splitflap-backend/internal/utils"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

var cfg = config.New()

func main() {
	send()
}

func replaceAt(s string, i int, c rune) string {
	r := []rune(s)
	r[i] = c
	return string(r)
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

	// sf.Calibrate()
	// sf.SetText(sender.MapForSending("ready or notfugees      "))
	// // sf.SetText(sender.MapForSending(strings.Repeat("a", 24)))
	// sf.SetText(utils.MapForSending(strings.Repeat("l", 24)))
	// time.Sleep(5 * time.Second)
	// sf.SetText(utils.MapForSending(cfg.Splitflap.AlphabetOffset))
	sf.SetText(utils.MapForSending(strings.Repeat("%", 24)))
	time.Sleep(5 * time.Second)
	sf.SetText(utils.MapForSending(strings.Repeat(",", 24)))
	time.Sleep(5 * time.Second)
	sf.SetText(utils.MapForSending(strings.Repeat(":", 24)))
	// fmt.Println(" mellanslag ")
	// spaces := sender.MapForSending(strings.Repeat(" ", 24))
	// fmt.Println(spaces)
	// spaces := sender.MapForSending(strings.Repeat(" ", 24))
	// fmt.Println(spaces)
	// sf.SetText(spaces) // linux
	// sf.SetText(sender.MapForSending(strings.Repeat(" ", 24))) // windows

	// sf.SetText(sender.MapForSending(manyOf(" ", 24)))
	// // sf.SetText(sender.MapForSending(teaturtle()))
	// // alphabet(sf)
	// alphabetInOrder(sf)

	time.Sleep(30 * time.Second)
}

func alphabet(sf *usb_serial.Splitflap) {
	for i := 0; i < len(utils.IndexToLetterMap); i++ {
		randomVal := rand.Intn(40)
		letter := utils.IndexToLetterMap[randomVal]
		sf.SetText(utils.MapForSending(strings.Repeat(letter, 24)))
		fmt.Println(strings.Repeat(letter, 24))
		time.Sleep(10 * time.Second)
	}
}

func alphabetInOrder(sf *usb_serial.Splitflap) {
	sf.SetText(utils.MapForSending(strings.Repeat(" ", 24)))
	time.Sleep(5 * time.Second)
	for i := 0; i < len(utils.IndexToLetterMap); i++ {
		letter := utils.IndexToLetterMap[i]
		err := sf.SetText(utils.MapForSending(strings.Repeat(letter, 24)))
		if err != nil {
			log.Info().Err(err).Msg("hmm")
		}
		time.Sleep(2 * time.Second)
	}
}
