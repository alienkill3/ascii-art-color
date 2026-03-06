# ASCII-Art Generator

## Description
This project is a command-line tool written in **Go** that converts a given string into a large, stylized **ASCII graphical representation**. It maps input characters to 8-line high blocks from specific banner templates to create banner-style art in the terminal. The tool also supports **color output** and **text alignment** with flexible flag-based configuration.

## Requirments
*   Golang installed (recommended version 1.25.6).
*   A terminal that supports ANSI color output.
## Usage
The program is executed through the terminal using the following command:

```bash
go run . [FLAGS] [SUBSTRING] 'INPUT_STRING' [BANNER_FONT]
```

*   Keep in mind that you have to put your string inside single quotes or escape the special characters, ex. `'!@#'` or `\!\@\#`
*   The single quote `(')` approach is strongly recommended

### Flags:
*   `--color=<value>`: Apply color to the output (required for substring coloring)
    *   Named colors: `red`, `green`, `blue`, `yellow`, `cyan`, `purple`, `orange`, `pink`
    *   Hex format: `--color=#FF5733` (6-digit hex)
        * Valid hex digits are `a-f` and `0-9`
    *   RGB format: `--color=rgb'(255,87,51)'`
        * Valid numbers are `0-255`
*   `--align=<value>`: Set text alignment
    *   Options: `left` (default), `right`, `center`, `justify`
    *   Example: `--align=right`
*   `(optional) -s <substring>` or `--substr=<substring>`: Apply color only to matching substring (requires `--color`). 
    * The program will also recognize the substring without the flag prefix
*   `--output=<file>`: Write output to a file instead of stdout
    *   Example: `--output=result.txt`

### Positional Arguments:
*   **[INPUT_STRING]**: The text you want to convert.
    * Supports uppercase/lowercase letters, numbers, spaces, special characters, and newline (`\n`) characters. 
        * To avoid any syntax conflicts with the terminal we should always put the `INPUT_STRING` inside single quotes `(')`
*   **[BANNER_FONT]**: Optional font selection. Defaults to `standard`.
    *   Options: `standard`, `shadow`, `thinkertoy`
*   **[SUBSTRING]**: When `--color` is used, the first positional becomes the substring, second becomes the string, third becomes the font. 
    * Can also be left `empty` to colorize the whole string instead.

## Features
*   **Standard Library Only**: Built entirely using standard Go packages.
*   **Template Support**: Includes three distinct styles: `shadow`, `standard`, and `thinkertoy`.
*   **Newline Handling**: Correctly interprets `\n` in the input string to move to a new block of ASCII art and manages multiple newlines for vertical spacing.
*   **Character Mapping**: Maps each character to a fixed **8-line height** graphical representation.
*   **Color Support**: Apply colors to the entire output or selectively to matching substrings.
*   **Text Alignment**: Support for left, right, center, and justify alignment options.
*   **File Output**: Option to write rendered ASCII art directly to a file.
*   **Flexible Input**: Handles arguments starting with `-` (negative numbers, flag-like text) as valid input.

## Implementation Approach

- Pipeline approach. For more info refer to [Architecture](/docs/architecture.md) and [Diagram](/docs/diagram.md).

## Example Usage

*   For example usage please refer to [Examples](/docs/examples.md).


