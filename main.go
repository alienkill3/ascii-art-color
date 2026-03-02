package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ascii-art/src"
)

type hexFlag int

func (h *hexFlag) String() string {
	return fmt.Sprintf("0x%x", int(*h))
}

func (h *hexFlag) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return err
	}
	*h = hexFlag(v)
	return nil
}

func main() {
	fontFlag := flag.String("f", "standard", "banner ASCII font")
	var colorFlag hexFlag = 0xffffff
	flag.Var(&colorFlag, "color", "Color Hex 0xRRGGBB")
	flag.Parse()
	inputArgs := flag.Args()

	if !src.IsValidFont(*fontFlag) {
		fmt.Fprintln(os.Stderr, "Error: wrong font flag")
		fmt.Fprintln(os.Stderr, "Valid fonts: standard, shadow, thinkertoy")
		os.Exit(1)
	}

	if len(inputArgs) < 1 || len(inputArgs) > 2 {
		fmt.Fprintln(os.Stderr, "Error: wrong arguments")
		fmt.Fprintln(os.Stderr, "Usage: go run . [OPTION] [STRING]")
		os.Exit(1)
	}

	var input, subStr string
	if len(inputArgs) == 1 {
		input = inputArgs[0]
	} else if len(inputArgs) == 2 {
		subStr = inputArgs[0]
		input = inputArgs[1]
	}
	if input == "" {
		return
	}

	if !src.IsValidAscii(input) || !src.IsValidAscii(subStr) {
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
		src.PrintBanner(p, subStr, fontMap, int(colorFlag))
	}
}
