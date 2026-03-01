package src

import (
	"os"
	"strings"
)

func PrintBanner(inputString, sub string, fontMap map[rune][]string, hexColor int) {
	var builder strings.Builder
	// Write in memory first, then print to stdout at once
	for line := range 8 {
		for _, colorSlice := range CutSubString(inputString, sub) {
			for _, rn := range colorSlice.Str {
				if ascii, exists := fontMap[rn]; exists {
					if colorSlice.Flag {
						builder.WriteString(Colorize(ascii[line], hexColor))
					} else {
						builder.WriteString(ascii[line])
					}
				}
			}
		}
		builder.WriteByte('\n')
	}

	os.Stdout.WriteString(builder.String())
}
