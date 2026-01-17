package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	file, err := os.Open("static/sephiroth.png")
	if err != nil {
		fmt.Printf("Failed to open image: %v\n", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Failed to decode image: %v\n", err)
		return
	}
	img = resize.Resize(800, 400, img, resize.Lanczos3)

	ascii := ImageToAscii(img, true)
	fmt.Println(string(ascii))
}
