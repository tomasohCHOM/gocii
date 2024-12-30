package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
	"sync"

	"github.com/nfnt/resize"
)

func main() {
	file, err := os.Open("static/sephiroth.png")
	if err != nil {
		fmt.Printf("Failed to open image: %v\n", err)
		return
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		fmt.Printf("Failed to decode image: %v\n", err)
		return
	}

	// Resize the image
	img = resize.Resize(400, 0, img, resize.Lanczos3)

	// Convert to ASCII
	ascii := convertToASCII(img)
	fmt.Println(ascii)
}

func convertToASCII(img image.Image) string {
	// asciiChars := "@#%XO+=:-. "
	asciiChars := " .-:=+OX%#@"
	lookup := createASCIILookup(asciiChars)
	bounds := img.Bounds()
	height := bounds.Dy()
	width := bounds.Dx()

	result := make([]string, height)
	var wg sync.WaitGroup

	for y := 0; y < height; y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			line := ""
			for x := 0; x < width; x++ {
				color := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
				line += string(lookup[color.Y])
			}
			result[y] = line
		}(y)
	}

	wg.Wait()
	return joinLines(result)
}

func createASCIILookup(asciiChars string) []rune {
	lookup := make([]rune, 256)
	for i := 0; i < 256; i++ {
		index := i * (len(asciiChars) - 1) / 255
		lookup[i] = rune(asciiChars[index])
	}
	return lookup
}

func joinLines(lines []string) string {
	return strings.Join(lines, "\n")
}
