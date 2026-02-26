package src

import (
	"os"
	"strings"
)

func PrintBanner(inputString string, fontMap map[rune][]string) {
	var builder strings.Builder
	// Write in memory first, then print to stdout at once
	for line := 0; line < 8; line++ {
		for _, rn := range inputString {
			if ascii, exists := fontMap[rn]; exists {
				builder.WriteString(ascii[line])
			}
		}
		builder.WriteByte('\n')
	}

	os.Stdout.WriteString(builder.String())
}
