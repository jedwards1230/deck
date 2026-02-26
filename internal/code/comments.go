package code

import (
	"strings"
)

// StripHiddenLines removes lines that start with /// (presenter-only lines).
// These lines are visible in the slide but hidden from execution output.
func StripHiddenLines(code string) string {
	var lines []string
	for _, line := range strings.Split(code, "\n") {
		if !strings.HasPrefix(strings.TrimSpace(line), "///") {
			lines = append(lines, line)
		}
	}
	return strings.Join(lines, "\n")
}
