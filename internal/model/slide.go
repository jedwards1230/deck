package model

// Slide represents a single slide in the presentation.
type Slide struct {
	Chunks       []Chunk
	Columns      []string       // per-column content (populated when Layout is set)
	SpeakerNotes []string
	Layout       *ColumnLayout // nil = full-width
}

// VisibleContent returns the concatenated content of chunks 0..chunkIndex.
func (s Slide) VisibleContent(chunkIndex int) string {
	if chunkIndex >= len(s.Chunks) {
		chunkIndex = len(s.Chunks) - 1
	}
	var content string
	for i := 0; i <= chunkIndex; i++ {
		content += s.Chunks[i].Content
	}
	return content
}
