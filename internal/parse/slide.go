package parse

import "strings"

// SplitSlides splits raw content into individual slide strings using \n---\n delimiter.
func SplitSlides(content string) []string {
	// Normalize line endings
	content = strings.ReplaceAll(content, "\r\n", "\n")

	slides := strings.Split(content, "\n---\n")

	// Remove empty trailing slide if present
	if len(slides) > 0 && strings.TrimSpace(slides[len(slides)-1]) == "" {
		slides = slides[:len(slides)-1]
	}

	return slides
}
