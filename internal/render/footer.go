package render

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/jedwards1230/deck/internal/model"
)

var (
	footerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))

	footerDividerStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("238"))
)

// RenderFooter creates the footer bar with author/date on the left and
// slide paging on the right, spanning the given width.
func RenderFooter(fm model.Frontmatter, currentSlide, totalSlides, width int) string {
	left := buildLeftFooter(fm)
	right := buildRightFooter(fm, currentSlide, totalSlides)

	leftWidth := lipgloss.Width(left)
	rightWidth := lipgloss.Width(right)
	gap := width - leftWidth - rightWidth
	if gap < 0 {
		gap = 0
	}

	divider := footerDividerStyle.Render(strings.Repeat("\u2500", width))
	content := left + strings.Repeat(" ", gap) + right

	return divider + "\n" + content
}

func buildLeftFooter(fm model.Frontmatter) string {
	var parts []string
	if fm.Author != "" {
		parts = append(parts, fm.Author)
	}
	if fm.Date != "" {
		parts = append(parts, fm.Date)
	}
	if len(parts) == 0 {
		return ""
	}
	return footerStyle.Render(strings.Join(parts, " \u00b7 "))
}

func buildRightFooter(fm model.Frontmatter, currentSlide, totalSlides int) string {
	// Custom footer template takes precedence.
	if fm.Footer != "" {
		return footerStyle.Render(expandFooterTemplate(fm.Footer, fm, currentSlide, totalSlides))
	}

	if fm.Paging != "" {
		// Use the user's paging template if it contains format verbs,
		// otherwise use it as a literal string.
		paging := fmt.Sprintf(fm.Paging, currentSlide+1, totalSlides)
		return footerStyle.Render(paging)
	}

	return ""
}

func expandFooterTemplate(tmpl string, fm model.Frontmatter, currentSlide, totalSlides int) string {
	r := strings.NewReplacer(
		"{author}", fm.Author,
		"{date}", fm.Date,
		"{current_slide}", fmt.Sprintf("%d", currentSlide+1),
		"{total_slides}", fmt.Sprintf("%d", totalSlides),
	)
	return r.Replace(tmpl)
}
