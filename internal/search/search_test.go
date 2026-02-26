package search

import (
	"testing"

	"github.com/jedwards1230/deck/internal/model"
)

func makeSlides(contents ...string) []model.Slide {
	slides := make([]model.Slide, len(contents))
	for i, c := range contents {
		slides[i] = model.Slide{
			Chunks: []model.Chunk{{Content: c}},
		}
	}
	return slides
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name         string
		slides       []model.Slide
		query        string
		currentSlide int
		wantFound    bool
		wantIndex    int
	}{
		{
			name:         "finds matching slide",
			slides:       makeSlides("# Intro", "# Topic A", "# Conclusion"),
			query:        "Topic",
			currentSlide: 0,
			wantFound:    true,
			wantIndex:    1,
		},
		{
			name:         "wraps around",
			slides:       makeSlides("# Target", "# Middle", "# End"),
			query:        "Target",
			currentSlide: 1,
			wantFound:    true,
			wantIndex:    0,
		},
		{
			name:         "case-insensitive with /i",
			slides:       makeSlides("# hello world", "# other"),
			query:        "HELLO/i",
			currentSlide: 0,
			wantFound:    true,
			wantIndex:    0,
		},
		{
			name:         "regex pattern",
			slides:       makeSlides("# intro", "has numbers 42", "# end"),
			query:        `\d{2,}`,
			currentSlide: 0,
			wantFound:    true,
			wantIndex:    1,
		},
		{
			name:         "no match returns Found=false",
			slides:       makeSlides("# Alpha", "# Beta", "# Gamma"),
			query:        "Delta",
			currentSlide: 0,
			wantFound:    false,
			wantIndex:    0,
		},
		{
			name:         "empty query returns Found=false",
			slides:       makeSlides("# Slide"),
			query:        "",
			currentSlide: 0,
			wantFound:    false,
			wantIndex:    0,
		},
		{
			name:         "empty slides returns Found=false",
			slides:       []model.Slide{},
			query:        "anything",
			currentSlide: 0,
			wantFound:    false,
			wantIndex:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Search(tt.slides, tt.query, tt.currentSlide)
			if got.Found != tt.wantFound {
				t.Errorf("Search().Found = %v, want %v", got.Found, tt.wantFound)
			}
			if got.Found && got.SlideIndex != tt.wantIndex {
				t.Errorf("Search().SlideIndex = %d, want %d", got.SlideIndex, tt.wantIndex)
			}
		})
	}
}

func TestSearchNext(t *testing.T) {
	t.Run("skips current slide", func(t *testing.T) {
		slides := makeSlides("# Match here", "# No match", "# Match here too")
		got := SearchNext(slides, "Match", 0)

		if !got.Found {
			t.Fatal("SearchNext().Found = false, want true")
		}
		if got.SlideIndex != 2 {
			t.Errorf("SearchNext().SlideIndex = %d, want 2", got.SlideIndex)
		}
	})

	t.Run("wraps to find match before current", func(t *testing.T) {
		slides := makeSlides("# Target", "# other", "# other")
		got := SearchNext(slides, "Target", 2)

		if !got.Found {
			t.Fatal("SearchNext().Found = false, want true")
		}
		if got.SlideIndex != 0 {
			t.Errorf("SearchNext().SlideIndex = %d, want 0", got.SlideIndex)
		}
	})
}

func TestSearchPrev(t *testing.T) {
	t.Run("searches backward", func(t *testing.T) {
		slides := makeSlides("# Match A", "# No match", "# Match B", "# Current")
		got := SearchPrev(slides, "Match", 3)

		if !got.Found {
			t.Fatal("SearchPrev().Found = false, want true")
		}
		if got.SlideIndex != 2 {
			t.Errorf("SearchPrev().SlideIndex = %d, want 2", got.SlideIndex)
		}
	})

	t.Run("wraps backward past beginning", func(t *testing.T) {
		slides := makeSlides("# other", "# current", "# Target")
		got := SearchPrev(slides, "Target", 0)

		if !got.Found {
			t.Fatal("SearchPrev().Found = false, want true")
		}
		if got.SlideIndex != 2 {
			t.Errorf("SearchPrev().SlideIndex = %d, want 2", got.SlideIndex)
		}
	})

	t.Run("empty query returns Found=false", func(t *testing.T) {
		slides := makeSlides("# Slide")
		got := SearchPrev(slides, "", 0)

		if got.Found {
			t.Error("SearchPrev().Found = true, want false")
		}
	})
}
