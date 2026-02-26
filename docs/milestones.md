# ASCII-Art Milestones (Audit Checklist)

## Milestone 1: CLI + Validation + Pipeline
- CLI argument parsing with optional `-f` flag.
- Validation for argument count and ASCII-only input.
- End-to-end flow: parse input -> load font -> build map -> print banner.
Proof: `main.go`, `src/fonts.go`, `src/fontMap.go`, `src/printBanner.go`

## Milestone 2: Font Support
- Support for `standard`, `shadow`, `thinkertoy`.
- Embedded font files with internal reader selection.
- Rune-to-glyph mapping for printable ASCII (8-line height).
Proof: `src/fonts.go`, `src/fonts/standard.txt`, `src/fonts/shadow.txt`, `src/fonts/thinkertoy.txt`, `src/fontMap.go`

## Milestone 3: Newline Handling and Rendering Behavior
- Escaped newline handling (`\\n`) in input.
- Correct behavior for empty input.
- Correct behavior for newline-only input.
- Correct behavior for internal/double newlines.
- Correct behavior for trailing newline case.
Proof: `main.go`, `output_test.go`, `testdata/*.golden`

## Milestone 4: Tests and Regression Safety
- Unit tests for ASCII validation.
- Unit tests for banner rendering.
- Integration/process tests for main execution and exit codes.
- Golden tests for critical output cases.
Proof: `tests/isValidAscii_test.go`, `tests/printBanner_test.go`, `main_test.go`, `output_test.go`, `testdata/`

## Current Status
- `go test ./...` passes.
