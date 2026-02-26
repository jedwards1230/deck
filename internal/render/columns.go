package render

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/jedwards1230/deck/internal/model"
)

// RenderColumns renders content in a multi-column layout. Each column's
// markdown is rendered at its proportional width, then the columns are
// joined horizontally with equal height padding.
func RenderColumns(columns []string, layout model.ColumnLayout, totalWidth int, cache *RendererCache) (string, error) {
	widths := layout.ColumnWidths(totalWidth)
	if len(widths) == 0 {
		return "", nil
	}

	rendered := make([]string, len(columns))
	maxLines := 0

	for i, col := range columns {
		if i >= len(widths) {
			break
		}

		w := widths[i]
		out, err := renderColumn(col, w, cache)
		if err != nil {
			rendered[i] = col
			continue
		}
		rendered[i] = out

		if lines := strings.Count(out, "\n") + 1; lines > maxLines {
			maxLines = lines
		}
	}

	paddedCols := make([]string, 0, len(rendered))
	for i, col := range rendered {
		if i >= len(widths) {
			break
		}
		paddedCols = append(paddedCols, padToHeight(col, widths[i], maxLines))
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, paddedCols...), nil
}

func renderColumn(content string, width int, cache *RendererCache) (string, error) {
	r, err := cache.Get(width)
	if err != nil {
		return content, err
	}

	out, err := r.Render(strings.TrimSpace(content))
	if err != nil {
		return content, err
	}

	return out, nil
}

// padToHeight pads each line to the target width and extends the content
// to the target number of lines.
func padToHeight(content string, width, targetLines int) string {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		if lineWidth := lipgloss.Width(line); lineWidth < width {
			lines[i] = line + strings.Repeat(" ", width-lineWidth)
		}
	}

	blank := strings.Repeat(" ", width)
	for len(lines) < targetLines {
		lines = append(lines, blank)
	}

	return strings.Join(lines, "\n")
}
