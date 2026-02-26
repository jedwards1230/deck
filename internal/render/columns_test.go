package render

import (
	"strings"
	"testing"

	"github.com/jedwards1230/deck/internal/model"
)

func TestPadToHeight(t *testing.T) {
	t.Run("pads short content to target lines", func(t *testing.T) {
		content := "line1\nline2"
		got := padToHeight(content, 10, 5)
		lines := strings.Split(got, "\n")

		if len(lines) != 5 {
			t.Errorf("padToHeight() produced %d lines, want 5", len(lines))
		}
	})

	t.Run("pads each line to target width", func(t *testing.T) {
		content := "hi"
		got := padToHeight(content, 10, 1)
		lines := strings.Split(got, "\n")

		if len(lines) != 1 {
			t.Fatalf("padToHeight() produced %d lines, want 1", len(lines))
		}

		// The line "hi" (2 chars) should be padded to width 10 with 8 spaces.
		if len(lines[0]) != 10 {
			t.Errorf("padToHeight() line length = %d, want 10", len(lines[0]))
		}
	})

	t.Run("does not truncate when content already at target height", func(t *testing.T) {
		content := "a\nb\nc"
		got := padToHeight(content, 5, 3)
		lines := strings.Split(got, "\n")

		if len(lines) != 3 {
			t.Errorf("padToHeight() produced %d lines, want 3", len(lines))
		}
	})

	t.Run("blank lines have correct width", func(t *testing.T) {
		content := "x"
		got := padToHeight(content, 8, 3)
		lines := strings.Split(got, "\n")

		if len(lines) != 3 {
			t.Fatalf("padToHeight() produced %d lines, want 3", len(lines))
		}

		// Blank padding lines should be exactly `width` spaces.
		for i := 1; i < len(lines); i++ {
			if len(lines[i]) != 8 {
				t.Errorf("padToHeight() blank line %d length = %d, want 8", i, len(lines[i]))
			}
		}
	})
}

func TestRenderColumns(t *testing.T) {
	t.Run("two equal columns", func(t *testing.T) {
		columns := []string{"Left content", "Right content"}
		layout := model.ColumnLayout{Ratios: []int{1, 1}}
		cache := NewRendererCache(true)

		got, err := RenderColumns(columns, layout, 82, cache)
		if err != nil {
			t.Fatalf("RenderColumns() error: %v", err)
		}

		if got == "" {
			t.Fatal("RenderColumns() returned empty string")
		}

		// Both columns' text should appear in the output.
		if !strings.Contains(got, "Left") {
			t.Errorf("RenderColumns() missing 'Left', got:\n%s", got)
		}
		if !strings.Contains(got, "Right") {
			t.Errorf("RenderColumns() missing 'Right', got:\n%s", got)
		}
	})

	t.Run("empty layout returns empty string", func(t *testing.T) {
		columns := []string{"content"}
		layout := model.ColumnLayout{Ratios: []int{}}
		cache := NewRendererCache(true)

		got, err := RenderColumns(columns, layout, 80, cache)
		if err != nil {
			t.Fatalf("RenderColumns() error: %v", err)
		}

		if got != "" {
			t.Errorf("RenderColumns() = %q, want empty string for empty layout", got)
		}
	})

	t.Run("more columns than widths does not panic", func(t *testing.T) {
		columns := []string{"A", "B", "C", "D"}
		layout := model.ColumnLayout{Ratios: []int{1, 1}} // only 2 ratios
		cache := NewRendererCache(true)

		// Should not panic; extra columns beyond the layout are ignored.
		got, err := RenderColumns(columns, layout, 82, cache)
		if err != nil {
			t.Fatalf("RenderColumns() error: %v", err)
		}

		if got == "" {
			t.Error("RenderColumns() returned empty string; expected some output")
		}
	})

	t.Run("single column", func(t *testing.T) {
		columns := []string{"Solo content"}
		layout := model.ColumnLayout{Ratios: []int{1}}
		cache := NewRendererCache(true)

		got, err := RenderColumns(columns, layout, 80, cache)
		if err != nil {
			t.Fatalf("RenderColumns() error: %v", err)
		}

		if !strings.Contains(got, "Solo") {
			t.Errorf("RenderColumns() missing 'Solo', got:\n%s", got)
		}
	})
}
