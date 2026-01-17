package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {
	in := flag.String("in", "", "Input path to your image file")
	out := flag.String("out", "", "Output path of converted ASCII image")
	noprint := flag.Bool("noprint", false, "Resulting ASCII image does not get printed")
	invert := flag.Bool("invert", false, "Invert ASCII symbols (recommended if using a dark background)")
	width := flag.Int("width", 0, "Modify the original's image size")
	flag.Parse()

	if *in == "" {
		log.Fatalf("provide a path to your image")
	}

	file, err := os.Open(*in)
	if err != nil {
		log.Fatalf("unable to open image: %v\n", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("unable to decode image: %v\n", err)
	}

	img = handleResizing(img, uint(*width))
	ascii := imageToAscii(img, *invert)

	if !*noprint {
		fmt.Printf("%s", ascii)
	}

	if *out != "" {
		err = os.WriteFile(*out, ascii, 0644)
		if err != nil {
			log.Fatalf("unable to write to file: %v\n", err)
		}
		log.Printf("successfully saved ASCII image to %s\n", *out)
	}
}
