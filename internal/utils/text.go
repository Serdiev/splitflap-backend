package utils

import (
	"fmt"
	"splitflap-backend/internal/models"
	"strings"
)

var ROW_LENGTH = 12

type TextHelper struct {
	upper string
	lower string
}

func NewText() TextHelper {
	return TextHelper{
		upper: strings.Repeat(" ", ROW_LENGTH),
		lower: strings.Repeat(" ", ROW_LENGTH),
	}
}

func (t *TextHelper) GetText() string {
	return strings.ToLower(t.upper + t.lower)
}

func (t *TextHelper) TopLeft(text string) {
	t.upper = t.BeginWith(t.upper, text)
}

func (t *TextHelper) BottomLeft(text string) {
	t.lower = t.BeginWith(t.lower, text)
}

func (t *TextHelper) TopRight(text string) {
	t.upper = t.EndWith(t.upper, text)
}

func (t *TextHelper) BottomRight(text string) {
	t.lower = t.EndWith(t.lower, text)
}

func (t *TextHelper) BeginWith(currentText string, text string) string {
	textLength := len(text)
	if textLength > ROW_LENGTH {
		textLength = ROW_LENGTH
	}

	newText := []byte(currentText)
	for i := 0; i < textLength; i++ {
		newText[i] = text[i]
	}

	return string(newText)
}

func (t *TextHelper) EndWith(currentText string, text string) string {
	textLength := len(text)
	if textLength > ROW_LENGTH {
		textLength = ROW_LENGTH
	}

	newText := []byte(currentText)
	for i := textLength - 1; i >= 0; i-- {
		newText[ROW_LENGTH-textLength+i] = text[i]
	}

	return string(newText)
}

func truncate(text string, length int) string {
	diff := len(text) - length

	// truncate to module_count
	if diff > 0 {
		return text[:length]
	}

	return text
}

func ReplaceDisallowedLetters(s string) string {
	output := s
	output = strings.ReplaceAll(output, "å", "a")
	output = strings.ReplaceAll(output, "ä", "a")
	output = strings.ReplaceAll(output, "ö", "0")
	output = strings.ReplaceAll(output, "’", "")
	output = strings.ReplaceAll(output, "'", "")
	output = strings.ReplaceAll(output, ";", ":")
	output = strings.ReplaceAll(output, ".", ",")
	output = strings.ReplaceAll(output, "o", "0")
	output = strings.ReplaceAll(output, "z", "s")
	output = strings.ReplaceAll(output, "é", "e")
	output = strings.ReplaceAll(output, "á", "a")
	output = strings.ReplaceAll(output, "–", "-")
	output = strings.ReplaceAll(output, "ë", "e")
	output = strings.ReplaceAll(output, "#", "")
	output = strings.ReplaceAll(output, "!", "")
	output = strings.ReplaceAll(output, "?", "")
	output = strings.ReplaceAll(output, "(", "")
	output = strings.ReplaceAll(output, ")", "")
	output = strings.ReplaceAll(output, "/", "i")
	output = strings.ReplaceAll(output, "\\", "i")
	return output
}

func getPlayingText(playing *models.SpotifyIsPlaying) string {
	text := NewText()
	if len(playing.Song+" - "+playing.Artist) <= cfg.GetRowLength() {
		text.TopLeft(playing.Song + " - ")
		text.TopRight(playing.Artist)

		sliderText := BottomSlider(playing.PercentageLeft())
		text.BottomLeft(sliderText)
	} else {
		text.TopLeft(playing.Song)
		text.BottomLeft(playing.Artist)
	}

	return text.GetText()
}

func BottomSlider(percentage int) string {
	pct := percentage
	if percentage > 100 {
		pct = 100
	} else if percentage <= 0 {
		pct = 1
	}

	left := float32(cfg.GetRowLength()) * (float32(pct) / 100)
	finishedString := strings.Repeat("-", int(left))
	leftString := strings.Repeat("+", cfg.GetRowLength()-int(left))
	msg := fmt.Sprintf("%s%s", finishedString, leftString)
	return msg
}
