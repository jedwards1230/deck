package model

import (
	"testing"
)

func TestVisibleContent(t *testing.T) {
	tests := []struct {
		name       string
		chunks     []Chunk
		chunkIndex int
		want       string
	}{
		{
			name:       "single chunk index 0",
			chunks:     []Chunk{{Content: "Hello world"}},
			chunkIndex: 0,
			want:       "Hello world",
		},
		{
			name: "multiple chunks index 0 returns only first",
			chunks: []Chunk{
				{Content: "First"},
				{Content: " Second"},
				{Content: " Third"},
			},
			chunkIndex: 0,
			want:       "First",
		},
		{
			name: "multiple chunks index 1 returns first two",
			chunks: []Chunk{
				{Content: "Alpha"},
				{Content: " Bravo"},
				{Content: " Charlie"},
			},
			chunkIndex: 1,
			want:       "Alpha Bravo",
		},
		{
			name: "multiple chunks last index returns all",
			chunks: []Chunk{
				{Content: "One"},
				{Content: " Two"},
				{Content: " Three"},
			},
			chunkIndex: 2,
			want:       "One Two Three",
		},
		{
			name: "index beyond chunks clamps to last",
			chunks: []Chunk{
				{Content: "A"},
				{Content: "B"},
			},
			chunkIndex: 10,
			want:       "AB",
		},
		{
			name: "preserves whitespace between chunks",
			chunks: []Chunk{
				{Content: "Line 1\n"},
				{Content: "\nLine 2"},
			},
			chunkIndex: 1,
			want:       "Line 1\n\nLine 2",
		},
		{
			name:       "empty single chunk",
			chunks:     []Chunk{{Content: ""}},
			chunkIndex: 0,
			want:       "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slide := Slide{Chunks: tt.chunks}
			got := slide.VisibleContent(tt.chunkIndex)

			if got != tt.want {
				t.Errorf("VisibleContent(%d) = %q, want %q", tt.chunkIndex, got, tt.want)
			}
		})
	}
}
