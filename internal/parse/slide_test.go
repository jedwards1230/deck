package parse

import (
	"testing"
)

func TestSplitSlides(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "single slide no delimiter",
			input: "# Hello World\n\nSome content.",
			want:  []string{"# Hello World\n\nSome content."},
		},
		{
			name:  "two slides",
			input: "# Slide 1\n---\n# Slide 2",
			want:  []string{"# Slide 1", "# Slide 2"},
		},
		{
			name:  "multiple slides",
			input: "# One\n---\n# Two\n---\n# Three\n---\n# Four",
			want:  []string{"# One", "# Two", "# Three", "# Four"},
		},
		{
			name:  "empty trailing slide removed",
			input: "# Slide 1\n---\n# Slide 2\n---\n",
			want:  []string{"# Slide 1", "# Slide 2"},
		},
		{
			name:  "empty trailing slide with whitespace removed",
			input: "# Slide 1\n---\n   \n  \t  ",
			want:  []string{"# Slide 1"},
		},
		{
			name:  "windows line endings",
			input: "# Slide 1\r\n---\r\n# Slide 2",
			want:  []string{"# Slide 1", "# Slide 2"},
		},
		{
			// The --- delimiter inside a fenced code block is NOT special-cased:
			// SplitSlides splits on any \n---\n occurrence, so code blocks
			// containing --- will be treated as a slide delimiter.
			name:  "delimiter inside code block splits (known behavior)",
			input: "# Slide 1\n\n```\nsome\n---\ncode\n```",
			want:  []string{"# Slide 1\n\n```\nsome", "code\n```"},
		},
		{
			// Empty string splits to [""], but then the empty-trailing-slide
			// logic removes it, resulting in an empty slice.
			name:  "empty content",
			input: "",
			want:  []string{},
		},
		{
			name:  "delimiter without surrounding content",
			input: "\n---\n",
			want:  []string{""},
		},
		{
			name:  "delimiter must have newlines on both sides",
			input: "hello---world",
			want:  []string{"hello---world"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitSlides(tt.input)

			if len(got) != len(tt.want) {
				t.Fatalf("len(slides) = %d, want %d\ngot:  %q\nwant: %q",
					len(got), len(tt.want), got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("slide[%d] = %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}
