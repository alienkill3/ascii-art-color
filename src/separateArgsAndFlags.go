package src

import "strings"

// separateArgsAndFlags separates recognized flags from positional arguments.
// This prevents arguments starting with "-" (like "-123" or "-abc") from being
// treated as invalid flags by the flag parser.
func separateArgsAndFlags(args []string) ([]string, []string) {
	var flagArgs []string
	var positionalArgs []string
	foundPositional := false

	// If only 1 argument and it looks like a flag, treat it as positional input
	if len(args) == 1 && strings.HasPrefix(args[0], "-") {
		return flagArgs, args
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if foundPositional {
			// Once a positional arg is found, everything after is positional
			positionalArgs = append(positionalArgs, arg)
			continue
		}

		// Check for double dash which marks end of flags
		if arg == "--" {
			flagArgs = append(flagArgs, arg)
			foundPositional = true
			continue
		}

		// If it starts with -, check if it's a known flag
		if strings.HasPrefix(arg, "-") {
			// Check if it looks like a known flag (color, align, s, substr, output)
			if isKnownFlag(arg) {
				flagArgs = append(flagArgs, arg)
				// If it's a flag with = (like --color=red), it's self-contained
				// If it's a flag without = but expects a value, get the next arg
				if !strings.Contains(arg, "=") && needsFlagValue(arg) {
					i++ // Skip next iteration since we'll add it to flags
					if i < len(args) {
						flagArgs = append(flagArgs, args[i])
					}
				}
			} else {
				// Unknown flag-like argument, treat as positional
				foundPositional = true
				positionalArgs = append(positionalArgs, arg)
			}
		} else {
			// Doesn't start with -, it's positional
			flagArgs = append(flagArgs, arg)
		}
	}

	return flagArgs, positionalArgs
}
