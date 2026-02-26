package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/jedwards1230/deck/internal/code"
	"github.com/jedwards1230/deck/internal/diff"
	"github.com/jedwards1230/deck/internal/model"
	"github.com/jedwards1230/deck/internal/nav"
	"github.com/jedwards1230/deck/internal/parse"
	"github.com/jedwards1230/deck/internal/render"
	"github.com/jedwards1230/deck/internal/search"
)

// Model is the bubbletea model for the slide presenter.
type Model struct {
	presentation *model.Presentation
	state        nav.State
	cache        *render.RendererCache
	width        int
	height       int
	ready        bool
	filePath     string // empty if reading from stdin
	codeOutput   string // virtual text from code execution

	// search state
	searching   bool
	searchQuery string
	lastSearch  string // persisted for n/N repeat
}

// New creates a new Model from the given content.
func New(content string, filePath string) Model {
	isDark := lipgloss.HasDarkBackground(os.Stdin, os.Stdout)
	pres := parse.ParsePresentation(content)

	totalSlides := len(pres.Slides)
	chunksInSlide := 1
	if totalSlides > 0 && len(pres.Slides[0].Chunks) > 0 {
		chunksInSlide = len(pres.Slides[0].Chunks)
	}

	return Model{
		presentation: pres,
		state: nav.State{
			TotalSlides:   totalSlides,
			ChunksInSlide: chunksInSlide,
		},
		cache:    render.NewRendererCache(isDark),
		filePath: filePath,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.cache.Invalidate()
		m.ready = true
		return m, nil

	case tea.KeyPressMsg:
		return m.handleKeyPress(msg)

	case FileChangedMsg:
		return m.handleFileChanged(msg)

	case CodeResultMsg:
		if msg.Err != nil {
			m.codeOutput = fmt.Sprintf("Error: %v", msg.Err)
		} else {
			m.codeOutput = msg.Output
		}
		return m, nil
	}

	return m, nil
}

func (m Model) handleKeyPress(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	key := msg.String()

	// Search mode input takes priority
	if m.searching {
		return m.handleSearchInput(msg)
	}

	// Guard against empty presentations
	if len(m.presentation.Slides) == 0 {
		if key == "q" || key == "ctrl+c" {
			return m, tea.Quit
		}
		return m, nil
	}

	// Global keybindings
	switch key {
	case "q", "ctrl+c":
		return m, tea.Quit

	case "esc":
		m.codeOutput = ""
		return m, nil

	case "/":
		m.searching = true
		m.searchQuery = ""
		return m, nil

	case "ctrl+e":
		return m, m.executeCode()

	case "y":
		m.yankCode()
		return m, nil

	case "ctrl+n":
		// Search next
		if m.lastSearch != "" {
			result := search.SearchNext(m.presentation.Slides, m.lastSearch, m.state.SlideIndex)
			if result.Found {
				m.jumpToSlide(result.SlideIndex)
			}
		}
		return m, nil

	case "N":
		// Search previous
		if m.lastSearch != "" {
			result := search.SearchPrev(m.presentation.Slides, m.lastSearch, m.state.SlideIndex)
			if result.Found {
				m.jumpToSlide(result.SlideIndex)
			}
		}
		return m, nil
	}

	// Clear code output on navigation
	m.codeOutput = ""

	// Delegate to pure navigation function
	newState := nav.Navigate(m.state, key)

	// Clamp chunk index when navigating to a different slide
	if newState.SlideIndex != m.state.SlideIndex {
		if newState.SlideIndex >= 0 && newState.SlideIndex < len(m.presentation.Slides) {
			slide := m.presentation.Slides[newState.SlideIndex]
			newState.ChunksInSlide = len(slide.Chunks)
			if newState.ChunkIndex >= newState.ChunksInSlide {
				newState.ChunkIndex = newState.ChunksInSlide - 1
			}
		}
	}

	m.state = newState
	return m, nil
}

func (m *Model) jumpToSlide(idx int) {
	if idx < 0 || idx >= len(m.presentation.Slides) {
		return
	}
	m.state.SlideIndex = idx
	m.state.ChunkIndex = 0
	m.state.ChunksInSlide = len(m.presentation.Slides[idx].Chunks)
}

