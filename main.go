package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/coalaura/logger"
)

var log = logger.New().WithOptions(logger.Options{
	NoLevel:    true,
	NoTime:     true,
	ParseCodes: true,
}).WithNoForeground()

func main() {
	log.Println("Reading input directory...")

	var (
		totalSize      uint64
		totalThumbSize uint64
		images         []string
	)

	err := filepath.Walk(InputDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !isImage(path) {
			return err
		}

		totalSize += uint64(info.Size())

		images = append(images, path)

		return nil
	})

	if err != nil {
		log.Fatalf("Failed to read directory: %s\n", err)

		os.Exit(1)
	}

	log.Printf("Found %d images.\n", len(images))

	if _, err := os.Stat(OutputDirectory); os.IsNotExist(err) {
		os.MkdirAll(OutputDirectory, 0755)

		log.Println("Creating output directory...")
	}

	for i, file := range images {
		log.Printf("Creating thumbnail %d of %d...\r", i+1, len(images))

		name := filepath.Base(file)

		img, err := readImage(file)
		if err != nil {
			log.Warningf("Failed to read image %s: %s\n", name, err)

			continue
		}

		img = thumbnail(img)

		ext := filepath.Ext(name)
		out := strings.Replace(name, ext, ".webp", 1)
		out = filepath.Join(OutputDirectory, out)

		err = exportWebP(img, out)
		if err != nil {
			log.Warningf("Failed to export image %s: %s\n", name, err)

			continue
		}

		stat, err := os.Stat(out)
		if err != nil {
			log.Warningf("Failed to stat image %s: %s\n", name, err)

			continue
		}

		totalThumbSize += uint64(stat.Size())
	}

	savings := totalSize - totalThumbSize
	percentage := float64(savings) * 100.0 / float64(totalSize)

	log.Println()
	log.Println()
	log.Printf("Saved ~120~%s ~r~of ~210~%s ~r~(%.2f%%).\n", formatBytes(totalSize-totalThumbSize), formatBytes(totalSize), percentage)
}

func formatBytes(bytes uint64) string {
	const unit = 1024

	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "kMGTPE"[exp])
}
