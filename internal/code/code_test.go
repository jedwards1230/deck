package code

import (
	"testing"
)

func TestExtractBlocks(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    []Block
	}{
		{
			name:    "finds go code block",
			content: "# Title\n\n```go\nfmt.Println(\"hello\")\n```\n",
			want: []Block{
				{Language: "go", Code: "fmt.Println(\"hello\")"},
			},
		},
		{
			name: "finds multiple blocks",
			content: "```python\nprint('a')\n```\n\ntext\n\n```bash\necho hi\n```\n",
			want: []Block{
				{Language: "python", Code: "print('a')"},
				{Language: "bash", Code: "echo hi"},
			},
		},
		{
			name:    "no blocks returns empty",
			content: "# Just a title\n\nSome plain text content.\n",
			want:    []Block{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractBlocks(tt.content)
			if len(got) != len(tt.want) {
				t.Fatalf("ExtractBlocks() returned %d blocks, want %d", len(got), len(tt.want))
			}
			for i := range got {
				if got[i].Language != tt.want[i].Language {
					t.Errorf("block[%d].Language = %q, want %q", i, got[i].Language, tt.want[i].Language)
				}
				if got[i].Code != tt.want[i].Code {
					t.Errorf("block[%d].Code = %q, want %q", i, got[i].Code, tt.want[i].Code)
				}
			}
		})
	}
}

func TestStripHiddenLines(t *testing.T) {
	tests := []struct {
		name string
		code string
		want string
	}{
		{
			name: "removes lines starting with ///",
			code: "/// package main\n/// import \"fmt\"\nfmt.Println(\"hello\")",
			want: "fmt.Println(\"hello\")",
		},
		{
			name: "keeps normal lines intact",
			code: "line 1\nline 2\nline 3",
			want: "line 1\nline 2\nline 3",
		},
		{
			name: "handles mixed content",
			code: "/// hidden setup\nvisible line\n/// hidden teardown\nanother visible",
			want: "visible line\nanother visible",
		},
		{
			name: "strips indented hidden lines",
			code: "  /// indented hidden\nvisible",
			want: "visible",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StripHiddenLines(tt.code)
			if got != tt.want {
				t.Errorf("StripHiddenLines() = %q, want %q", got, tt.want)
			}
		})
	}
}
