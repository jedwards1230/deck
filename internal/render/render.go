package render

import (
	"strings"

	"github.com/jedwards1230/deck/internal/model"
)

// RenderSlide renders the visible portion of a slide at the given width.
// If the slide has a column layout with column content, it renders in
// multi-column mode. Otherwise renders as a single block of markdown.
func RenderSlide(slide model.Slide, chunkIndex, width int, cache *RendererCache) (string, error) {
	// Column layout rendering
	if slide.Layout != nil && len(slide.Columns) > 0 {
		return RenderColumns(slide.Columns, *slide.Layout, width, cache)
	}

	// Standard single-column rendering
	content := strings.TrimSpace(slide.VisibleContent(chunkIndex))
	if content == "" {
		return "", nil
	}

	renderer, err := cache.Get(width)
	if err != nil {
		return content, err
	}

	rendered, err := renderer.Render(content)
	if err != nil {
		return content, err
	}

	return rendered, nil
}
