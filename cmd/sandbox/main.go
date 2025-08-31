package main

import (
	"context"
	"fmt"
	"math"
	"splitflap-backend/internal/handlers"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/stocks"

	// "splitflap-backend/internal/stocks"
	"splitflap-backend/internal/utils"
)

func main() {

	fmt.Println("boop")

	// svc := handlers.CreateService(context.Background())

	// img := utils.ConvertUrlToImage("https://fakeimg.pl/64x64")

	// fmt.Println("send img")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Image sent")
	// }
}

func main2() {

	svc := handlers.CreateService(context.Background())
	fmt.Println(svc)

	cl := stocks.NewAvanzaClient()

	msft := stocks.TRACKED_STOCKS[1]
	ohlc, err := cl.GetOHLCData(msft)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ohlc)

	// ohlc.OHLC = []models.OHLC{}

	// fakeData := []int{0, 1, 2, 1, 0, -1, -2, -1, 0, 1, 2, 1, 0}
	// for _, val := range fakeData {
	// 	ohlc.OHLC = append(ohlc.OHLC, models.OHLC{Close: float64(val)})
	// }

	height := 135
	width := 240

	interpolated := InterpolateOHLC(ohlc.OHLC, width)

	scaledValuesInt := make([]int, len(interpolated))
	for i, val := range interpolated {
		scaledValuesInt[i] = int(math.Round(ScaleToDisplay(val, ohlc.GetMinValue(), ohlc.GetMaxValue(), height)))
	}

	fmt.Println(scaledValuesInt)
	img := utils.Image{
		Hash:  "https://www.avanza.se/_api/market-guide/stock/MSFT",
		Image: make([][]utils.Color, height),
	}

	img.InitSize(width, height)

	fmt.Println(img.Image)

	for x, intVal := range scaledValuesInt {
		img.Image[intVal][x] = utils.Color{R: 0, G: 255, B: 0}
	}

	// img := utils.Image{
	// 	Url:   "https://www.avanza.se/_api/market-guide/stock/MSFT",
	// 	Image: [][]utils.Color{},
	// }

	// img.InitSize(240, 135)
	// fmt.Println(img.Image)
	// // color all pixels in the image to green
	// for y := 0; y < len(img.Image); y++ {
	// 	for x := 0; x < len(img.Image[0]); x++ {
	// 		img.Image[y][x] = utils.Color{R: uint8(x), G: 0, B: uint8(y)}
	// 	}
	// }
}

func ScaleToDisplay(value, minSource, maxSource float64, height int) float64 {
	// Normalize to [0, 1]
	normalized := (value - minSource) / (maxSource - minSource)

	// Scale to [0, displayHeight - 1] and invert Y-axis
	scaled := (1 - normalized) * (float64(height - 1))

	return scaled
}

func InterpolateOHLC(data []models.OHLC, targetLength int) []float64 {
	result := make([]float64, targetLength)
	srcLen := len(data)

	for i := 0; i < targetLength; i++ {
		// Map index to source fractional position
		pos := float64(i) * float64(srcLen-1) / float64(targetLength-1)

		// Lower index
		idx := int(math.Floor(pos))
		frac := pos - float64(idx) // Fractional part

		// Linear interpolate
		if idx >= srcLen-1 {
			result[i] = data[srcLen-1].Close
		} else {
			start := data[idx].Close
			end := data[idx+1].Close
			result[i] = start + (end-start)*frac
		}
	}

	return result
}
