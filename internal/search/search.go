package search

import (
	"regexp"
	"strings"

	"github.com/jedwards1230/deck/internal/model"
)

// Result represents a search match.
type Result struct {
	SlideIndex int
	Found      bool
}

// Search finds the next slide matching the query, starting from currentSlide.
// Wraps around to the beginning. Supports regex patterns.
// Append /i to query for case-insensitive search.
func Search(slides []model.Slide, query string, currentSlide int) Result {
	if query == "" || len(slides) == 0 {
		return Result{Found: false}
	}

	// Check for case-insensitive flag
	caseInsensitive := false
	if strings.HasSuffix(query, "/i") {
		query = strings.TrimSuffix(query, "/i")
		caseInsensitive = true
	}

	// Build regex
	flags := ""
	if caseInsensitive {
		flags = "(?i)"
	}
	pattern, err := regexp.Compile(flags + query)
	if err != nil {
		// Fall back to literal search
		return searchLiteral(slides, query, currentSlide, caseInsensitive)
	}

	// Search forward from current slide (wrapping)
	for i := range len(slides) {
		idx := (currentSlide + i) % len(slides)
		content := slides[idx].VisibleContent(len(slides[idx].Chunks) - 1)
		if pattern.MatchString(content) {
			return Result{SlideIndex: idx, Found: true}
		}
	}

	return Result{Found: false}
}

// SearchNext finds the next match after the current slide.
func SearchNext(slides []model.Slide, query string, currentSlide int) Result {
	return Search(slides, query, (currentSlide+1)%len(slides))
}

// SearchPrev finds the previous match before the current slide.
func SearchPrev(slides []model.Slide, query string, currentSlide int) Result {
	if query == "" || len(slides) == 0 {
		return Result{Found: false}
	}

	caseInsensitive := false
	if strings.HasSuffix(query, "/i") {
		query = strings.TrimSuffix(query, "/i")
		caseInsensitive = true
	}

	flags := ""
	if caseInsensitive {
		flags = "(?i)"
	}
	pattern, err := regexp.Compile(flags + query)
	if err != nil {
		return Result{Found: false}
	}

	// Search backward from current slide (wrapping)
	for i := range len(slides) {
		idx := (currentSlide - 1 - i + len(slides)) % len(slides)
		content := slides[idx].VisibleContent(len(slides[idx].Chunks) - 1)
		if pattern.MatchString(content) {
			return Result{SlideIndex: idx, Found: true}
		}
	}

	return Result{Found: false}
}

func searchLiteral(slides []model.Slide, query string, currentSlide int, caseInsensitive bool) Result {
	if caseInsensitive {
		query = strings.ToLower(query)
	}

	for i := range len(slides) {
		idx := (currentSlide + i) % len(slides)
		content := slides[idx].VisibleContent(len(slides[idx].Chunks) - 1)
		if caseInsensitive {
			content = strings.ToLower(content)
		}
		if strings.Contains(content, query) {
			return Result{SlideIndex: idx, Found: true}
		}
	}

	return Result{Found: false}
}
