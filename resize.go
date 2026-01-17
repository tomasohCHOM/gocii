package main

import (
	"fmt"
	"image"

	"github.com/nfnt/resize"
)

const HEIGHT_SCALE = 0.5

// Resizes the image given the new desired width. If set to 0, it will
// default to the image's original width and height. It will also modify
// its aspect ratio to be used with mono-like fonts.
func handleResizing(img image.Image, w uint) image.Image {
	bounds := img.Bounds()
	W, H := bounds.Dx(), bounds.Dy()

	if w == 0 {
		w = uint(W)
	}

	aspect := float64(H) / float64(W)
	h := uint(float64(w) * aspect * HEIGHT_SCALE)
	if h == 0 {
		h = 1
	}

	fmt.Println(w, h)

	return resize.Resize(w, h, img, resize.Lanczos3)
}
