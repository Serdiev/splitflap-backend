package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/stocks"
	"splitflap-backend/internal/usb_serial"
	"splitflap-backend/internal/utils"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

var cfg = config.New()

func main() {

	// Specify the URL you want to make a request to
	url := "https://www.avanza.se/_api/market-guide/stock/3873"

	// Create a custom HTTP client with HTTP/1.1 transport
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30,
			DisableCompression: true, // Disable HTTP/2
		},
	}

	// Make the GET request using the custom client
	response, err := client.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", response.StatusCode)
		return
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body
	fmt.Println("Response Body:")
	fmt.Println(string(body))

	utils.SetTimeZone()

	a := stocks.NewAvanzaClient()
	r, err := a.GetStockInfo(stocks.TRACKED_STOCKS[1])

	fmt.Println(r)
	fmt.Println(err)
	// fmt.Println(time.Now())
	// send()
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
