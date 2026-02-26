package nav

// State tracks navigation position within a presentation.
type State struct {
	SlideIndex    int
	ChunkIndex    int
	TotalSlides   int
	ChunksInSlide int    // number of chunks in current slide
	Buffer        string // numeric prefix buffer for vim-style navigation
}
