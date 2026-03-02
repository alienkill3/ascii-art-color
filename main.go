package main

import (
	"fmt"
	"os"

	"ascii-art/src"
)

func main() {
	cfg, err := src.ParseCLIArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if cfg.Input == "" {
		return
	}

	if !src.IsValidAscii(cfg.Input) {
		fmt.Fprintln(os.Stderr, "Error: found non ascii character")
		os.Exit(1)
	}

	fontData := src.GetReader(cfg.Font)
	fontMap, err := src.BuildFontMap(fontData)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing font map")
		os.Exit(1)
	}

	// Render once, then choose output target (stdout or file).
	content := src.RenderInputString(cfg.Input, fontMap, cfg.Options)
	if cfg.OutputFile != "" {
		err = os.WriteFile(cfg.OutputFile, []byte(content), 0o644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error writing output file")
			os.Exit(1)
		}
		return
	}
	fmt.Print(content)
}
