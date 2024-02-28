package main

import (
	"context"
	"fmt"
	"math/rand"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/handlers"
	"splitflap-backend/internal/usb_serial"
	"splitflap-backend/internal/utils"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

var cfg = config.New()

func removeValue(slice []int, valueToRemove int) []int {
	var result []int

	for _, v := range slice {
		if v != valueToRemove {
			result = append(result, v)
		}
	}

	return result
}

func matrixReplacement(textBefore string, newText string) {
	currentText := textBefore
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
			currentText = replaceAt(currentText, randomIndex, rune(newText[randomIndex]))
			indexesLeft = removeValue(indexesLeft, randomIndex)
		}

		fmt.Println(currentText)
	}
}

func main() {
	send2()
	// matrixReplacement(strings.Repeat(" ", cfg.Splitflap.ModuleCount), strings.Repeat("a", cfg.Splitflap.ModuleCount))
}

func replaceAt(s string, i int, c rune) string {
	r := []rune(s)
	r[i] = c
	return string(r)
}

func send2() {

	app := handlers.CreateService(context.Background())

	app.Sender.SendMessage(strings.Repeat(" ", cfg.Splitflap.ModuleCount), "sand")
	time.Sleep(5 * time.Second)
	app.Sender.SendMessage(strings.Repeat("x", cfg.Splitflap.ModuleCount), "sand")
	// for i := 1; i < cfg.Splitflap.ModuleCount+1; i++ {
	// 	sf.SetText(utils.MapForSending(strings.Repeat("a", i)))
	// 	time.Sleep(20 * time.Millisecond)
	// }

	// sf.SetText(utils.MapForSending(strings.Repeat("%", 24)))
	// time.Sleep(5 * time.Second)
	// sf.SetText(utils.MapForSending(strings.Repeat(",", 24)))
	// time.Sleep(5 * time.Second)
	// sf.SetText(utils.MapForSending(strings.Repeat(":", 24)))
	// // alphabet(sf)
	// alphabetInOrder(sf)

	time.Sleep(30 * time.Second)
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

	sf.SetText(utils.MapForSending(strings.Repeat(" ", cfg.Splitflap.ModuleCount)))
	time.Sleep(2 * time.Second)
	sf.SetText(utils.MapForSending(strings.Repeat("x", cfg.Splitflap.ModuleCount)))
	// for i := 1; i < cfg.Splitflap.ModuleCount+1; i++ {
	// 	sf.SetText(utils.MapForSending(strings.Repeat("a", i)))
	// 	time.Sleep(20 * time.Millisecond)
	// }

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
