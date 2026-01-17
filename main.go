package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"

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
	img = resize.Resize(400, 200, img, resize.Lanczos3)

	// Convert to ASCII
	ascii := convertToASCII(img)
	fmt.Println(strings.Index(ascii, "\n"))
	fmt.Println(len(ascii))
	fmt.Println(ascii)
}

func convertToASCII(img image.Image) string {
	asciiChars := " .-:=+OX%#@"
	lookup := createASCIILookup(asciiChars)
	bounds := img.Bounds()
	height := bounds.Dy()
	width := bounds.Dx()

	result := make([]string, height)

	for y := range height {
		var line strings.Builder
		for x := range width {
			color := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			line.WriteString(string(lookup[color.Y]))
		}
		result[y] = line.String()
	}

	return joinLines(result)
}

func createASCIILookup(asciiChars string) []rune {
	lookup := make([]rune, 256)
	for i := range 256 {
		index := i * (len(asciiChars) - 1) / 255
		lookup[i] = rune(asciiChars[index])
	}
	return lookup
}

func joinLines(lines []string) string {
	return strings.Join(lines, "\n")
}
