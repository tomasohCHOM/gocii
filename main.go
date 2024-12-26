package main

import (
	"fmt"
	"image/png"
	"os"
)

func main() {
	file, err := os.Open("static/sephiroth.png")
	if err != nil {
		fmt.Println("Error opening image:", err)
		return
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}
	bounds := img.Bounds()
	fmt.Println("Width:", bounds.Dx())
	fmt.Println("Height:", bounds.Dy())
}
