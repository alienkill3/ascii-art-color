package src

import "strings"

// isKnownFlag checks if an argument is a known flag
func isKnownFlag(arg string) bool {
	// Remove leading dashes and check the flag name
	flagName := strings.TrimLeft(arg, "-")

	// Handle flags with = separator
	if idx := strings.Index(flagName, "="); idx > 0 {
		flagName = flagName[:idx]
	}

	// List of known flag names
	knownFlags := map[string]bool{
		"color":  true,
		"align":  true,
		"s":      true,
		"substr": true,
		"output": true,
	}

	return knownFlags[flagName]
}

// needsFlagValue checks if a flag requires a value on the next argument
func needsFlagValue(arg string) bool {
	flags := map[string]bool{
		"-color":   true,
		"--color":  true,
		"-align":   true,
		"--align":  true,
		"-s":       true,
		"-substr":  true,
		"--substr": true,
		"-output":  true,
		"--output": true,
	}
	return flags[arg]
}
