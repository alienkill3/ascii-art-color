package src

import (
	"fmt"
	"strings"
)

type BannerOptions struct {
	ColorHex *int
	SubStr   string
}

func PrintBannerWithOptions(inputString string, fontMap map[rune][]string, options BannerOptions) {
	// No color option: preserve default rendering behavior.
	if options.ColorHex == nil {
		PrintBanner(inputString, fontMap)
		return
	}

	parts := []StringColor{{Str: inputString, Flag: true}}
	if options.SubStr != "" {
		parts = CutSubString(inputString, options.SubStr)
	}

	var builder strings.Builder
	for line := 0; line < 8; line++ {
		for _, part := range parts {
			for _, rn := range part.Str {
				ascii, exists := fontMap[rn]
				if !exists {
					continue
				}

				glyph := ascii[line]
				if part.Flag {
					builder.WriteString(Colorize(glyph, *options.ColorHex))
					continue
				}
				builder.WriteString(glyph)
			}
		}
		builder.WriteByte('\n')
	}

	fmt.Print(builder.String())
}

func RenderInput(input string, fontMap map[rune][]string, options BannerOptions) {
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
		PrintBannerWithOptions(p, fontMap, options)
	}
}
