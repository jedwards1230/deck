package code

// Language defines how to execute a code block for a given language.
type Language struct {
	Extension string   // file extension (e.g., ".go")
	Command   []string // command template: <file>, <name>, <path> are replaced
}

// Languages maps language names to their execution configurations.
var Languages = map[string]Language{
	"go":         {Extension: ".go", Command: []string{"go", "run", "<file>"}},
	"bash":       {Extension: ".sh", Command: []string{"bash", "<file>"}},
	"sh":         {Extension: ".sh", Command: []string{"sh", "<file>"}},
	"python":     {Extension: ".py", Command: []string{"python3", "<file>"}},
	"python3":    {Extension: ".py", Command: []string{"python3", "<file>"}},
	"javascript": {Extension: ".js", Command: []string{"node", "<file>"}},
	"js":         {Extension: ".js", Command: []string{"node", "<file>"}},
	"ruby":       {Extension: ".rb", Command: []string{"ruby", "<file>"}},
	"rb":         {Extension: ".rb", Command: []string{"ruby", "<file>"}},
}
