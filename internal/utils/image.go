package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg" // Support for JPEG
	_ "image/png"  // Support for PNG
	"splitflap-backend/internal/logger"
	"splitflap-backend/pkg/fluent"
)

type Image struct {
	Hash  string    `json:"hash"`
	Image [][]Color `json:"image"`
}

func NewImage(img [][]Color) *Image {
	i := &Image{
		Image: img,
	}

	i.ComputeHash()
	return i
}

func (img *Image) ComputeHash() {
	data, err := json.Marshal(img.Image)
	if err != nil {
		img.Hash = err.Error()
	}

	hash := sha256.Sum256(data)
	img.Hash = hex.EncodeToString(hash[:])[:12]
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

	return NewImage(imageData)
}

// Constructs a byte array [width, height, R, G, B, R, G, B, ...]
// Always outputs a 64x64 image, padded with black if smaller.
// Returns nil if the source image is larger than 64x64.
func (img *Image) ToBytes() ([]byte, error) {
	const targetWidth = 64
	const targetHeight = 64

	// Basic validation
	if len(img.Image) == 0 || len(img.Image[0]) == 0 {
		return nil, fmt.Errorf("image is empty")
	}

	srcHeight := len(img.Image)
	srcWidth := len(img.Image[0])

	// Ensure all rows are the same width
	for y := range img.Image {
		if len(img.Image[y]) != srcWidth {
			return nil, fmt.Errorf("image rows have inconsistent widths")
		}
	}

	// Reject images larger than 64x64
	if srcWidth > targetWidth || srcHeight > targetHeight {
		return nil, fmt.Errorf("image too large: got %dx%d, max is 64x64", srcWidth, srcHeight)
	}

	// Calculate offsets to center the image
	offsetX := (targetWidth - srcWidth) / 2
	offsetY := (targetHeight - srcHeight) / 2

	// Preallocate exact size:
	// 2 bytes for width/height + (64 * 64 * 3) for RGB data
	imgBytes := make([]uint8, 0, 2+(targetWidth*targetHeight*3))

	// Output dimensions are always 64x64
	imgBytes = append(imgBytes, uint8(targetWidth))
	imgBytes = append(imgBytes, uint8(targetHeight))

	// Build 64x64 output
	for y := 0; y < targetHeight; y++ {
		for x := 0; x < targetWidth; x++ {
			srcX := x - offsetX
			srcY := y - offsetY

			// If inside source image bounds, use source pixel
			if srcX >= 0 && srcX < srcWidth && srcY >= 0 && srcY < srcHeight {
				pixel := img.Image[srcY][srcX]
				imgBytes = append(imgBytes, pixel.R, pixel.G, pixel.B)
			} else {
				// Otherwise pad with black
				imgBytes = append(imgBytes, 0, 0, 0)
			}
		}
	}

	return imgBytes, nil
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

			returnImg = NewImage(imageData)
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
