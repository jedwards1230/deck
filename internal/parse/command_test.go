package parse

import (
	"testing"

	"github.com/jedwards1230/deck/internal/model"
)

func TestExtractCommands(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantCmds     []model.Command
		wantCleaned  string
	}{
		{
			name:  "pause command",
			input: "Before\n<!-- pause -->\nAfter",
			wantCmds: []model.Command{
				{Type: model.CmdPause},
			},
			wantCleaned: "Before\n\nAfter",
		},
		{
			name:  "speaker note extracted",
			input: "Content\n<!-- speaker_note: Remember to explain this -->",
			wantCmds: []model.Command{
				{Type: model.CmdSpeakerNote, Value: "Remember to explain this"},
			},
			wantCleaned: "Content\n",
		},
		{
			name:  "column layout with ratios",
			input: "<!-- column_layout: [3, 2] -->",
			wantCmds: []model.Command{
				{Type: model.CmdColumnLayout, Ratios: []int{3, 2}},
			},
			wantCleaned: "",
		},
		{
			name:  "column index",
			input: "<!-- column: 0 -->\nLeft column content",
			wantCmds: []model.Command{
				{Type: model.CmdColumn, Column: 0},
			},
			wantCleaned: "\nLeft column content",
		},
		{
			name:  "reset layout",
			input: "<!-- reset_layout -->\nFull width again",
			wantCmds: []model.Command{
				{Type: model.CmdResetLayout},
			},
			wantCleaned: "\nFull width again",
		},
		{
			name:         "unrecognized comment left in content",
			input:        "Before\n<!-- TODO: fix this -->\nAfter",
			wantCmds:     nil,
			wantCleaned:  "Before\n<!-- TODO: fix this -->\nAfter",
		},
		{
			name:  "multiple commands in one block",
			input: "<!-- column_layout: [1, 1] -->\n<!-- speaker_note: Two cols -->\nBody",
			wantCmds: []model.Command{
				{Type: model.CmdColumnLayout, Ratios: []int{1, 1}},
				{Type: model.CmdSpeakerNote, Value: "Two cols"},
			},
			wantCleaned: "\n\nBody",
		},
		{
			name:  "extra whitespace in commands",
			input: "<!--   pause   -->",
			wantCmds: []model.Command{
				{Type: model.CmdPause},
			},
			wantCleaned: "",
		},
		{
			name:  "speaker note with extra whitespace",
			input: "<!--   speaker_note:   lots of spaces   -->",
			wantCmds: []model.Command{
				{Type: model.CmdSpeakerNote, Value: "lots of spaces"},
			},
			wantCleaned: "",
		},
		{
			name:         "empty content",
			input:        "",
			wantCmds:     nil,
			wantCleaned:  "",
		},
		{
			name:  "column layout three ratios",
			input: "<!-- column_layout: [1, 2, 3] -->",
			wantCmds: []model.Command{
				{Type: model.CmdColumnLayout, Ratios: []int{1, 2, 3}},
			},
			wantCleaned: "",
		},
		{
			name:         "column with non-numeric index is not recognized",
			input:        "<!-- column: abc -->",
			wantCmds:     nil,
			wantCleaned:  "<!-- column: abc -->",
		},
		{
			name:         "column layout with zero ratio is rejected",
			input:        "<!-- column_layout: [0, 1] -->",
			wantCmds:     nil,
			wantCleaned:  "<!-- column_layout: [0, 1] -->",
		},
		{
			name:         "column layout with negative ratio is rejected",
			input:        "<!-- column_layout: [-1, 2] -->",
			wantCmds:     nil,
			wantCleaned:  "<!-- column_layout: [-1, 2] -->",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCmds, gotCleaned := ExtractCommands(tt.input)

			if len(gotCmds) != len(tt.wantCmds) {
				t.Fatalf("len(commands) = %d, want %d\ngot:  %+v",
					len(gotCmds), len(tt.wantCmds), gotCmds)
			}

			for i := range gotCmds {
				got := gotCmds[i].Command
				want := tt.wantCmds[i]

				if got.Type != want.Type {
					t.Errorf("cmd[%d].Type = %v, want %v", i, got.Type, want.Type)
				}
				if got.Value != want.Value {
					t.Errorf("cmd[%d].Value = %q, want %q", i, got.Value, want.Value)
				}
				if got.Column != want.Column {
					t.Errorf("cmd[%d].Column = %d, want %d", i, got.Column, want.Column)
				}
				if !intSliceEqual(got.Ratios, want.Ratios) {
					t.Errorf("cmd[%d].Ratios = %v, want %v", i, got.Ratios, want.Ratios)
				}
			}

			if gotCleaned != tt.wantCleaned {
				t.Errorf("cleaned = %q, want %q", gotCleaned, tt.wantCleaned)
			}
		})
	}
}

func TestParseRatios(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []int
	}{
		{name: "two values", input: "[3, 2]", want: []int{3, 2}},
		{name: "three values", input: "[1, 2, 3]", want: []int{1, 2, 3}},
		{name: "no spaces", input: "[3,2]", want: []int{3, 2}},
		{name: "single value", input: "[5]", want: []int{5}},
		{name: "empty brackets", input: "[]", want: nil},
		{name: "zero value rejected", input: "[0, 1]", want: nil},
		{name: "negative value rejected", input: "[-1, 2]", want: nil},
		{name: "non-numeric rejected", input: "[a, b]", want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseRatios(tt.input)
			if !intSliceEqual(got, tt.want) {
				t.Errorf("parseRatios(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func intSliceEqual(a, b []int) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
