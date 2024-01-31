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

	// utils.SetTimeZone()

	// a := stocks.NewAvanzaClient()
	// r, err := a.GetStockInfo(stocks.TRACKED_STOCKS[1])

	// fmt.Println(r)
	// fmt.Println(r.GetYTD())
	// fmt.Println(err)
	// fmt.Println(time.Now())
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

	time.Sleep(2 * time.Second)
	for i := 1; i < cfg.Splitflap.ModuleCount+1; i++ {
		sf.SetText(utils.MapForSending(strings.Repeat("a", i)))
		time.Sleep(200 * time.Millisecond)
	}

	// sf.SetText(utils.MapForSending(strings.Repeat("%", 24)))
	// time.Sleep(5 * time.Second)
	// sf.SetText(utils.MapForSending(strings.Repeat(",", 24)))
	// time.Sleep(5 * time.Second)
	// sf.SetText(utils.MapForSending(strings.Repeat(":", 24)))
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
