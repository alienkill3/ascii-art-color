package src

import (
	"fmt"
	"strconv"
	"strings"
)

// Colorize applies RGB color formatting to a string using ANSI escape codes.
// for 24-bit true color terminal output.
// Example: colorized := Colorize("Hello", 0xFF0000) // Red color
func Colorize(str string, colorHex int) string {
	if len(str) == 0 {
		return ""
	}
	r := (colorHex >> 16) & 0xFF
	g := (colorHex >> 8) & 0xFF
	b := colorHex & 0xFF
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, str)
}

// ParseHexColor parses a hex RGB string to integer.
// Accepts both "RRGGBB" and "#RRGGBB".
func ParseHexColor(color string) (int, error) {
	clean := strings.TrimSpace(strings.TrimPrefix(color, "#"))
	if len(clean) != 6 {
		return 0, fmt.Errorf("invalid color format: use RRGGBB or #RRGGBB")
	}

	val, err := strconv.ParseInt(clean, 16, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid hex color: %w", err)
	}

	return int(val), nil
}
