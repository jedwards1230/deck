package model

// CommandType identifies the type of HTML comment command.
type CommandType int

const (
	CmdPause CommandType = iota
	CmdSpeakerNote
	CmdColumnLayout
	CmdColumn
	CmdResetLayout
)

// Command represents a parsed HTML comment command.
type Command struct {
	Type   CommandType
	Value  string // for speaker notes: the note text
	Ratios []int  // for column_layout: proportional widths
	Column int    // for column: the column index (0-based)
}
