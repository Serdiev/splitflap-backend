package utils

import "strings"

type TextHelper struct {
	upper     string
	lower     string
	rowLength int
}

func NewText() TextHelper {
	rowLength := 12
	return TextHelper{
		upper:     strings.Repeat(" ", rowLength),
		lower:     strings.Repeat(" ", rowLength),
		rowLength: rowLength,
	}
}

func (t *TextHelper) GetText() string {
	return t.upper + t.lower
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
	if textLength > t.rowLength {
		textLength = 12
	}

	newText := []byte(currentText)
	for i := 0; i < textLength; i++ {
		newText[i] = text[i]
	}

	return string(newText)
}

func (t *TextHelper) EndWith(currentText string, text string) string {
	textLength := len(text)
	if textLength > t.rowLength {
		textLength = 12
	}

	newText := []byte(currentText)
	for i := textLength - 1; i >= 0; i-- {
		newText[t.rowLength-textLength+i] = text[i]
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
