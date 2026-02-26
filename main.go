package main

import (
	_ "embed"
	"fmt"
	"io"
	"os"

	tea "charm.land/bubbletea/v2"

	"github.com/jedwards1230/deck/internal/app"
	"github.com/jedwards1230/deck/internal/version"
	"github.com/jedwards1230/deck/internal/watch"
)

//go:embed tutorial.md
var tutorial string

func main() {
	// Handle version flag
	if len(os.Args) == 2 && (os.Args[1] == "--version" || os.Args[1] == "-v" || os.Args[1] == "version") {
		fmt.Println("deck", version.Info())
		return
	}

	content, filePath, err := loadContent()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	m := app.New(content, filePath)

	p := tea.NewProgram(m)

	// Start file watcher if reading from a file
	if filePath != "" {
		go watch.Watch(filePath, p)
	}

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}

func loadContent() (content string, filePath string, err error) {
	// Check for piped input
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", "", fmt.Errorf("reading stdin: %w", err)
		}
		return string(data), "", nil
	}

	// Check for file argument
	if len(os.Args) < 2 {
		return tutorial, "", nil
	}

	path := os.Args[1]
	data, err := os.ReadFile(path)
	if err != nil {
		return "", "", fmt.Errorf("reading %s: %w", path, err)
	}

	return string(data), path, nil
}
