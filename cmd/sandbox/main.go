package main

import (
	"fmt"
	"splitflap-backend/internal/usb_serial"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {

	// connection := usb_serial.NewSerialConnection()
	connection := usb_serial.NewMockConnection()
	if connection == nil {
		log.Info().Msg("no connection")
		return
	}

	sf := usb_serial.NewSplitflap(connection)
	sf.Start()
	time.AfterFunc(time.Second, func() {
		sf.SetText("d")
	})
	defer sf.Shutdown()

	time.Sleep(30 * time.Second)
	fmt.Println("main finished")
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
