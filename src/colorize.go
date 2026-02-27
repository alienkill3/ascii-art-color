package src

import "fmt"

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
