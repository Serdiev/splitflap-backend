package utils_test

import (
	"splitflap-backend/internal/utils"
	"testing"
)

func TestMapForSending(t *testing.T) {
	expected := "123456123456123456123456"

	output := utils.MapForSending(expected)
	input := utils.MapForReading(output)

	if input != expected {
		t.Errorf("MapForSending(%q) = %q; want %q", input, input, expected)
	}
}
