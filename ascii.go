package main

import (
	"image"
	"image/color"
)

const DEPTH = 256

func imageToAscii(img image.Image, invert bool) []byte {
	asciiChars := "@#%XO+=:-. "
	if invert {
		asciiChars = " .-:=+OX%#@"
	}
	lookup := asciiLookup(asciiChars)

	bounds := img.Bounds()
	M, N := bounds.Dy(), bounds.Dx()
	result := make([]byte, 0, M*N+N)

	for r := range M {
		for c := range N {
			color := color.GrayModel.Convert(img.At(c, r)).(color.Gray)
			result = append(result, lookup[color.Y])
		}
		if r < M-1 {
			result = append(result, '\n')
		}
	}
	return result
}

func asciiLookup(asciiChars string) []byte {
	lookup := make([]byte, DEPTH)
	for i := range DEPTH {
		idx := i * (len(asciiChars) - 1) / (DEPTH - 1)
		lookup[i] = byte(asciiChars[idx])
	}

	return lookup
}
