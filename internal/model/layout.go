package model

// ColumnLayout defines a multi-column layout with proportional ratios.
type ColumnLayout struct {
	Ratios []int // e.g., [3, 2] means 3/5 and 2/5
}

// ColumnWidths resolves ratios to actual character widths for the given total width.
// It accounts for column gaps (2 chars between columns).
func (l ColumnLayout) ColumnWidths(totalWidth int) []int {
	if len(l.Ratios) == 0 {
		return nil
	}

	gaps := len(l.Ratios) - 1
	availableWidth := totalWidth - (gaps * 2) // 2 chars gap between columns
	if availableWidth < len(l.Ratios) {
		// Not enough space, give each column minimum 1.
		widths := make([]int, len(l.Ratios))
		for i := range widths {
			widths[i] = 1
		}
		return widths
	}

	total := 0
	for _, r := range l.Ratios {
		total += r
	}

	widths := make([]int, len(l.Ratios))
	remaining := availableWidth
	for i, r := range l.Ratios {
		if i == len(l.Ratios)-1 {
			widths[i] = remaining // last column gets remainder to avoid rounding
		} else {
			widths[i] = availableWidth * r / total
			remaining -= widths[i]
		}
	}

	return widths
}
