package diff

import (
	"testing"

	"github.com/jedwards1230/deck/internal/model"
)

func TestFindModified(t *testing.T) {
	tests := []struct {
		name string
		old  *model.Presentation
		new  *model.Presentation
		want int
	}{
		{
			name: "identical presentations returns -1",
			old: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "# Hello"}}},
					{Chunks: []model.Chunk{{Content: "# World"}}},
				},
			},
			new: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "# Hello"}}},
					{Chunks: []model.Chunk{{Content: "# World"}}},
				},
			},
			want: -1,
		},
		{
			name: "added slide returns index of new slide",
			old: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "# Slide 1"}}},
				},
			},
			new: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "# Slide 1"}}},
					{Chunks: []model.Chunk{{Content: "# Slide 2"}}},
				},
			},
			want: 1,
		},
		{
			name: "removed slide returns the shorter length",
			old: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "# Slide 1"}}},
					{Chunks: []model.Chunk{{Content: "# Slide 2"}}},
					{Chunks: []model.Chunk{{Content: "# Slide 3"}}},
				},
			},
			new: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "# Slide 1"}}},
					{Chunks: []model.Chunk{{Content: "# Slide 2"}}},
				},
			},
			want: 2,
		},
		{
			name: "modified slide content returns its index",
			old: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "# Slide 1"}}},
					{Chunks: []model.Chunk{{Content: "# Original"}}},
					{Chunks: []model.Chunk{{Content: "# Slide 3"}}},
				},
			},
			new: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "# Slide 1"}}},
					{Chunks: []model.Chunk{{Content: "# Changed"}}},
					{Chunks: []model.Chunk{{Content: "# Slide 3"}}},
				},
			},
			want: 1,
		},
		{
			name: "modified chunk count returns its index",
			old: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "# Slide 1"}}},
					{Chunks: []model.Chunk{{Content: "part 1"}}},
				},
			},
			new: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "# Slide 1"}}},
					{Chunks: []model.Chunk{{Content: "part 1"}, {Content: "part 2"}}},
				},
			},
			want: 1,
		},
		{
			name: "nil old presentation returns 0",
			old:  nil,
			new: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "# Slide 1"}}},
				},
			},
			want: 0,
		},
		{
			name: "old presentation with empty slides returns 0",
			old:  &model.Presentation{Slides: []model.Slide{}},
			new: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "# Slide 1"}}},
				},
			},
			want: 0,
		},
		{
			name: "whitespace differences are ignored",
			old: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "# Hello\n"}}},
					{Chunks: []model.Chunk{{Content: "# World\n\n"}}},
				},
			},
			new: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "\n# Hello"}}},
					{Chunks: []model.Chunk{{Content: "# World"}}},
				},
			},
			want: -1,
		},
		{
			name: "trailing newline does not cause false positive",
			old: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "# Slide 1"}}},
					{Chunks: []model.Chunk{{Content: "# Target"}}},
				},
			},
			new: &model.Presentation{
				Slides: []model.Slide{
					{Chunks: []model.Chunk{{Content: "# Slide 1\n"}}},
					{Chunks: []model.Chunk{{Content: "# CHANGED"}}},
				},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindModified(tt.old, tt.new)
			if got != tt.want {
				t.Errorf("FindModified() = %d, want %d", got, tt.want)
			}
		})
	}
}
