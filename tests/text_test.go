package tests

import (
	"splitflap-backend/internal/utils"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

var twelveSpaces = strings.Clone(" ")

func getSpaces(n int) string {
	return strings.Repeat(" ", n)
}

func Test_Text(t *testing.T) {
	text := utils.NewText()

	assert.Equal(t, 24, len(text.GetText()))
}

func Test_UpperLeft(t *testing.T) {
	text := utils.NewText()

	text.TopLeft("hej")
	assert.Equal(t, "hej"+getSpaces(21), text.GetText())
}

func Test_BottomLeft(t *testing.T) {
	text := utils.NewText()

	text.BottomLeft("1234561234561337")
	assert.Equal(t, getSpaces(12)+"123456123456", text.GetText())
}

func Test_UpperRight(t *testing.T) {
	text := utils.NewText()

	text.TopRight("hej")
	assert.Equal(t, getSpaces(9)+"hej"+getSpaces(12), text.GetText())
}

func Test_BottomRight(t *testing.T) {
	text := utils.NewText()

	text.BottomRight("hej")
	assert.Equal(t, getSpaces(21)+"hej", text.GetText())
}

func Test_AllThings(t *testing.T) {
	text := utils.NewText()

	text.TopLeft("Microsoft")
	text.TopRight("+1.24%")
	text.BottomLeft("YTD")
	text.BottomRight("+23.34%")
	assert.Equal(t, "Microsoft +1.24%YTD  +23.34%", text.GetText())

}
