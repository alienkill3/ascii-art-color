package tests

import (
	"bytes"
	"io"
	"os"
	"testing"

	"ascii-art/src"
)

func TestPrintBanner(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		fontMap map[rune][]string
		want    string
	}{
		{
			name:  "empty string",
			input: "",
			fontMap: map[rune][]string{
				'A': {"line1", "line2", "line3", "line4", "line5", "line6", "line7", "line8"},
			},
			want: "\n\n\n\n\n\n\n\n",
		},
		{
			name:  "single character",
			input: "A",
			fontMap: map[rune][]string{
				'A': {"1", "2", "3", "4", "5", "6", "7", "8"},
			},
			want: "1\n2\n3\n4\n5\n6\n7\n8\n",
		},
		{
			name:  "multiple characters",
			input: "AB",
			fontMap: map[rune][]string{
				'A': {"a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8"},
				'B': {"b1", "b2", "b3", "b4", "b5", "b6", "b7", "b8"},
			},
			want: "a1b1\na2b2\na3b3\na4b4\na5b5\na6b6\na7b7\na8b8\n",
		},
		{
			name:  "missing character in fontMap",
			input: "AZ",
			fontMap: map[rune][]string{
				'A': {"a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8"},
			},
			want: "a1\na2\na3\na4\na5\na6\na7\na8\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// redirect the file descriptor to old and
			// create a pipe in memory so we can check the output of PrintBanner
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			src.PrintBanner(tt.input, "", tt.fontMap, 0xffffff)

			// we close the write end of pipe (w) so the (r) observes EOF
			w.Close()
			os.Stdout = old
			var buf bytes.Buffer
			io.Copy(&buf, r)
			r.Close()
			if buf.String() != tt.want {
				t.Errorf("got %q, want %q", buf.String(), tt.want)
			}
		})
	}
}
