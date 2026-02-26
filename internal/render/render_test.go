package render

import (
	"strings"
	"testing"

	"github.com/jedwards1230/deck/internal/model"
)

func TestRenderSlide(t *testing.T) {
	t.Run("renders simple slide", func(t *testing.T) {
		slide := model.Slide{
			Chunks: []model.Chunk{
				{Content: "# Hello\n\nThis is a test slide."},
			},
		}
		cache := NewRendererCache(true)

		got, err := RenderSlide(slide, 0, 80, cache)
		if err != nil {
			t.Fatalf("RenderSlide() error: %v", err)
		}

		if got == "" {
			t.Fatal("RenderSlide() returned empty string for non-empty content")
		}

		// Glamour renders the heading; the word "Hello" should appear.
		if !strings.Contains(got, "Hello") {
			t.Errorf("RenderSlide() output missing 'Hello', got:\n%s", got)
		}
	})

	t.Run("handles empty content", func(t *testing.T) {
		slide := model.Slide{
			Chunks: []model.Chunk{
				{Content: ""},
			},
		}
		cache := NewRendererCache(true)

		got, err := RenderSlide(slide, 0, 80, cache)
		if err != nil {
			t.Fatalf("RenderSlide() error: %v", err)
		}

		if got != "" {
			t.Errorf("RenderSlide() = %q, want empty string for empty content", got)
		}
	})

	t.Run("handles whitespace-only content", func(t *testing.T) {
		slide := model.Slide{
			Chunks: []model.Chunk{
				{Content: "   \n\n  "},
			},
		}
		cache := NewRendererCache(true)

		got, err := RenderSlide(slide, 0, 80, cache)
		if err != nil {
			t.Fatalf("RenderSlide() error: %v", err)
		}

		if got != "" {
			t.Errorf("RenderSlide() = %q, want empty string for whitespace-only content", got)
		}
	})

	t.Run("renders only first chunk at index 0", func(t *testing.T) {
		slide := model.Slide{
			Chunks: []model.Chunk{
				{Content: "First chunk content"},
				{Content: "\n\nSecond chunk content"},
			},
		}
		cache := NewRendererCache(true)

		got, err := RenderSlide(slide, 0, 80, cache)
		if err != nil {
			t.Fatalf("RenderSlide() error: %v", err)
		}

		if !strings.Contains(got, "First") {
			t.Errorf("RenderSlide() at chunk 0 missing 'First', got:\n%s", got)
		}
		if strings.Contains(got, "Second") {
			t.Errorf("RenderSlide() at chunk 0 should not contain 'Second', got:\n%s", got)
		}
	})

	t.Run("renders multiple chunks at higher index", func(t *testing.T) {
		slide := model.Slide{
			Chunks: []model.Chunk{
				{Content: "Alpha content"},
				{Content: "\n\nBravo content"},
			},
		}
		cache := NewRendererCache(true)

		got, err := RenderSlide(slide, 1, 80, cache)
		if err != nil {
			t.Fatalf("RenderSlide() error: %v", err)
		}

		if !strings.Contains(got, "Alpha") {
			t.Errorf("RenderSlide() at chunk 1 missing 'Alpha', got:\n%s", got)
		}
		if !strings.Contains(got, "Bravo") {
			t.Errorf("RenderSlide() at chunk 1 missing 'Bravo', got:\n%s", got)
		}
	})
}
