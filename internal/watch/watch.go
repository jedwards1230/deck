package watch

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/fsnotify/fsnotify"

	"github.com/jedwards1230/deck/internal/app"
)

// Watch monitors a file for changes and sends FileChangedMsg to the program.
func Watch(filePath string, p *tea.Program) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Fprintf(os.Stderr, "watch error: %v\n", err)
		return
	}
	defer func() { _ = watcher.Close() }()

	if err := watcher.Add(filePath); err != nil {
		fmt.Fprintf(os.Stderr, "watch error: %v\n", err)
		return
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Has(fsnotify.Write) {
				data, err := os.ReadFile(filePath)
				if err != nil {
					continue
				}
				p.Send(app.FileChangedMsg{Content: string(data)})
			}
		case _, ok := <-watcher.Errors:
			if !ok {
				return
			}
		}
	}
}
