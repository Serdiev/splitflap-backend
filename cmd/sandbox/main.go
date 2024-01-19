package main

import (
	"fmt"
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

	sf.SetText(sender.MapForSending(manyOf(":", 24)))
	// alphabet(sf)

	time.Sleep(30 * time.Second)
}
func alphabet(sf *usb_serial.Splitflap) {
	for i := 0; i < len(sender.IndexToLetterMap); i++ {
		letter := sender.IndexToLetterMap[i]
		sf.SetText(sender.MapForSending(manyOf(letter, 24)))
		fmt.Println(manyOf(letter, 24))
		time.Sleep(1 * time.Second)
	}
}

func sendManually() {

	// c := context.Background()
	// app := handlers.CreateService(c)

	// statemachine.HandleStocksState(&app)

	// a, b := ss.GetPortsList()
	// fmt.Println(a, b)

	// req := []byte{4, 8, 100, 34, 5, 47, 210, 156, 12, 0}
	// c := &serial.Config{Name: "COM5", Baud: 230400}
	// s, err := serial.OpenPort(c)

	// if err != nil {
	// 	fmt.Println("err")
	// 	return
	// }

	// t, err := s.Write(req)
	// if err != nil {
	// 	fmt.Println("err")
	// 	return
	// }
	// fmt.Println(t)

	// go func() {
	// 	for {
	// 		buffer := []byte{}
	// 		s.Read(buffer)
	// 		if len(buffer) > 0 {
	// 			fmt.Println("got something")
	// 		}
	// 	}
	// }()

	// time.Sleep(30 * time.Second)

	// return
}
