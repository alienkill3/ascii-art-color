package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ascii-art/src"
)

func main() {
	fontFlag := flag.String("f", "standard", "banner ASCII font")
	flag.Parse()
	inputArgs := flag.Args()

	if !src.IsValidFont(*fontFlag) {
		fmt.Fprintln(os.Stderr, "Error: wrong font flag")
		fmt.Fprintln(os.Stderr, "Valid fonts: standard, shadow, thinkertoy")
		os.Exit(1)
	}

	if len(inputArgs) != 1 {
		fmt.Fprintln(os.Stderr, "Error: wrong arguments")
		fmt.Fprintln(os.Stderr, "Usage: go run . \"INPUT STRING\"")
		os.Exit(1)
	}

	input := inputArgs[0]
	if input == "" {
		return
	}

	if !src.IsValidAscii(input) {
		fmt.Fprintln(os.Stderr, "Error: found non ascii character")
		os.Exit(1)
	}

	fontData := src.GetReader(src.FontType(*fontFlag))
	fontMap, err := src.BuildFontMap(fontData)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing font map")
		os.Exit(1)
	}

	parts := strings.Split(input, "\\n")

	// Ignore trailing "\n" (it should not create an extra empty output block)
	if len(parts) > 0 && parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}

	for _, p := range parts {
		if p == "" {
			fmt.Println()
			continue
		}
		src.PrintBanner(p, fontMap)
	}
}
