package src

import (
	"fmt"
	"strings"
)

type BannerOptions struct {
	ColorHex *int
	SubStr   string
}

func bannerBlockWithOptions(inputString string, fontMap map[rune][]string, options BannerOptions) string {
	parts := []StringColor{{Str: inputString, Flag: true}}
	if options.ColorHex == nil {
		parts[0].Flag = false
	} else if options.SubStr != "" {
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
				if part.Flag && options.ColorHex != nil {
					builder.WriteString(Colorize(glyph, *options.ColorHex))
					continue
				}
				builder.WriteString(glyph)
			}
		}
		builder.WriteByte('\n')
	}
	return builder.String()
}

func PrintBannerWithOptions(inputString string, fontMap map[rune][]string, options BannerOptions) {
	fmt.Print(bannerBlockWithOptions(inputString, fontMap, options))
}

func RenderInputString(input string, fontMap map[rune][]string, options BannerOptions) string {
	parts := strings.Split(input, "\\n")
	var builder strings.Builder

	// Ignore trailing "\n" (it should not create an extra empty output block)
	if len(parts) > 0 && parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}

	for _, p := range parts {
		if p == "" {
			builder.WriteByte('\n')
			continue
		}
		builder.WriteString(bannerBlockWithOptions(p, fontMap, options))
	}
	return builder.String()
}

func RenderInput(input string, fontMap map[rune][]string, options BannerOptions) {
	fmt.Print(RenderInputString(input, fontMap, options))
}
