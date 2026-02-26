package parse

import (
	"regexp"
	"strings"

	"github.com/jedwards1230/deck/internal/model"
)

var (
	pauseRegex  = regexp.MustCompile(`(?m)^\s*<!--\s*pause\s*-->\s*$`)
	columnRegex = regexp.MustCompile(`(?m)^\s*<!--\s*column:\s*\d+\s*-->\s*$`)
)

// ParsePresentation parses raw markdown content into a Presentation.
func ParsePresentation(content string) *model.Presentation {
	// Extract frontmatter BEFORE splitting slides, since frontmatter uses
	// the same --- delimiter as slide boundaries.
	fm, remaining := ParseFrontmatter(content)

	rawSlides := SplitSlides(remaining)
	if len(rawSlides) == 0 {
		return &model.Presentation{Frontmatter: fm}
	}

	slides := make([]model.Slide, 0, len(rawSlides))
	for _, raw := range rawSlides {
		slide := parseSlide(raw)
		slides = append(slides, slide)
	}

	return &model.Presentation{
		Slides:      slides,
		Frontmatter: fm,
	}
}

func parseSlide(raw string) model.Slide {
	var slide model.Slide

	// Extract speaker notes and layout commands first
	cmds, _ := ExtractCommands(raw)
	var layout *model.ColumnLayout
	for _, cmd := range cmds {
		switch cmd.Command.Type {
		case model.CmdSpeakerNote:
			slide.SpeakerNotes = append(slide.SpeakerNotes, cmd.Command.Value)
		case model.CmdColumnLayout:
			layout = &model.ColumnLayout{Ratios: cmd.Command.Ratios}
		}
	}
	slide.Layout = layout

	// Extract column content if layout is present
	if layout != nil {
		slide.Columns = extractColumns(raw, len(layout.Ratios))
	}

	// Split at pause markers to create chunks
	parts := pauseRegex.Split(raw, -1)
	for _, part := range parts {
		_, cleaned := extractNonPauseCommands(part)
		slide.Chunks = append(slide.Chunks, model.Chunk{Content: cleaned})
	}

	if len(slide.Chunks) == 0 {
		slide.Chunks = []model.Chunk{{Content: raw}}
	}

	return slide
}

// extractColumns splits slide content into per-column buckets using
// <!-- column: N --> markers. Content before the first column marker
// is discarded (it's typically just the layout command).
func extractColumns(raw string, numColumns int) []string {
	columns := make([]string, numColumns)

	// Split at column markers
	parts := columnRegex.Split(raw, -1)
	markers := columnRegex.FindAllString(raw, -1)

	if len(markers) == 0 {
		return nil
	}

	// Parse column indices from markers
	currentCol := -1
	for i, marker := range markers {
		cmds, _ := ExtractCommands(marker)
		for _, cmd := range cmds {
			if cmd.Command.Type == model.CmdColumn {
				currentCol = cmd.Command.Column
			}
		}

		// parts[i+1] is the content after this marker
		if currentCol >= 0 && currentCol < numColumns && i+1 < len(parts) {
			content := parts[i+1]
			// Strip non-column commands (speaker notes, layout, reset_layout, pauses)
			_, cleaned := extractNonPauseCommands(content)
			// Also strip pause markers
			cleaned = pauseRegex.ReplaceAllString(cleaned, "")
			columns[currentCol] += cleaned
		}
	}

	return columns
}

func extractNonPauseCommands(content string) ([]CommandWithPosition, string) {
	var commands []CommandWithPosition

	cleaned := commentRegex.ReplaceAllStringFunc(content, func(match string) string {
		sub := commentRegex.FindStringSubmatch(match)
		if len(sub) < 2 {
			return match
		}

		inner := strings.TrimSpace(sub[1])
		if inner == "pause" {
			return match
		}

		cmd, ok := parseCommand(inner)
		if !ok {
			return match
		}

		commands = append(commands, CommandWithPosition{Command: cmd})
		return ""
	})

	return commands, cleaned
}
