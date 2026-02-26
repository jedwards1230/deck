package render

import (
	"testing"
)

func TestNewRendererCache(t *testing.T) {
	t.Run("creates non-nil cache", func(t *testing.T) {
		cache := NewRendererCache(true)
		if cache == nil {
			t.Fatal("NewRendererCache() returned nil")
		}
	})

	t.Run("dark mode flag is stored", func(t *testing.T) {
		dark := NewRendererCache(true)
		light := NewRendererCache(false)

		if dark.isDark != true {
			t.Error("expected isDark=true for dark cache")
		}
		if light.isDark != false {
			t.Error("expected isDark=false for light cache")
		}
	})
}

func TestRendererCacheGet(t *testing.T) {
	t.Run("returns a renderer", func(t *testing.T) {
		cache := NewRendererCache(true)
		r, err := cache.Get(80)
		if err != nil {
			t.Fatalf("Get(80) error: %v", err)
		}
		if r == nil {
			t.Fatal("Get(80) returned nil renderer")
		}
	})

	t.Run("returns cached renderer on same width", func(t *testing.T) {
		cache := NewRendererCache(true)

		r1, err := cache.Get(80)
		if err != nil {
			t.Fatalf("first Get(80) error: %v", err)
		}

		r2, err := cache.Get(80)
		if err != nil {
			t.Fatalf("second Get(80) error: %v", err)
		}

		if r1 != r2 {
			t.Error("Get(80) returned different renderer on second call with same width")
		}
	})

	t.Run("creates new renderer when width changes", func(t *testing.T) {
		cache := NewRendererCache(true)

		r1, err := cache.Get(80)
		if err != nil {
			t.Fatalf("Get(80) error: %v", err)
		}

		r2, err := cache.Get(120)
		if err != nil {
			t.Fatalf("Get(120) error: %v", err)
		}

		if r1 == r2 {
			t.Error("Get returned same renderer for different widths")
		}
	})

	t.Run("clamps minimum render width", func(t *testing.T) {
		cache := NewRendererCache(true)

		// Width 5 minus glamourGutter (3) = 2, which is below 10, so it clamps to 10.
		r, err := cache.Get(5)
		if err != nil {
			t.Fatalf("Get(5) error: %v", err)
		}
		if r == nil {
			t.Fatal("Get(5) returned nil renderer")
		}

		// The stored width should be the clamped value (10).
		if cache.width != 10 {
			t.Errorf("cache.width = %d, want 10", cache.width)
		}
	})

	t.Run("light mode returns a renderer", func(t *testing.T) {
		cache := NewRendererCache(false)
		r, err := cache.Get(80)
		if err != nil {
			t.Fatalf("Get(80) light mode error: %v", err)
		}
		if r == nil {
			t.Fatal("Get(80) light mode returned nil renderer")
		}
	})
}

func TestRendererCacheInvalidate(t *testing.T) {
	t.Run("clears the cache", func(t *testing.T) {
		cache := NewRendererCache(true)

		r1, err := cache.Get(80)
		if err != nil {
			t.Fatalf("Get(80) error: %v", err)
		}

		cache.Invalidate()

		r2, err := cache.Get(80)
		if err != nil {
			t.Fatalf("Get(80) after invalidate error: %v", err)
		}

		if r1 == r2 {
			t.Error("Invalidate did not clear cache; Get returned same renderer pointer")
		}
	})

	t.Run("resets stored width", func(t *testing.T) {
		cache := NewRendererCache(true)

		_, err := cache.Get(80)
		if err != nil {
			t.Fatalf("Get(80) error: %v", err)
		}

		cache.Invalidate()

		if cache.width != 0 {
			t.Errorf("after Invalidate, cache.width = %d, want 0", cache.width)
		}
		if cache.renderer != nil {
			t.Error("after Invalidate, cache.renderer is not nil")
		}
	})
}
