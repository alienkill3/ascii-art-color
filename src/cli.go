package src

// cli.go handles command-line argument parsing and configuration.
// It parses flags (--color, --align, --output, --s), validates positional arguments,
// and constructs a CLIConfig struct containing the input string, font selection,
// and rendering options like color, alignment, and substring filtering.
// The ParseCLIArgs function separates flag arguments from positional arguments to
// properly handle user input and font specification across multiple argument patterns.

import (
	"errors"
	"flag"
	"io"
	"strings"
)

const OutputUsage = "Usage: go run . [FLAGS] [SUBSTRING] [STRING] [BANNER]\n\nEX: go run . --color=red so something standard"

var ErrUsage = errors.New(OutputUsage)

type CLIConfig struct {
	Input      string
	Font       FontType
	OutputFile string
	Options    BannerOptions
}

func ParseCLIArgs(args []string) (CLIConfig, error) {
	cfg := CLIConfig{Font: Standard}

	// Separate flags from positional arguments to handle values starting with "-"
	flagArgs, positionalArgs := separateArgsAndFlags(args)

	fs := flag.NewFlagSet("ascii-art", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	// fontFlag := fs.String("f", "", "banner ASCII font")
	colorFlag := fs.String("color", "", "text color: color name, hex format (RRGGBB), or RGB format rgb(R, G, B)")
	alignFlag := fs.String("align", "left", "alignment of the output")
	subStrFlag := fs.String("s", "", "apply color only to matching substring")
	// fs.String("substr", "", "apply color only to matching substring")
	outputFlag := fs.String("output", "", "write output to file")

	if err := fs.Parse(flagArgs); err != nil {
		return cfg, ErrUsage
	}

	// Combine parsed positional args with separated positional args
	parts := append(fs.Args(), positionalArgs...)
	if len(parts) > 3 {
		return cfg, ErrUsage
	}
	// fmt.Println(parts)
	// Handle positional arguments: [substring] string [font]
	// Logic:
	// 1 arg: string
	// 2 args: string font (if args[1] is valid font)
	// 3 args: color string font

	if len(parts) == 1 {
		cfg.Input = parts[0]
	} else if len(parts) == 2 {
		// check if color flag is set
		if *colorFlag != "" && !IsValidFont(parts[1]) {

			*subStrFlag = parts[0]
			cfg.Input = parts[1]
		} else if IsValidFont(parts[1]) {
			// Two args with second being font
			cfg.Input = parts[0]
			cfg.Font = FontType(parts[1])

		} else {
			// Two args without font being second arg
			// This is substring string
			return cfg, errors.New("Error: wrong font flag\nValid fonts: standard, shadow, thinkertoy")
		}
	} else if len(parts) == 3 {
		// Three args: substring string font
		if *subStrFlag == "" {
			*subStrFlag = parts[0]
		}
		cfg.Input = parts[1]
		if !IsValidFont(parts[2]) {
			return cfg, errors.New("Error: wrong font flag\nValid fonts: standard, shadow, thinkertoy\n")
		}
		cfg.Font = FontType(parts[2])
	}

	if strings.TrimSpace(*colorFlag) != "" {
		var parsedColor int
		var err error

		// Try RGB format first, then named color, then hex
		if strings.HasPrefix(*colorFlag, "rgb(") {
			parsedColor, err = ParseRGBColor(*colorFlag)
			if err != nil {
				return cfg, errors.New("Error: " + err.Error())
			}
		} else if strings.HasPrefix(*colorFlag, "#") {
			// If it starts with #, treat it as hex
			parsedColor, err = ParseHexColor(*colorFlag)
			if err != nil {
				return cfg, errors.New("Error: " + err.Error())
			}
		} else {
			// Otherwise, try named color first
			parsedColor, err = ParseNamedColor(*colorFlag)
			if err != nil {
				return cfg, errors.New("Error: " + err.Error())
			}
		}
		cfg.Options.ColorHex = &parsedColor
	}

	// Validate that substring is only used with color flag
	if strings.TrimSpace(*subStrFlag) != "" && cfg.Options.ColorHex == nil {
		return cfg, errors.New("Error: substring flag requires --color flag to be set\nCorrect format: go run . [color] [substring] [string] [banner]\nExample: go run . --color=red e hello shadow\n")
	}

	cfg.Options.SubStr = *subStrFlag
	cfg.OutputFile = *outputFlag
	cfg.Options.Align = *alignFlag
	return cfg, nil
}
