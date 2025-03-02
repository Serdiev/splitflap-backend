package utils

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg" // Support for JPEG
	_ "image/png"  // Support for PNG
	"splitflap-backend/pkg/fluent"
)

type Image struct {
	Image [][]Color `json:"image"`
}

type Color struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
}

// Constructs a byte array [width, height, R, G, B, R, G, B, ...]
func (img *Image) ToBytes() ([]byte, error) {
	ints := []uint8{}

	height := uint8(len(img.Image))
	width := uint8(len(img.Image[0])) // Assume non-empty rows

	ints = append(ints, width)
	ints = append(ints, height)

	for y := uint8(0); y < height; y++ {
		for x := uint8(0); x < width; x++ {
			pixel := img.Image[y][x]
			ints = append(ints, pixel.R)
			ints = append(ints, pixel.G)
			ints = append(ints, pixel.B)
		}
	}

	return ints, nil
}

func NewColor(r, g, b uint8) Color {
	return Color{R: r, G: g, B: b}
}

func ConvertUrlToImage(url string) (*Image, error) {
	var returnImage Image

	err := fluent.
		Get(url).
		OnSuccess(func(payload []byte) error {
			img, _, err := image.Decode(bytes.NewBuffer(payload))
			if err != nil {
				return fmt.Errorf("failed to decode image: %w", err)
			}

			// Get image bounds
			bounds := img.Bounds()
			width, height := bounds.Dx(), bounds.Dy()

			// Convert image pixels to a 2D array of Color
			imageData := make([][]Color, height)
			for y := 0; y < height; y++ {
				imageData[y] = make([]Color, width)
				for x := 0; x < width; x++ {
					r, g, b, _ := img.At(x, y).RGBA()
					imageData[y][x] = Color{
						R: uint8(r >> 8),
						G: uint8(g >> 8),
						B: uint8(b >> 8),
					}
				}
			}

			returnImage = Image{Image: imageData}
			return nil
		}).
		Execute()

	if err != nil {
		return nil, err
	}

	return &returnImage, nil
}
