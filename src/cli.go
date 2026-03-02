package src

import (
	"errors"
	"flag"
	"io"
	"strings"
)

const OutputUsage = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"

var ErrUsage = errors.New(OutputUsage)

type CLIConfig struct {
	Input      string
	Font       FontType
	OutputFile string
	Options    BannerOptions
}

func ParseCLIArgs(args []string) (CLIConfig, error) {
	cfg := CLIConfig{Font: Standard}
	cleanArgs := removeTestArgs(args)

	// Keep --output validation strict to match subject requirements.
	if err := validateOutputFlagFormat(cleanArgs); err != nil {
		return cfg, err
	}

	fs := flag.NewFlagSet("ascii-art", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	fontFlag := fs.String("f", "", "banner ASCII font")
	colorFlag := fs.String("color", "", "text color in hex format: RRGGBB or #RRGGBB")
	subStrFlag := fs.String("substr", "", "apply color only to matching substring")
	outputFlag := fs.String("output", "", "write output to file")

	if err := fs.Parse(cleanArgs); err != nil {
		return cfg, ErrUsage
	}

	parts := fs.Args()
	if len(parts) < 1 || len(parts) > 2 {
		return cfg, ErrUsage
	}
	cfg.Input = parts[0]

	if len(parts) == 2 {
		banner := FontType(parts[1])
		if !IsValidFont(string(banner)) {
			return cfg, errors.New("Error: wrong font flag\nValid fonts: standard, shadow, thinkertoy")
		}
		cfg.Font = banner
	}

	if *fontFlag != "" {
		if !IsValidFont(*fontFlag) {
			return cfg, errors.New("Error: wrong font flag\nValid fonts: standard, shadow, thinkertoy")
		}
		if len(parts) == 2 && FontType(*fontFlag) != cfg.Font {
			return cfg, errors.New("Error: conflicting banner values")
		}
		cfg.Font = FontType(*fontFlag)
	}

	if strings.TrimSpace(*colorFlag) != "" {
		parsedColor, err := ParseHexColor(*colorFlag)
		if err != nil {
			return cfg, errors.New("Error: invalid color flag\nUse --color RRGGBB (example: --color FF0000)")
		}
		cfg.Options.ColorHex = &parsedColor
	}

	cfg.Options.SubStr = *subStrFlag
	cfg.OutputFile = *outputFlag
	return cfg, nil
}

func removeTestArgs(args []string) []string {
	filtered := make([]string, 0, len(args))
	for _, arg := range args {
		if strings.HasPrefix(arg, "-test.") {
			continue
		}
		filtered = append(filtered, arg)
	}
	return filtered
}

func validateOutputFlagFormat(args []string) error {
	for _, arg := range args {
		switch {
		case arg == "--output":
			return ErrUsage
		case strings.HasPrefix(arg, "--output"):
			if !strings.HasPrefix(arg, "--output=") {
				return ErrUsage
			}
			if strings.TrimSpace(strings.TrimPrefix(arg, "--output=")) == "" {
				return ErrUsage
			}
		case strings.HasPrefix(arg, "-output"):
			return ErrUsage
		}
	}
	return nil
}
