package utils

import (
	"fmt"
	config "splitflap-backend/configs"
	"strings"
)

var cfg = config.New()

var LetterToIndexMap = createLetterToIndexMap()
var IndexToLetterMap = createIndexToLetterMap()
var InvertedWireMapping = map[int]int{}
var WireMapping = createWireMapping()
var CustomToArduinoMapping = createArduinoMap()

func createArduinoMap() map[string]string {
	m := map[string]string{}
	for i := 0; i < len(cfg.Splitflap.AlphabetCustomOrder); i++ {
		m[string(cfg.Splitflap.AlphabetCustomOrder[i])] = string(cfg.Splitflap.AlphabetESP32Order[i])
	}

	return m
}

// Read as from row 0-11 then 12-23 (left to right)
// var flapSpoolOffsets = "::::  "

// key -> driver index
// var singleMapping = map[int]int{
// 	0: 5, 1: 3, 2: 4,
// 	3: 1, 4: 0, 5: 2,
// }

// var upper = []int{5, 4, 3}
// var lower = []int{1, 0, 2}

// var wireMapping = map[int]int{
// 	0: 22, 1: 21, 2: 19, 3: 16, 4: 15, 5: 13, 6: 10, 7: 9, 8: 7, 9: 4, 10: 3, 11: 1,
// 	12: 23, 13: 20, 14: 18, 15: 17, 16: 14, 17: 12, 18: 11, 19: 8, 20: 6, 21: 5, 22: 2, 23: 0,
// }

// var wireMapping = map[int]int{
// 	0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 10: 10, 11: 11,
// 	12: 12, 13: 13, 14: 14, 15: 15, 16: 16, 17: 17, 18: 18, 19: 19, 20: 20, 21: 21, 22: 22, 23: 23,
// }

func createWireMapping() map[int]int {
	mapping := map[int]int{}
	mapping[23] = 0
	mapping[11] = 1
	mapping[22] = 2
	mapping[10] = 3
	mapping[9] = 4
	mapping[21] = 5

	mapping[20] = 6
	mapping[8] = 7
	mapping[19] = 8
	mapping[7] = 9
	mapping[6] = 10
	mapping[18] = 11

	mapping[17] = 12
	mapping[5] = 13
	mapping[16] = 14
	mapping[4] = 15
	mapping[3] = 16
	mapping[15] = 17

	mapping[14] = 18
	mapping[2] = 19
	mapping[13] = 20
	mapping[1] = 21
	mapping[0] = 22
	mapping[12] = 23

	InvertedWireMapping = mapping

	inverted := map[int]int{}
	for k, v := range mapping {
		inverted[v] = k
	}

	return inverted
}

func createLetterToIndexMap() map[string]int {
	m := map[string]int{}
	for i, v := range cfg.Splitflap.AlphabetCustomOrder {
		letter := string(v)
		m[letter] = i
	}
	return m
}

func createIndexToLetterMap() map[int]string {
	m := map[int]string{}
	for i, v := range cfg.Splitflap.AlphabetCustomOrder {
		letter := string(v)
		m[i] = letter
	}
	return m
}

func MapForSending(s string) string {
	output := s
	output = ReplaceDisallowedLetters(output)
	output = SpoolOffsetMapping(output)
	output = AdjustWireMapping(output)
	output = MapToArduinoLetters(output)
	fmt.Println(output)
	return output
}

func ReplaceDisallowedLetters(s string) string {
	output := s
	output = strings.ReplaceAll(output, "o", "0")
	output = strings.ReplaceAll(output, "z", "s")
	output = strings.ReplaceAll(output, "å", "a")
	output = strings.ReplaceAll(output, "ä", "a")
	output = strings.ReplaceAll(output, "ö", "o")
	return output
}

// adjust wiring (order on drivers => text in order)
func AdjustWireMapping(text string) string {
	output := ""
	for i := 0; i < len(WireMapping); i++ {
		mapIndex := WireMapping[i]
		output += string(text[mapIndex])
	}

	return output
}

// adjust offset
func SpoolOffsetMapping(text string) string {
	output := ""

	for i := 0; i < len(text); i++ {
		letterOffset := string(cfg.Splitflap.AlphabetOffset[i])
		offset := LetterToIndexMap[letterOffset]
		currentIndex := LetterToIndexMap[string(text[i])]
		adjustedIndex := (40 + currentIndex - offset) % 40

		output += IndexToLetterMap[adjustedIndex]
	}

	return output
}

// will convert custom letters to the ESP32 "order"
func MapToArduinoLetters(text string) string {
	output := ""
	for i := 0; i < len(text); i++ {
		output += CustomToArduinoMapping[string(text[i])]
	}

	return output
}
