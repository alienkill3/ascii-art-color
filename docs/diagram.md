```
┌──────────────────────────────────┐
│  Command-Line Input              │
│  (ParseCLIArgs)                  │
│  ├─ Input String                 │
│  ├─ Color (named/hex/RGB)        │
│  ├─ Substring (optional)         │
│  ├─ MainString                   │
│  └─ Font (standard/shadow/       │
│         thinkertoy)              │
│                                  │
└────────────┬─────────────────────┘
             │
             ▼
┌──────────────────────────────────┐
│  Color Parsing Stage             │
│  (ParseNamedColor/ParseHexColor/ │
│   ParseRGBColor)                 │
│  → RGB integer value             │
└────────────┬─────────────────────┘
             │
             ▼
┌──────────────────────────────────┐
│  Font Processing Stage           │
│  BuildFontMap(font file)         │
│  → map[rune][]string             │
│  (fully buffered in memory)      │
└────────────┬─────────────────────┘
             │
             ▼
┌──────────────────────────────────┐
│  Banner Rendering Stage          │
│  RenderInputString(input,        │
│                   fontMap,       │
│                   options)       │
│  → rendered ASCII art            │
└────────────┬─────────────────────┘
             │
             ▼
┌──────────────────────────────────┐
│  Colorization Stage              │
│  (Colorize with Substring Match) │
│  ├─ Apply color to full output   │
│  └─ Or only to substring matches │
└────────────┬─────────────────────┘
             │
             ▼
┌──────────────────────────────────┐
│  Output Stage                    │
│  ├─ stdout (default)             │
│  └─ File (-o/--output option)    │
└──────────────────────────────────┘
```
