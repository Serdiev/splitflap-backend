package utils

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"

	"github.com/dim13/cobs"
	"github.com/rs/zerolog/log"
)

// Adds CRC32 checksum + null ending
func CreatePayloadWithCRC32Checksum(payload []byte) []byte {
	payload = binary.LittleEndian.AppendUint32(payload, CalculateCRC32(payload))
	payload = cobs.Encode(payload)
	payload = append(payload, 0)

	return payload
}

func CalculateCRC32(data []byte) uint32 {
	crc32hash := crc32.NewIEEE()
	crc32hash.Write(data)
	var t = crc32hash.Sum32()
	return t
}

func ParseCRC32EncodedPayload(data []byte) ([]byte, bool) {
	decoded := cobs.Decode(data)
	fmt.Println(decoded)
	decoded = decoded[:len(decoded)-1]

	if len(data) < 4 {
		log.Info().Msg("< 4 bytes")
		return []byte{}, false
	}
	payload := decoded[:len(decoded)-4]

	crc32bytes := decoded[len(decoded)-4:]

	expectedCRC := CalculateCRC32(payload)
	providedCRC := binary.LittleEndian.Uint32(crc32bytes)
	if expectedCRC != providedCRC {
		log.Info().Msgf("Bad CRC. expected=%#x, actual=%#x\n", expectedCRC, providedCRC)
		return []byte{}, false
	}

	return payload, true
}