package render

import (
	"strings"
	"testing"

	"github.com/jedwards1230/deck/internal/model"
)

func TestRenderFooter(t *testing.T) {
	t.Run("with author and date", func(t *testing.T) {
		fm := model.Frontmatter{
			Author: "Alice",
			Date:   "2025-01-15",
			Paging: "Slide %d / %d",
		}

		got := RenderFooter(fm, 0, 5, 80)

		if !strings.Contains(got, "Alice") {
			t.Errorf("RenderFooter() missing author, got %q", got)
		}
		if !strings.Contains(got, "2025-01-15") {
			t.Errorf("RenderFooter() missing date, got %q", got)
		}
	})

	t.Run("with empty frontmatter", func(t *testing.T) {
		fm := model.Frontmatter{}

		got := RenderFooter(fm, 0, 1, 80)

		// Should still render without panicking; divider line is always present
		if !strings.Contains(got, "\u2500") {
			t.Errorf("RenderFooter() missing divider, got %q", got)
		}
	})

	t.Run("with custom footer template", func(t *testing.T) {
		fm := model.Frontmatter{
			Author: "Bob",
			Date:   "2025-06",
			Footer: "{author} - {current_slide}/{total_slides}",
		}

		got := RenderFooter(fm, 2, 10, 80)

		if !strings.Contains(got, "Bob") {
			t.Errorf("RenderFooter() missing author in custom footer, got %q", got)
		}
		if !strings.Contains(got, "3/10") {
			t.Errorf("RenderFooter() missing slide numbers in custom footer, got %q", got)
		}
	})
}

func TestExpandFooterTemplate(t *testing.T) {
	tests := []struct {
		name         string
		tmpl         string
		fm           model.Frontmatter
		currentSlide int
		totalSlides  int
		wantContains []string
	}{
		{
			name:         "all variables",
			tmpl:         "{author} | {date} | {current_slide} of {total_slides}",
			fm:           model.Frontmatter{Author: "Eve", Date: "2025"},
			currentSlide: 3,
			totalSlides:  20,
			wantContains: []string{"Eve", "2025", "4", "20"},
		},
		{
			name:         "partial variables",
			tmpl:         "Slide {current_slide}",
			fm:           model.Frontmatter{},
			currentSlide: 0,
			totalSlides:  5,
			wantContains: []string{"Slide 1"},
		},
		{
			name:         "no variables",
			tmpl:         "static footer",
			fm:           model.Frontmatter{},
			currentSlide: 0,
			totalSlides:  1,
			wantContains: []string{"static footer"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expandFooterTemplate(tt.tmpl, tt.fm, tt.currentSlide, tt.totalSlides)
			for _, want := range tt.wantContains {
				if !strings.Contains(got, want) {
					t.Errorf("expandFooterTemplate() = %q, want it to contain %q", got, want)
				}
			}
		})
	}
}
