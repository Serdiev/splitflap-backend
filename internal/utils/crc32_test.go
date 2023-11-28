package utils_test

import (
	"splitflap-backend/internal/utils"
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_Crc32(t *testing.T) {
	encoded := utils.CreatePayloadWithCRC32Checksum([]byte{1, 2, 3})

	decoded, ok := utils.ParseCRC32EncodedPayload(encoded)

	assert.Equal(t, true, ok)
	assert.Equal(t, []byte{1, 2, 3}, decoded)
}