func (m Model) executeCode() tea.Cmd {
	if m.state.SlideIndex >= len(m.presentation.Slides) {
		return nil
	}

	slide := m.presentation.Slides[m.state.SlideIndex]
	content := slide.VisibleContent(m.state.ChunkIndex)
	blocks := code.ExtractBlocks(content)

	if len(blocks) == 0 {
		return nil
	}

	// Execute the last code block on the slide
	block := blocks[len(blocks)-1]

	return func() tea.Msg {
		output, err := code.Execute(block)
		return CodeResultMsg{Output: output, Err: err}
	}
}

func (m Model) yankCode() {
	if m.state.SlideIndex >= len(m.presentation.Slides) {
		return
	}

	slide := m.presentation.Slides[m.state.SlideIndex]
	content := slide.VisibleContent(m.state.ChunkIndex)
	blocks := code.ExtractBlocks(content)

	if len(blocks) == 0 {
		return
	}

	// Copy the last code block to clipboard
	block := blocks[len(blocks)-1]
	_ = clipboard.WriteAll(block.Code)
}

func (m Model) handleSearchInput(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	key := msg.String()

	switch key {
	case "enter":
		m.searching = false
		if m.searchQuery != "" {
			m.lastSearch = m.searchQuery
			result := search.Search(m.presentation.Slides, m.searchQuery, m.state.SlideIndex)
			if result.Found {
				m.jumpToSlide(result.SlideIndex)
			}
		}
		return m, nil
	case "esc":
		m.searching = false
		m.searchQuery = ""
		return m, nil
	case "backspace":
		if len(m.searchQuery) > 0 {
			m.searchQuery = m.searchQuery[:len(m.searchQuery)-1]
		}
		return m, nil
	default:
		if len(key) == 1 || key == "space" {
			if key == "space" {
				key = " "
			}
			m.searchQuery += key
		}
		return m, nil
	}
}

func (m Model) handleFileChanged(msg FileChangedMsg) (tea.Model, tea.Cmd) {
	newPres := parse.ParsePresentation(msg.Content)

	jumpTo := diff.FindModified(m.presentation, newPres)

	m.presentation = newPres
	m.state.TotalSlides = len(newPres.Slides)

	if jumpTo >= 0 && jumpTo < len(newPres.Slides) {
		m.state.SlideIndex = jumpTo
		m.state.ChunkIndex = 0
	}

	// Clamp slide index
	if m.state.SlideIndex >= m.state.TotalSlides {
		m.state.SlideIndex = m.state.TotalSlides - 1
	}
	if m.state.SlideIndex < 0 {
		m.state.SlideIndex = 0
	}

	// Update chunks in slide
	if m.state.SlideIndex < len(newPres.Slides) {
		m.state.ChunksInSlide = len(newPres.Slides[m.state.SlideIndex].Chunks)
		if m.state.ChunkIndex >= m.state.ChunksInSlide {
			m.state.ChunkIndex = m.state.ChunksInSlide - 1
		}
	}

	return m, nil
}

func (m Model) View() tea.View {
	var v tea.View
	v.AltScreen = true

	if !m.ready || len(m.presentation.Slides) == 0 {
		v.SetContent("\n  Loading...")
		return v
	}

	// Reserve space for footer (2 lines: divider + content)
	footerHeight := 2
	contentHeight := m.height - footerHeight

	// Render slide content
	slide := m.presentation.Slides[m.state.SlideIndex]
	rendered, _ := render.RenderSlide(slide, m.state.ChunkIndex, m.width, m.cache)

	// Append code output if present
	if m.codeOutput != "" {
		rendered += "\n" + m.codeOutput
	}

	// Pad content to fill available height
	renderedLines := strings.Count(rendered, "\n") + 1
	if renderedLines < contentHeight {
		rendered += strings.Repeat("\n", contentHeight-renderedLines)
	}

	// Render footer
	footer := render.RenderFooter(
		m.presentation.Frontmatter,
		m.state.SlideIndex,
		m.state.TotalSlides,
		m.width,
	)

	// Search bar overlay
	if m.searching {
		searchBar := fmt.Sprintf("/%s█", m.searchQuery)
		footer = searchBar + strings.Repeat(" ", max(0, m.width-lipgloss.Width(searchBar))) + "\n"
	}

	v.SetContent(rendered + footer)
	return v
}
