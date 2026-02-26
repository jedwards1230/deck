package parse

import (
	"strings"

	"github.com/jedwards1230/deck/internal/model"
	"gopkg.in/yaml.v3"
)

// ParseFrontmatter extracts YAML frontmatter from content.
// Returns the frontmatter and remaining content (without the frontmatter block).
func ParseFrontmatter(content string) (model.Frontmatter, string) {
	var fm model.Frontmatter

	trimmed := strings.TrimSpace(content)
	if !strings.HasPrefix(trimmed, "---") {
		return fm, content
	}

	// Find closing ---
	rest := trimmed[3:] // skip opening ---
	idx := strings.Index(rest, "\n---")
	if idx < 0 {
		return fm, content
	}

	yamlContent := rest[:idx]
	remaining := strings.TrimPrefix(rest[idx+4:], "\n")

	_ = yaml.Unmarshal([]byte(yamlContent), &fm)

	// Apply defaults
	if fm.Paging == "" {
		fm.Paging = "Slide %d / %d"
	}

	return fm, remaining
}
