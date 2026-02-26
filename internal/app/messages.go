package app

// FileChangedMsg signals that the watched file has been modified.
type FileChangedMsg struct {
	Content string
}

// CodeResultMsg carries the output of a code block execution.
type CodeResultMsg struct {
	Output string
	Err    error
}

