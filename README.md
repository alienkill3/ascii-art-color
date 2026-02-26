# ASCII-Art Generator

## Description
This project is a command-line tool written in **Go** that converts a given string into a large, stylized **ASCII graphical representation**. It maps input characters to 8-line high blocks from specific banner templates to create banner-style art in the terminal.

## Usage
The program is executed through the terminal using the following command:

```bash
go run . [-f banner_font] 'INPUT_STRING'
```
Keep in mind that you have to put your string inside single quotes or escape the special characters, ex. ``` '!@#' or \!\@\# ```

*   **-f banner_font** (Optional): Specifies the banner template to use (**shadow**, **standard**, or **thinkertoy**). If omitted, it defaults to **standard**.
*   **INPUT_STRING**: The text you want to convert. It supports uppercase and lowercase letters, numbers, spaces, special characters, and newline (`\n`) characters.

## Features
*   **Standard Library Only**: Built entirely using standard Go packages.
*   **Template Support**: Includes three distinct styles: `shadow`, `standard`, and `thinkertoy`.
*   **Newline Handling**: Correctly interprets `\n` in the input string to move to a new block of ASCII art and manages multiple newlines for vertical spacing.
*   **Character Mapping**: Maps each character to a fixed **8-line height** graphical representation.

## Examples
The following are examples of how the program handles different inputs:

*   **Basic String:** `go run . "Hello"` renders "Hello" in the standard ASCII font.
*   **Numbers and Newlines:** `go run . "1Hello 2There\n"` correctly renders numbers, spaces, and the trailing newline.
*   **Internal Newlines:** `go run . "Hello\n\nThere"` ensures the correct vertical space is generated between the ASCII words.
*   **Empty Input:** `go run . ""` results in no output or a single empty line.
