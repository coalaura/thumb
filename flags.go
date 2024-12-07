package main

import (
	"os"
	"path/filepath"

	"github.com/coalaura/arguments"
)

var (
	OutputDirectory string
	InputDirectory  string
	MaxWidth        uint
	Quality         float32
)

func init() {
	arguments.Parse()

	if arguments.Bool("h", "help", false) {
		usage()

		os.Exit(0)
	}

	OutputDirectory = arguments.String("o", "output", "")
	InputDirectory = arguments.String("i", "input", "")

	MaxWidth = arguments.UIntN("w", "width", 0, arguments.Options[uint]{
		Min: 4,
		Max: 4096,
	})

	Quality = arguments.FloatN("q", "quality", 0, arguments.Options[float32]{
		Min: 10,
		Max: 100,
	})

	if OutputDirectory == "" || InputDirectory == "" {
		log.Fatal("Missing input or output directory")

		usage()

		os.Exit(1)
	}

	if _, err := os.Stat(InputDirectory); os.IsNotExist(err) {
		log.Fatal("Input directory does not exist")

		usage()

		os.Exit(1)
	}

	if MaxWidth <= 0 {
		MaxWidth = 512
	}

	if Quality <= 0 {
		Quality = qualityFromWidth(MaxWidth)
	}

	if abs, err := filepath.Abs(InputDirectory); err == nil {
		InputDirectory = abs
	}

	if abs, err := filepath.Abs(OutputDirectory); err == nil {
		OutputDirectory = abs
	}

	log.Println("~1~_|_~2~|_ ~3~   ~4~__ ~5~|_")
	log.Println("~1~ |_~2~| |~3~|_|~4~|||~5~|_)")

	log.Printf("Input:   ~250~%s\n", InputDirectory)
	log.Printf("Output:  ~250~%s\n", OutputDirectory)
	log.Printf("Width:   ~250~%dpx\n", MaxWidth)
	log.Printf("Quality: ~250~%.0f%%\n", Quality)
	log.Println()
}

func qualityFromWidth(width uint) float32 {
	if width <= 512 {
		return 85.0
	} else if width <= 1024 {
		return 80.0
	} else if width <= 2048 {
		return 75.0
	} else if width <= 4096 {
		return 70.0
	} else {
		return 65.0
	}
}

func usage() {
	log.Println("Usage: thumb -i <input> -o <output> [-w <width>] [-q <quality>]")
}
