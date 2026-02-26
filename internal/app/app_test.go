package app

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
)

const testPresentation = `---
author: Tester
date: 2025
---
# Slide One

Hello world!
---
# Slide Two

<!-- pause -->
First reveal
<!-- pause -->
Second reveal
---
# Slide Three

` + "```bash\necho hello\n```"

func TestNewModel(t *testing.T) {
	m := New(testPresentation, "test.md")

	if m.presentation == nil {
		t.Fatal("presentation should not be nil")
	}
	if len(m.presentation.Slides) != 3 {
		t.Fatalf("expected 3 slides, got %d", len(m.presentation.Slides))
	}
	if m.presentation.Frontmatter.Author != "Tester" {
		t.Errorf("author = %q, want %q", m.presentation.Frontmatter.Author, "Tester")
	}
	if m.state.SlideIndex != 0 {
		t.Errorf("initial slide index = %d, want 0", m.state.SlideIndex)
	}
	if m.state.TotalSlides != 3 {
		t.Errorf("total slides = %d, want 3", m.state.TotalSlides)
	}
}

func TestModelInit(t *testing.T) {
	m := New(testPresentation, "")
	cmd := m.Init()
	if cmd != nil {
		t.Error("Init() should return nil")
	}
}

func TestModelUpdateWindowSize(t *testing.T) {
	m := New(testPresentation, "")
	newModel, _ := m.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	updated := newModel.(Model)

	if !updated.ready {
		t.Error("model should be ready after WindowSizeMsg")
	}
	if updated.width != 120 {
		t.Errorf("width = %d, want 120", updated.width)
	}
	if updated.height != 40 {
		t.Errorf("height = %d, want 40", updated.height)
	}
}

func TestModelNavigation(t *testing.T) {
	m := New(testPresentation, "")
	// Set ready + dimensions
	newModel, _ := m.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	m = newModel.(Model)

	// Navigate forward
	newModel, _ = m.Update(tea.KeyPressMsg(tea.Key{Code: 'l'}))
	m = newModel.(Model)
	if m.state.SlideIndex != 1 {
		t.Errorf("after 'l': slide = %d, want 1", m.state.SlideIndex)
	}

	// Navigate forward again — slide 2 has 3 chunks, should advance chunk first
	newModel, _ = m.Update(tea.KeyPressMsg(tea.Key{Code: 'l'}))
	m = newModel.(Model)
	if m.state.SlideIndex != 1 || m.state.ChunkIndex != 1 {
		t.Errorf("after second 'l': slide=%d chunk=%d, want slide=1 chunk=1",
			m.state.SlideIndex, m.state.ChunkIndex)
	}

	// Navigate backward
	newModel, _ = m.Update(tea.KeyPressMsg(tea.Key{Code: 'h'}))
	m = newModel.(Model)
	if m.state.ChunkIndex != 0 {
		t.Errorf("after 'h': chunk = %d, want 0", m.state.ChunkIndex)
	}
}

func TestModelQuit(t *testing.T) {
	m := New(testPresentation, "")
	_, cmd := m.Update(tea.KeyPressMsg(tea.Key{Code: 'q'}))
	if cmd == nil {
		t.Error("'q' should produce a quit command")
	}
}

func TestModelSearch(t *testing.T) {
	m := New(testPresentation, "")
	newModel, _ := m.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	m = newModel.(Model)

	// Enter search mode
	newModel, _ = m.Update(tea.KeyPressMsg(tea.Key{Code: '/'}))
	m = newModel.(Model)
	if !m.searching {
		t.Error("should be in search mode after '/'")
	}

	// Type search query
	for _, ch := range "Three" {
		newModel, _ = m.Update(tea.KeyPressMsg(tea.Key{Code: ch}))
		m = newModel.(Model)
	}
	if m.searchQuery != "Three" {
		t.Errorf("search query = %q, want %q", m.searchQuery, "Three")
	}

	// Press enter to execute search
	newModel, _ = m.Update(tea.KeyPressMsg(tea.Key{Code: tea.KeyEnter}))
	m = newModel.(Model)
	if m.searching {
		t.Error("should exit search mode after enter")
	}
	if m.state.SlideIndex != 2 {
		t.Errorf("should jump to slide 2 (Slide Three), got %d", m.state.SlideIndex)
	}
}

func TestModelFileChanged(t *testing.T) {
	m := New(testPresentation, "test.md")
	newModel, _ := m.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	m = newModel.(Model)

	// Navigate to slide 1
	newModel, _ = m.Update(tea.KeyPressMsg(tea.Key{Code: 'l'}))
	m = newModel.(Model)

	// Simulate file change — modify slide 2
	newContent := strings.Replace(testPresentation, "Second reveal", "MODIFIED reveal", 1)
	newModel, _ = m.Update(FileChangedMsg{Content: newContent})
	m = newModel.(Model)

	// Should jump to modified slide (slide 1, which has the pause content)
	if m.state.SlideIndex != 1 {
		t.Errorf("should jump to modified slide 1, got %d", m.state.SlideIndex)
	}
}

func TestModelView(t *testing.T) {
	m := New(testPresentation, "")
	// Before ready
	v := m.View()
	if !v.AltScreen {
		t.Error("should use alt screen")
	}

	// After ready
	newModel, _ := m.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	m = newModel.(Model)
	v = m.View()
	if !v.AltScreen {
		t.Error("should still use alt screen after ready")
	}
}

func TestModelEscClearsCodeOutput(t *testing.T) {
	m := New(testPresentation, "")
	newModel, _ := m.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	m = newModel.(Model)

	// Simulate code output
	newModel, _ = m.Update(CodeResultMsg{Output: "hello"})
	m = newModel.(Model)
	if m.codeOutput != "hello" {
		t.Errorf("code output = %q, want %q", m.codeOutput, "hello")
	}

	// Esc clears it
	newModel, _ = m.Update(tea.KeyPressMsg(tea.Key{Code: tea.KeyEscape}))
	m = newModel.(Model)
	if m.codeOutput != "" {
		t.Errorf("esc should clear code output, got %q", m.codeOutput)
	}
}
