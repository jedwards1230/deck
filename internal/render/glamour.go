package render

import (
	"sync"

	"github.com/charmbracelet/glamour/v2"
	"github.com/charmbracelet/glamour/v2/styles"
)

// glamourGutter is the internal padding glamour adds to rendered content.
const glamourGutter = 3

// RendererCache caches a glamour TermRenderer keyed by width, avoiding
// re-creation on every render when the terminal size has not changed.
type RendererCache struct {
	mu       sync.Mutex
	renderer *glamour.TermRenderer
	width    int
	isDark   bool
}

// NewRendererCache creates a new renderer cache for the given color scheme.
func NewRendererCache(isDark bool) *RendererCache {
	return &RendererCache{isDark: isDark}
}

// Get returns a glamour renderer configured for the given width. If the
// cached renderer already matches, it is returned without allocation.
func (c *RendererCache) Get(width int) (*glamour.TermRenderer, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	renderWidth := width - glamourGutter
	if renderWidth < 10 {
		renderWidth = 10
	}

	if c.renderer != nil && c.width == renderWidth {
		return c.renderer, nil
	}

	style := styles.DarkStyleConfig
	if !c.isDark {
		style = styles.LightStyleConfig
	}

	r, err := glamour.NewTermRenderer(
		glamour.WithStyles(style),
		glamour.WithWordWrap(renderWidth),
	)
	if err != nil {
		return nil, err
	}

	c.renderer = r
	c.width = renderWidth
	return r, nil
}

// Invalidate clears the cached renderer, forcing re-creation on the next Get.
func (c *RendererCache) Invalidate() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.renderer = nil
	c.width = 0
}
