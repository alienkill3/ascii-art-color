package src

import (
	"fmt"
	"strings"
)

type BannerOptions struct {
	ColorHex *int
	SubStr   string
	Align    string
}

func bannerBlockWithOptions(inputString string, fontMap map[rune][]string, options BannerOptions, width uint16) string {
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

func PrintBannerWithOptions(inputString string, fontMap map[rune][]string, options BannerOptions, width uint16) {
	fmt.Print(bannerBlockWithOptions(inputString, fontMap, options, width))
	fmt.Println("test")
}

func RenderInputString(input string, fontMap map[rune][]string, options BannerOptions, width uint16) string {
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
		builder.WriteString(bannerBlockWithOptions(p, fontMap, options, width))
	}
	return builder.String()
}

// func RenderInput(input string, fontMap map[rune][]string, options BannerOptions, width uint16) {
// 	fmt.Printf("%*s", width, RenderInputString(input, fontMap, options, width))
// 	fmt.Println("test")
// }
