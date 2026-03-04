package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"ascii-art/src"
)

type hexFlag int

func (h *hexFlag) String() string {
	return fmt.Sprintf("0x%x", int(*h))
}

func (h *hexFlag) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return err
	}
	*h = hexFlag(v)
	return nil
}

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

	ws, err := src.IoctlGetWinsize(int(os.Stdout.Fd()))
	width := uint16(80) // default width
	if err == nil {
		width = ws.Cols
	}

	// Render once, then choose output target (stdout or file).
	input := strings.ReplaceAll(cfg.Input, "\\n", "\n")
	content := src.RenderInputString(input, fontMap, cfg.Options, width)
	if cfg.OutputFile != "" {
		err = os.WriteFile(cfg.OutputFile, []byte(content), 0o644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error writing output file")
			os.Exit(1)
		}
		return
	}
	// for _, ln := range strings.Split(content, "\n") {
	// 	if cfg.Options.ColorHex != nil {
	// 		fmt.Printf("%*s\n", int(ws.Cols)+57, ln)
	// 	} else {
	// 		fmt.Printf("%*s\n", int(ws.Cols), ln)
	// 	}
	// }
	fmt.Print(content)
	// fmt.Printf("%*s\n", ws.Cols, "test")
}
