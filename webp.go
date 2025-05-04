package main

import (
	"image"
	"os"

	"github.com/gen2brain/webp"
)

func exportWebP(img image.Image, file string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	return webp.Encode(f, img, webp.Options{
		Quality: Quality,
		Method:  6,
		Exact:   true,
	})
}
