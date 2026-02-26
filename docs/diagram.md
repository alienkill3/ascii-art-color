┌─────────────────────┐
│  Command-Line Input │
│  (flag.Parse, args) │
└─────────┬───────────┘
          │
          ▼
┌─────────────────────────────┐
│  Font Processing Stage      │
│  BuildFontMap(font file)    │
│  → map[rune][]string        │
│  (fully buffered in memory) │
└─────────┬──────────────────┘
          │
          ▼
┌─────────────────────────────┐
│  Banner Rendering Stage     │
│  PrintBanner(inputString,   │
│              fontMap)       │
│  → strings.Builder          │
│  → stdout (single write)    │
└─────────┬──────────────────┘
          │
          ▼
┌─────────────────────┐
│     Terminal Output │
│  ASCII Art Printed  │
└─────────────────────┘
