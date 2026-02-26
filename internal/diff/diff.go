package diff

import (
	"strings"

	"github.com/jedwards1230/deck/internal/model"
)

// FindModified compares two presentations and returns the index of the first
// modified slide. Returns -1 if no modifications were found.
func FindModified(old, new *model.Presentation) int {
	if old == nil || len(old.Slides) == 0 {
		return 0
	}

	maxLen := min(len(old.Slides), len(new.Slides))

	for i := range maxLen {
		if slideChanged(old.Slides[i], new.Slides[i]) {
			return i
		}
	}

	// If slide count changed, jump to first new slide
	if len(new.Slides) != len(old.Slides) {
		return maxLen
	}

	return -1
}

func slideChanged(a, b model.Slide) bool {
	if len(a.Chunks) != len(b.Chunks) {
		return true
	}
	for i := range a.Chunks {
		// Normalize whitespace to avoid false positives from editor
		// trailing newlines, whitespace changes, etc.
		if normalizeContent(a.Chunks[i].Content) != normalizeContent(b.Chunks[i].Content) {
			return true
		}
	}
	return false
}

func normalizeContent(s string) string {
	return strings.TrimSpace(s)
}
