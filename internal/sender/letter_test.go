package sender_test

import (
	"splitflap-backend/internal/utils"
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_WireMapping(t *testing.T) {
	a := "123456123456123456123456"

	res := utils.AdjustWireMapping(a)
	assert.Equal(t, "#112233445566112233445566#", "#"+res+"#")
}

func Test_Adjustments(t *testing.T) {
	a := "                        "

	res := utils.SpoolOffsetMapping(a)
	assert.Equal(t, "# bb:a: : aaabbr  ,aaaa b#", "#"+res+"#")
}

func Test_MapToArdunio(t *testing.T) {
	a := " abcdefghijklmnpqrstuvwxy0123456789%+-,:"

	res := utils.MapToArduinoLetters(a)
	assert.Equal(t, "# abcdefghijklmnopqrstuvwxyz0123456789.,'#", "#"+res+"#")
}

func Test_MapReplacedLetters(t *testing.T) {
	a := "o0zåäö"

	res := utils.MapToArduinoLetters(a)
	assert.Equal(t, "#00saao#", "#"+res+"#")
}
