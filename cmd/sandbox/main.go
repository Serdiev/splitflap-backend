package main

import (
	"errors"
	"fmt"
	"splitflap-backend/internal/utils"
	"splitflap-backend/pkg/fluent"
)

func test() {
	img := utils.Image{
		Image: [][]utils.Color{
			{{255, 0, 0}, {0, 255, 0}},   // Red, Green
			{{0, 0, 255}, {255, 255, 0}}, // Blue, Yellow
		},
	}

	binaryData, err := img.ToBytes()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Binary Data:", binaryData)
}

func main() {
	img := utils.ConvertUrlToImage("https://i.scdn.co/image/ab67616d000048514e0362c225863f6ae2432651")
	// fmt.Println(img, err)

	// // img = utils.ConvertUrlToImage("https://fakeimg.pl/64/?text=Hello")
	// if err != nil {
	// 	fmt.Println("failed to convert url to image:", err)
	// 	return
	// }

	// img := utils.Image{
	// 	Image: [][]utils.Color{
	// 		{{255, 0, 0}, {0, 255, 0}},   // Red, Green
	// 		{{0, 0, 255}, {255, 255, 0}}, // Blue, Yellow
	// 	},
	// }
	// img := utils.Image{
	// 	Image: [][]utils.Color{
	// 		{{255, 0, 0}, {255, 255, 0}, {177, 177, 177}}, // Red, Green
	// 		{{0, 0, 255}, {6, 71, 163}, {255, 0, 127}},    // Blue, Yellow
	// 		{{0, 0, 255}, {255, 255, 0}, {0, 255, 0}},     // Blue, Yellow
	// 	},
	// }

	// bts, err := json.Marshal(img)
	bts, err := img.ToBytes()
	// bytes, err := json.Marshal(img)
	if err != nil {
		fmt.Println("failed to get image bytes", err)
		return
	}

	// fmt.Println("Binary Data:", bts)
	// fmt.Println("[]byte length:", len(bts))

	url := fmt.Sprintf("http://%s:8080/image", "192.168.1.233")

	// resp, err := http.Post(url, "application/octet-stream", bytes.NewReader(bts))
	// if err != nil {
	// 	fmt.Println("Error sending request:", err)
	// 	return
	// }
	// defer resp.Body.Close()

	fmt.Println("content", fmt.Sprintf("%d", len(bts)))
	err = fluent.Post(url, bts).
		WithHeader("Content-Length", fmt.Sprintf("%d", len(bts))).
		OnSuccess(func(bytes []byte) error {
			return nil
		}).
		OnError(func(bytes []byte) error {
			return errors.New(string(bytes))
		}).
		Execute()

	fmt.Println(err)
}
