package nav

import (
	"math"
	"strconv"
)

// Navigate receives the current State and a key press, returns the new State.
// This is a pure function with no side effects.
func Navigate(state State, keyPress string) State {
	switch keyPress {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		newBuffer := keyPress
		if bufferIsNumeric(state.Buffer) {
			newBuffer = state.Buffer + keyPress
		}
		return State{
			SlideIndex:    state.SlideIndex,
			ChunkIndex:    state.ChunkIndex,
			TotalSlides:   state.TotalSlides,
			ChunksInSlide: state.ChunksInSlide,
			Buffer:        newBuffer,
		}

	case "g":
		if state.Buffer == "g" {
			return State{
				SlideIndex:    0,
				ChunkIndex:    0,
				TotalSlides:   state.TotalSlides,
				ChunksInSlide: state.ChunksInSlide,
			}
		}
		return State{
			SlideIndex:    state.SlideIndex,
			ChunkIndex:    state.ChunkIndex,
			TotalSlides:   state.TotalSlides,
			ChunksInSlide: state.ChunksInSlide,
			Buffer:        "g",
		}

	case "G":
		if bufferIsNumeric(state.Buffer) {
			target := navigateToSlide(state.Buffer, state.TotalSlides)
			return State{
				SlideIndex:    target,
				ChunkIndex:    0,
				TotalSlides:   state.TotalSlides,
				ChunksInSlide: state.ChunksInSlide,
			}
		}
		return State{
			SlideIndex:    state.TotalSlides - 1,
			ChunkIndex:    0,
			TotalSlides:   state.TotalSlides,
			ChunksInSlide: state.ChunksInSlide,
		}

	case "space", "down", "j", "right", "l", "enter", "n", "pgdown":
		return advanceForward(state)

	case "up", "k", "left", "h", "p", "pgup":
		return advanceBackward(state)

	default:
		// Unknown key clears buffer.
		return State{
			SlideIndex:    state.SlideIndex,
			ChunkIndex:    state.ChunkIndex,
			TotalSlides:   state.TotalSlides,
			ChunksInSlide: state.ChunksInSlide,
		}
	}
}

func advanceForward(state State) State {
	repeat := 1
	if bufferIsNumeric(state.Buffer) {
		r, _ := strconv.Atoi(state.Buffer)
		if r > 0 {
			repeat = r
		}
	}

	slide := state.SlideIndex
	chunk := state.ChunkIndex
	chunksInSlide := state.ChunksInSlide

	for range repeat {
		if chunk < chunksInSlide-1 {
			chunk++
		} else if slide < state.TotalSlides-1 {
			slide++
			chunk = 0
			chunksInSlide = 1 // will be corrected by caller
		}
	}

	return State{
		SlideIndex:    slide,
		ChunkIndex:    chunk,
		TotalSlides:   state.TotalSlides,
		ChunksInSlide: chunksInSlide,
	}
}

func advanceBackward(state State) State {
	repeat := 1
	if bufferIsNumeric(state.Buffer) {
		r, _ := strconv.Atoi(state.Buffer)
		if r > 0 {
			repeat = r
		}
	}

	slide := state.SlideIndex
	chunk := state.ChunkIndex

	for range repeat {
		if chunk > 0 {
			chunk--
		} else if slide > 0 {
			slide--
			chunk = math.MaxInt // sentinel: show all chunks of previous slide (caller clamps)
		}
	}

	return State{
		SlideIndex:    slide,
		ChunkIndex:    chunk,
		TotalSlides:   state.TotalSlides,
		ChunksInSlide: state.ChunksInSlide,
	}
}

func navigateToSlide(buffer string, totalSlides int) int {
	target, _ := strconv.Atoi(buffer)
	target-- // 1-indexed to 0-indexed

	if target < 0 {
		return 0
	}
	if target >= totalSlides {
		return totalSlides - 1
	}
	return target
}

func bufferIsNumeric(buffer string) bool {
	_, err := strconv.Atoi(buffer)
	return err == nil
}
