package src

import (
	"reflect"
	"testing"
)

func TestCutSubString(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		subStr string
		want   []StringColor
	}{
		{
			name:   "empty substring returns full string without color",
			str:    "hello",
			subStr: "",
			want:   []StringColor{{Str: "hello", Flag: false}},
		},
		{
			name:   "leading and trailing matches",
			str:    "abXXcdXX",
			subStr: "XX",
			want: []StringColor{
				{Str: "ab", Flag: false},
				{Str: "XX", Flag: true},
				{Str: "cd", Flag: false},
				{Str: "XX", Flag: true},
			},
		},
		{
			name:   "overlapping pattern uses first non-overlapping matches",
			str:    "aaaa",
			subStr: "aa",
			want: []StringColor{
				{Str: "aa", Flag: true},
				{Str: "aa", Flag: true},
			},
		},
		{
			name:   "no match returns original string as plain",
			str:    "hello",
			subStr: "z",
			want:   []StringColor{{Str: "hello", Flag: false}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CutSubString(tt.str, tt.subStr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("CutSubString(%q, %q) = %#v, want %#v", tt.str, tt.subStr, got, tt.want)
			}
		})
	}
}

