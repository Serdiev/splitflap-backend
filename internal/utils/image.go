package utils

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg" // Support for JPEG
	_ "image/png"  // Support for PNG
	"splitflap-backend/internal/logger"
	"splitflap-backend/pkg/fluent"
)

type Image struct {
	Url   string    `json:"url"`
	Image [][]Color `json:"image"`
}

func (i *Image) InitSize(width, height int) {
	i.Image = make([][]Color, height)
	for y := 0; y < height; y++ {
		i.Image[y] = make([]Color, width)
		for x := 0; x < width; x++ {
			i.Image[y][x] = Color{R: 0, G: 0, B: 0}
		}
	}
}

type Color struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
}

func EmptyImage() *Image {
	imageData := make([][]Color, 64)
	for y := 0; y < 64; y++ {
		imageData[y] = make([]Color, 64)
		for x := 0; x < 64; x++ {
			imageData[y][x] = Color{R: 0, G: 0, B: 0}
		}
	}

	return &Image{Image: imageData, Url: "empty..."}
}

// Constructs a byte array [width, height, R, G, B, R, G, B, ...]
func (img *Image) ToBytes() ([]byte, error) {
	ints := []uint8{}

	height := uint8(len(img.Image))
	width := uint8(len(img.Image[0])) // Assume non-empty rows
	fmt.Println("width", width)
	fmt.Println("height", height)

	ints = append(ints, width)
	ints = append(ints, height)

	for y := uint8(0); y < height; y++ {
		for x := uint8(0); x < width; x++ {
			ints = append(ints, img.Image[y][x].R)
			ints = append(ints, img.Image[y][x].G)
			ints = append(ints, img.Image[y][x].B)
		}
	}

	return ints, nil
}

func NewColor(r, g, b uint8) Color {
	return Color{R: r, G: g, B: b}
}

func ConvertUrlToImage(url string) (returnImg *Image) {
	err := fluent.
		Get(url).
		OnSuccess(func(payload []byte) error {
			img, _, iErr := image.Decode(bytes.NewBuffer(payload))
			if iErr != nil {
				return fmt.Errorf("failed to decode image: %w", iErr)
			}

			bounds := img.Bounds()
			width, height := bounds.Dx(), bounds.Dy()

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

			returnImg = &Image{Image: imageData, Url: url}
			return nil
		}).
		OnError(func(payload []byte) error {
			return fmt.Errorf("failed to get image: %s", string(payload))
		}).
		Execute()

	if err != nil {
		logger.Error().Err(err).Msg("failed to get image")
	}

	return
}
