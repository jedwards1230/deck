package parse

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jedwards1230/deck/internal/model"
)

var commentRegex = regexp.MustCompile(`<!--\s*(.*?)\s*-->`)

// ExtractCommands parses HTML comments from content, returning commands and cleaned content.
func ExtractCommands(content string) ([]CommandWithPosition, string) {
	var commands []CommandWithPosition

	cleaned := commentRegex.ReplaceAllStringFunc(content, func(match string) string {
		sub := commentRegex.FindStringSubmatch(match)
		if len(sub) < 2 {
			return match
		}

		inner := strings.TrimSpace(sub[1])
		cmd, ok := parseCommand(inner)
		if !ok {
			return match // not a recognized command, keep as-is
		}

		commands = append(commands, CommandWithPosition{Command: cmd})
		return "" // strip the command from content
	})

	return commands, cleaned
}

// CommandWithPosition pairs a command with its position in the content.
type CommandWithPosition struct {
	Command model.Command
}

func parseCommand(s string) (model.Command, bool) {
	switch {
	case s == "pause":
		return model.Command{Type: model.CmdPause}, true

	case strings.HasPrefix(s, "speaker_note:"):
		note := strings.TrimSpace(strings.TrimPrefix(s, "speaker_note:"))
		return model.Command{Type: model.CmdSpeakerNote, Value: note}, true

	case strings.HasPrefix(s, "column_layout:"):
		ratioStr := strings.TrimSpace(strings.TrimPrefix(s, "column_layout:"))
		ratios := parseRatios(ratioStr)
		if len(ratios) == 0 {
			return model.Command{}, false
		}
		return model.Command{Type: model.CmdColumnLayout, Ratios: ratios}, true

	case strings.HasPrefix(s, "column:"):
		colStr := strings.TrimSpace(strings.TrimPrefix(s, "column:"))
		col, err := strconv.Atoi(colStr)
		if err != nil {
			return model.Command{}, false
		}
		return model.Command{Type: model.CmdColumn, Column: col}, true

	case s == "reset_layout":
		return model.Command{Type: model.CmdResetLayout}, true

	default:
		return model.Command{}, false
	}
}

func parseRatios(s string) []int {
	// Parse [3,2] or [1, 2, 3] format
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "[]")
	parts := strings.Split(s, ",")

	var ratios []int
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		n, err := strconv.Atoi(p)
		if err != nil || n <= 0 {
			return nil
		}
		ratios = append(ratios, n)
	}

	return ratios
}
