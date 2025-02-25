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
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

func NewColor(r, g, b int) Color {
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
						R: int(r >> 8), // Convert from 16-bit to 8-bit
						G: int(g >> 8),
						B: int(b >> 8),
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
