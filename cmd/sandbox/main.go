package main

import (
	"fmt"
	"splitflap-backend/internal/utils"
)

func main() {
	// img, err := utils.ConvertUrlToImage("https://i.scdn.co/image/ab67616d00004851b11a5489e8cb11dd22b930a0")
	// fmt.Println(img, err)

	fmt.Println(len("ö"))
	t := utils.ReplaceDisallowedLetters("ööööö")
	fmt.Println(t)
}
