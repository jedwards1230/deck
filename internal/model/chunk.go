package model

// Chunk is a unit of progressive reveal (content between pause commands).
type Chunk struct {
	Content string // raw markdown (commands stripped)
}
