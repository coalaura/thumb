package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
	_ "golang.org/x/image/webp"
)

func readImage(file string) (image.Image, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	img, _, err := image.Decode(f)

	return img, err
}

func thumbnail(img image.Image) image.Image {
	height := img.Bounds().Max.Y

	return resize.Thumbnail(MaxWidth, uint(height), img, resize.Lanczos3)
}

func isImage(file string) bool {
	ext := strings.ToLower(filepath.Ext(file))

	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".webp"
}
