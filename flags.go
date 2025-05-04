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
	Quality         int
)

func init() {
	arguments.Register("output", 'o', &OutputDirectory).WithHelp("The output directory")
	arguments.Register("input", 'i', &OutputDirectory).WithHelp("The input directory")
	arguments.Register("width", 'w', &MaxWidth).WithHelp("The maximum width the thumbnail should fit in")
	arguments.Register("quality", 'q', &MaxWidth).WithHelp("The quality of the generated the thumbnail")

	arguments.RegisterHelp(true, "Show this help page")

	arguments.Parse()

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

	if MaxWidth <= 0 || MaxWidth > 4096 {
		MaxWidth = 512
	}

	if Quality <= 0 || Quality > 100 {
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
	log.Printf("Quality: ~250~%d%%\n", Quality)
	log.Println()
}

func qualityFromWidth(width uint) int {
	if width <= 512 {
		return 85
	} else if width <= 1024 {
		return 80
	} else if width <= 2048 {
		return 75
	} else if width <= 4096 {
		return 70
	} else {
		return 65
	}
}

func usage() {
	log.Println("Usage: thumb -i <input> -o <output> [-w <width>] [-q <quality>]")
}
