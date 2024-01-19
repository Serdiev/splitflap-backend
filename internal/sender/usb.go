package sender

var LetterToIndexMap = createLetterToIndexMap()
var IndexToLetterMap = createIndexToLetterMap()
var wireMapping = createWireMapping()

// 40 letters.
// Maps a letter from my own format to standard format in SplitFlap Arduino code.
var splitflapMapping = map[string]string{
	" ": " ",
	"a": "a",
	"b": "b",
	"c": "c",
	"d": "d",
	"e": "e",
	"f": "f",
	"g": "g",
	"h": "h",
	"i": "i",
	"j": "j",
	"k": "k",
	"l": "l",
	"m": "m",
	"n": "n",
	"p": "o",
	"q": "p",
	"r": "q",
	"s": "r",
	"t": "s",
	"u": "t",
	"v": "u",
	"w": "v",
	"x": "w",
	"y": "x",
	"0": "y",
	"1": "z",
	"2": "0",
	"3": "1",
	"4": "2",
	"5": "3",
	"6": "4",
	"7": "5",
	"8": "6",
	"9": "7",
	"%": "8",
	"+": "9",
	"-": ".",
	",": ",",
	":": "'",
	// extra
	"o": "0",
	"z": "s",
	"å": "a",
	"ä": "a",
	"ö": "o",
}

// Read as from row 0-11 then 12-23 (left to right)
// var flapSpoolOffsets = "::::  "

// key -> driver index
var singleMapping = map[int]int{
	0: 5, 1: 3, 2: 4,
	3: 1, 4: 0, 5: 2,
}

var upper = []int{5, 4, 3}
var lower = []int{1, 0, 2}

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

	inverted := map[int]int{}
	for k, v := range mapping {
		inverted[v] = k
	}

	return inverted
}

func createLetterToIndexMap() map[string]int {
	m := map[string]int{}
	for i, v := range cfg.ALPHABET_CUSTOM_ORDER {
		letter := string(v)
		m[letter] = i
	}
	return m
}

func createIndexToLetterMap() map[string]int {
	m := map[string]int{}
	for i, v := range cfg.ALPHABET_CUSTOM_ORDER {
		letter := string(v)
		m[i] = letter
	}
	return m
}

type SerialSender struct {
}

func NewSerialSender() *SerialSender {
	return &SerialSender{}
}

// SendMessage sends the given text over the serial usb
func (m *SerialSender) SendMessage(text string) error {
	alteredMessage := AdjustWireMapping(text)
	alteredMessage = SpoolOffsetMapping(alteredMessage)

	return nil
}

func MapForSending(s string) string {
	output := SpoolOffsetMapping(s)
	output = AdjustWireMapping(output)
	output = MapToArduinoLetters(output)
	return output
}

// adjust wiring (order on drivers => text in order)
func AdjustWireMapping(text string) string {
	newMsg := ""
	for i := 0; i < len(wireMapping); i++ {
		mapIndex := wireMapping[i]
		newMsg += string(text[mapIndex])
	}

	return newMsg
}

// adjust offset
func SpoolOffsetMapping(text string) string {
	newText := ""

	for i := 0; i < len(text); i++ {
		letterOffset := string(cfg.ALPHABET_OFFSET[i])
		offset := LetterToIndexMap[letterOffset]
		currentIndex := LetterToIndexMap[string(text[i])]
		adjustedIndex := (40 + currentIndex - offset) % 40

		newText += IndexToLetterMap[adjustedIndex]
	}

	return newText
}

// will convert custom letters to the default
func MapToArduinoLetters(text string) string {
	newText := ""
	for i := 0; i < len(text); i++ {
		newText += splitflapMapping[string(text[i])]
	}

	return newText
}
