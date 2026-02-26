package nav

import (
	"fmt"
	"math"
	"testing"
)

func TestNavigate(t *testing.T) {
	tests := []struct {
		name     string
		state    State
		keyPress string
		want     State
	}{
		// -----------------------------------------------------------
		// 1. Basic forward navigation
		// -----------------------------------------------------------
		{
			name:     "space advances forward one slide",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "space",
			want:     State{SlideIndex: 1, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "j advances forward one slide",
			state:    State{SlideIndex: 1, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "j",
			want:     State{SlideIndex: 2, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "down advances forward one slide",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "down",
			want:     State{SlideIndex: 1, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "l advances forward one slide",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "l",
			want:     State{SlideIndex: 1, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "right advances forward one slide",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "right",
			want:     State{SlideIndex: 1, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "enter advances forward one slide",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "enter",
			want:     State{SlideIndex: 1, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "n advances forward one slide",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "n",
			want:     State{SlideIndex: 1, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "pgdown advances forward one slide",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "pgdown",
			want:     State{SlideIndex: 1, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
		},

		// -----------------------------------------------------------
		// 2. Basic backward navigation
		// -----------------------------------------------------------
		{
			name:     "up goes backward one slide",
			state:    State{SlideIndex: 3, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "up",
			want:     State{SlideIndex: 2, ChunkIndex: math.MaxInt, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "k goes backward one slide",
			state:    State{SlideIndex: 2, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "k",
			want:     State{SlideIndex: 1, ChunkIndex: math.MaxInt, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "left goes backward one slide",
			state:    State{SlideIndex: 2, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "left",
			want:     State{SlideIndex: 1, ChunkIndex: math.MaxInt, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "h goes backward one slide",
			state:    State{SlideIndex: 2, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "h",
			want:     State{SlideIndex: 1, ChunkIndex: math.MaxInt, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "p goes backward one slide",
			state:    State{SlideIndex: 2, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "p",
			want:     State{SlideIndex: 1, ChunkIndex: math.MaxInt, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "pgup goes backward one slide",
			state:    State{SlideIndex: 2, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "pgup",
			want:     State{SlideIndex: 1, ChunkIndex: math.MaxInt, TotalSlides: 5, ChunksInSlide: 1},
		},

		// -----------------------------------------------------------
		// 3. Boundary clamping
		// -----------------------------------------------------------
		{
			name:     "forward at last slide stays clamped",
			state:    State{SlideIndex: 4, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "j",
			want:     State{SlideIndex: 4, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "backward at first slide stays clamped",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "k",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "forward at last slide last chunk stays clamped",
			state:    State{SlideIndex: 4, ChunkIndex: 2, TotalSlides: 5, ChunksInSlide: 3},
			keyPress: "j",
			want:     State{SlideIndex: 4, ChunkIndex: 2, TotalSlides: 5, ChunksInSlide: 3},
		},

		// -----------------------------------------------------------
		// 4. Numeric buffer building
		// -----------------------------------------------------------
		{
			name:     "pressing 1 sets buffer to 1",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "1",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1, Buffer: "1"},
		},
		{
			name:     "pressing 5 sets buffer to 5",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
			keyPress: "5",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "5"},
		},
		{
			name:     "pressing 0 on empty buffer sets buffer to 0",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "0",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1, Buffer: "0"},
		},
		{
			name:     "pressing 2 after 1 builds buffer 12",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 20, ChunksInSlide: 1, Buffer: "1"},
			keyPress: "2",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 20, ChunksInSlide: 1, Buffer: "12"},
		},
		{
			name:     "pressing 0 after 1 builds buffer 10",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 20, ChunksInSlide: 1, Buffer: "1"},
			keyPress: "0",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 20, ChunksInSlide: 1, Buffer: "10"},
		},
		{
			name:     "pressing 3 after non-numeric buffer g starts fresh buffer",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 20, ChunksInSlide: 1, Buffer: "g"},
			keyPress: "3",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 20, ChunksInSlide: 1, Buffer: "3"},
		},
		{
			name:     "pressing 9 with three-digit buffer appends",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 1000, ChunksInSlide: 1, Buffer: "12"},
			keyPress: "9",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 1000, ChunksInSlide: 1, Buffer: "129"},
		},

		// -----------------------------------------------------------
		// 5. Repeatable forward (e.g. 3j)
		// -----------------------------------------------------------
		{
			name:     "3j moves forward 3 slides",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "3"},
			keyPress: "j",
			want:     State{SlideIndex: 3, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:     "5j from slide 3 moves to slide 8",
			state:    State{SlideIndex: 3, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "5"},
			keyPress: "j",
			want:     State{SlideIndex: 8, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:     "repeat forward clamps at last slide",
			state:    State{SlideIndex: 7, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "5"},
			keyPress: "j",
			want:     State{SlideIndex: 9, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:     "repeat forward clears buffer",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "2"},
			keyPress: "space",
			want:     State{SlideIndex: 2, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},

		// -----------------------------------------------------------
		// 6. Repeatable backward (e.g. 2k)
		// -----------------------------------------------------------
		{
			name:     "2k from chunk 0 goes back one slide then decrements sentinel",
			state:    State{SlideIndex: 5, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "2"},
			keyPress: "k",
			// iter1: chunk=0 -> slide boundary -> slide=4, chunk=MaxInt
			// iter2: chunk=MaxInt>0 -> chunk=MaxInt-1
			want: State{SlideIndex: 4, ChunkIndex: math.MaxInt - 1, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:     "3k from slide 1 clamps slide at 0 and decrements sentinel",
			state:    State{SlideIndex: 1, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "3"},
			keyPress: "k",
			// iter1: chunk=0 -> slide boundary -> slide=0, chunk=MaxInt
			// iter2: chunk=MaxInt>0 -> chunk=MaxInt-1 (slide already 0, can't go further back)
			// iter3: chunk=MaxInt-1>0 -> chunk=MaxInt-2
			want: State{SlideIndex: 0, ChunkIndex: math.MaxInt - 2, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:     "repeat backward clears buffer and applies sentinel logic",
			state:    State{SlideIndex: 4, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "2"},
			keyPress: "up",
			// iter1: chunk=0 -> slide boundary -> slide=3, chunk=MaxInt
			// iter2: chunk=MaxInt>0 -> chunk=MaxInt-1
			want: State{SlideIndex: 3, ChunkIndex: math.MaxInt - 1, TotalSlides: 10, ChunksInSlide: 1},
		},

		// -----------------------------------------------------------
		// 7. gg - go to first slide
		// -----------------------------------------------------------
		{
			name:     "gg from middle goes to first slide",
			state:    State{SlideIndex: 5, ChunkIndex: 2, TotalSlides: 10, ChunksInSlide: 3, Buffer: "g"},
			keyPress: "g",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 3},
		},
		{
			name:     "gg from first slide is no-op",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "g"},
			keyPress: "g",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:     "single g sets buffer to g",
			state:    State{SlideIndex: 3, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
			keyPress: "g",
			want:     State{SlideIndex: 3, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "g"},
		},
		{
			name:     "g with numeric buffer sets buffer to g (not gg)",
			state:    State{SlideIndex: 3, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "5"},
			keyPress: "g",
			want:     State{SlideIndex: 3, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "g"},
		},

		// -----------------------------------------------------------
		// 8. G - go to last slide
		// -----------------------------------------------------------
		{
			name:     "G goes to last slide",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
			keyPress: "G",
			want:     State{SlideIndex: 9, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:     "G from last slide is no-op",
			state:    State{SlideIndex: 9, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
			keyPress: "G",
			want:     State{SlideIndex: 9, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:     "G with non-numeric buffer goes to last slide",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "g"},
			keyPress: "G",
			want:     State{SlideIndex: 9, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},

		// -----------------------------------------------------------
		// 9. 3G - go to slide 3 (1-indexed)
		// -----------------------------------------------------------
		{
			name:     "3G goes to slide index 2 (1-indexed input)",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "3"},
			keyPress: "G",
			want:     State{SlideIndex: 2, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:     "1G goes to first slide",
			state:    State{SlideIndex: 5, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "1"},
			keyPress: "G",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:     "10G goes to last slide in 10-slide deck",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "10"},
			keyPress: "G",
			want:     State{SlideIndex: 9, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:     "5G resets chunk to 0",
			state:    State{SlideIndex: 0, ChunkIndex: 2, TotalSlides: 10, ChunksInSlide: 3, Buffer: "5"},
			keyPress: "G",
			want:     State{SlideIndex: 4, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 3},
		},

		// -----------------------------------------------------------
		// 10. G with out-of-range number
		// -----------------------------------------------------------
		{
			name:     "0G clamps to first slide",
			state:    State{SlideIndex: 5, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "0"},
			keyPress: "G",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:     "99G clamps to last slide",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "99"},
			keyPress: "G",
			want:     State{SlideIndex: 9, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:     "11G clamps to last slide in 10-slide deck",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "11"},
			keyPress: "G",
			want:     State{SlideIndex: 9, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},

		// -----------------------------------------------------------
		// 11. Unknown key clears buffer
		// -----------------------------------------------------------
		{
			name:     "unknown key clears numeric buffer",
			state:    State{SlideIndex: 3, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "5"},
			keyPress: "x",
			want:     State{SlideIndex: 3, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:     "unknown key clears g buffer",
			state:    State{SlideIndex: 3, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1, Buffer: "g"},
			keyPress: "x",
			want:     State{SlideIndex: 3, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:     "unknown key preserves position",
			state:    State{SlideIndex: 7, ChunkIndex: 2, TotalSlides: 10, ChunksInSlide: 4},
			keyPress: "z",
			want:     State{SlideIndex: 7, ChunkIndex: 2, TotalSlides: 10, ChunksInSlide: 4},
		},

		// -----------------------------------------------------------
		// 12. Chunk-aware forward navigation
		// -----------------------------------------------------------
		{
			name:     "forward increments chunk when not at last chunk",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 3},
			keyPress: "j",
			want:     State{SlideIndex: 0, ChunkIndex: 1, TotalSlides: 5, ChunksInSlide: 3},
		},
		{
			name:     "forward increments chunk from 1 to 2",
			state:    State{SlideIndex: 0, ChunkIndex: 1, TotalSlides: 5, ChunksInSlide: 3},
			keyPress: "j",
			want:     State{SlideIndex: 0, ChunkIndex: 2, TotalSlides: 5, ChunksInSlide: 3},
		},
		{
			name:     "forward at last chunk advances slide and resets chunk",
			state:    State{SlideIndex: 0, ChunkIndex: 2, TotalSlides: 5, ChunksInSlide: 3},
			keyPress: "j",
			want:     State{SlideIndex: 1, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "repeat forward walks through chunks then slides",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 3, Buffer: "4"},
			keyPress: "j",
			// chunk 0->1, 1->2, 2->last so slide 0->1 chunk=0, then slide 1->2 chunk=0
			want: State{SlideIndex: 2, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "repeat forward 2 advances two chunks within same slide",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 4, Buffer: "2"},
			keyPress: "j",
			want:     State{SlideIndex: 0, ChunkIndex: 2, TotalSlides: 5, ChunksInSlide: 4},
		},

		// -----------------------------------------------------------
		// 13. Chunk-aware backward navigation
		// -----------------------------------------------------------
		{
			name:     "backward decrements chunk when not at first chunk",
			state:    State{SlideIndex: 2, ChunkIndex: 2, TotalSlides: 5, ChunksInSlide: 3},
			keyPress: "k",
			want:     State{SlideIndex: 2, ChunkIndex: 1, TotalSlides: 5, ChunksInSlide: 3},
		},
		{
			name:     "backward from chunk 1 to chunk 0",
			state:    State{SlideIndex: 2, ChunkIndex: 1, TotalSlides: 5, ChunksInSlide: 3},
			keyPress: "k",
			want:     State{SlideIndex: 2, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 3},
		},
		{
			name:     "backward at chunk 0 goes to previous slide with sentinel MaxInt",
			state:    State{SlideIndex: 2, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 3},
			keyPress: "k",
			want:     State{SlideIndex: 1, ChunkIndex: math.MaxInt, TotalSlides: 5, ChunksInSlide: 3},
		},
		{
			name:     "repeat backward walks through chunks then slides",
			state:    State{SlideIndex: 2, ChunkIndex: 2, TotalSlides: 5, ChunksInSlide: 3, Buffer: "3"},
			keyPress: "k",
			// chunk 2->1, 1->0, 0->prev slide: slide=1, chunk=MaxInt
			want: State{SlideIndex: 1, ChunkIndex: math.MaxInt, TotalSlides: 5, ChunksInSlide: 3},
		},
		{
			name:     "repeat backward 2 decrements two chunks within same slide",
			state:    State{SlideIndex: 2, ChunkIndex: 3, TotalSlides: 5, ChunksInSlide: 5, Buffer: "2"},
			keyPress: "k",
			want:     State{SlideIndex: 2, ChunkIndex: 1, TotalSlides: 5, ChunksInSlide: 5},
		},

		// -----------------------------------------------------------
		// 14. Edge case: single slide
		// -----------------------------------------------------------
		{
			name:     "forward on single slide single chunk is no-op",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 1, ChunksInSlide: 1},
			keyPress: "j",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 1, ChunksInSlide: 1},
		},
		{
			name:     "backward on single slide single chunk is no-op",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 1, ChunksInSlide: 1},
			keyPress: "k",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 1, ChunksInSlide: 1},
		},
		{
			name:     "G on single slide goes to slide 0",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 1, ChunksInSlide: 1},
			keyPress: "G",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 1, ChunksInSlide: 1},
		},
		{
			name:     "gg on single slide goes to slide 0",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 1, ChunksInSlide: 1, Buffer: "g"},
			keyPress: "g",
			want:     State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 1, ChunksInSlide: 1},
		},
		{
			name:     "single slide with multiple chunks forward increments chunk",
			state:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 1, ChunksInSlide: 3},
			keyPress: "j",
			want:     State{SlideIndex: 0, ChunkIndex: 1, TotalSlides: 1, ChunksInSlide: 3},
		},
		{
			name:     "single slide at last chunk forward is no-op",
			state:    State{SlideIndex: 0, ChunkIndex: 2, TotalSlides: 1, ChunksInSlide: 3},
			keyPress: "j",
			want:     State{SlideIndex: 0, ChunkIndex: 2, TotalSlides: 1, ChunksInSlide: 3},
		},

		// -----------------------------------------------------------
		// 15. Edge case: single chunk per slide
		// -----------------------------------------------------------
		{
			name:     "forward with single chunk advances slide directly",
			state:    State{SlideIndex: 2, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "j",
			want:     State{SlideIndex: 3, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
		},
		{
			name:     "backward with single chunk goes to previous slide sentinel",
			state:    State{SlideIndex: 2, ChunkIndex: 0, TotalSlides: 5, ChunksInSlide: 1},
			keyPress: "k",
			want:     State{SlideIndex: 1, ChunkIndex: math.MaxInt, TotalSlides: 5, ChunksInSlide: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Navigate(tt.state, tt.keyPress)
			if got != tt.want {
				t.Errorf("Navigate(%+v, %q)\n  got  = %+v\n  want = %+v", tt.state, tt.keyPress, got, tt.want)
			}
		})
	}
}

// TestNavigateSequence verifies multi-step navigation sequences to ensure
// state transitions compose correctly across multiple key presses.
func TestNavigateSequence(t *testing.T) {
	tests := []struct {
		name     string
		initial  State
		keys     []string
		want     State
	}{
		{
			name:    "type 1 then 2 then G goes to slide 12",
			initial: State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 20, ChunksInSlide: 1},
			keys:    []string{"1", "2", "G"},
			want:    State{SlideIndex: 11, ChunkIndex: 0, TotalSlides: 20, ChunksInSlide: 1},
		},
		{
			name:    "type g then g performs gg to first slide",
			initial: State{SlideIndex: 8, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
			keys:    []string{"g", "g"},
			want:    State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:    "forward then backward returns to original position with sentinel chunk",
			initial: State{SlideIndex: 3, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
			keys:    []string{"j", "k"},
			// j: 3->4 chunk=0, k: 4->3 chunk=MaxInt (sentinel for caller to clamp)
			want: State{SlideIndex: 3, ChunkIndex: math.MaxInt, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:    "numeric buffer then unknown key then forward moves 1",
			initial: State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
			keys:    []string{"5", "x", "j"},
			// 5: buffer="5", x: clears buffer, j: forward 1
			want: State{SlideIndex: 1, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
		{
			name:    "traverse all chunks in a multi-chunk slide",
			initial: State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 3, ChunksInSlide: 3},
			keys:    []string{"j", "j", "j"},
			// j: chunk 0->1, j: chunk 1->2, j: chunk 2 is last, advance slide to 1
			want: State{SlideIndex: 1, ChunkIndex: 0, TotalSlides: 3, ChunksInSlide: 1},
		},
		{
			name:    "G after g uses g as non-numeric buffer so G goes to last",
			initial: State{SlideIndex: 0, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
			keys:    []string{"g", "G"},
			// g sets buffer to "g", G sees non-numeric buffer => go to last slide
			want: State{SlideIndex: 9, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := tt.initial
			for _, key := range tt.keys {
				state = Navigate(state, key)
			}
			if state != tt.want {
				t.Errorf("after keys %v from %+v\n  got  = %+v\n  want = %+v", tt.keys, tt.initial, state, tt.want)
			}
		})
	}
}

// TestAllForwardKeysEquivalent verifies that all forward-navigation keys
// produce identical results from the same starting state.
func TestAllForwardKeysEquivalent(t *testing.T) {
	forwardKeys := []string{"space", "down", "j", "right", "l", "enter", "n", "pgdown"}
	initial := State{SlideIndex: 2, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1}
	expected := Navigate(initial, forwardKeys[0])

	for _, key := range forwardKeys[1:] {
		t.Run(key, func(t *testing.T) {
			got := Navigate(initial, key)
			if got != expected {
				t.Errorf("key %q produced %+v, want %+v (same as %q)", key, got, expected, forwardKeys[0])
			}
		})
	}
}

// TestAllBackwardKeysEquivalent verifies that all backward-navigation keys
// produce identical results from the same starting state.
func TestAllBackwardKeysEquivalent(t *testing.T) {
	backwardKeys := []string{"up", "k", "left", "h", "p", "pgup"}
	initial := State{SlideIndex: 5, ChunkIndex: 0, TotalSlides: 10, ChunksInSlide: 1}
	expected := Navigate(initial, backwardKeys[0])

	for _, key := range backwardKeys[1:] {
		t.Run(key, func(t *testing.T) {
			got := Navigate(initial, key)
			if got != expected {
				t.Errorf("key %q produced %+v, want %+v (same as %q)", key, got, expected, backwardKeys[0])
			}
		})
	}
}

// TestBufferIsNumeric verifies the internal bufferIsNumeric helper via
// observable Navigate behavior (digit appending vs. fresh buffer).
func TestBufferIsNumeric(t *testing.T) {
	tests := []struct {
		name           string
		existingBuffer string
		digit          string
		wantBuffer     string
	}{
		{"empty buffer starts fresh", "", "7", "7"},
		{"numeric buffer appends", "3", "4", "34"},
		{"multi-digit numeric buffer appends", "12", "3", "123"},
		{"g buffer starts fresh", "g", "5", "5"},
		{"zero buffer appends", "0", "1", "01"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := State{
				SlideIndex:    0,
				ChunkIndex:    0,
				TotalSlides:   100,
				ChunksInSlide: 1,
				Buffer:        tt.existingBuffer,
			}
			got := Navigate(state, tt.digit)
			if got.Buffer != tt.wantBuffer {
				t.Errorf("Navigate with buffer=%q, digit=%q: got buffer %q, want %q",
					tt.existingBuffer, tt.digit, got.Buffer, tt.wantBuffer)
			}
		})
	}
}

// TestNavigateToSlide verifies the navigateToSlide helper (1-indexed clamping).
func TestNavigateToSlide(t *testing.T) {
	tests := []struct {
		buffer      string
		totalSlides int
		want        int
	}{
		{"1", 10, 0},   // 1-indexed: slide 1 -> index 0
		{"5", 10, 4},   // slide 5 -> index 4
		{"10", 10, 9},  // slide 10 -> index 9
		{"0", 10, 0},   // 0 -> clamps to 0
		{"11", 10, 9},  // out of range -> clamps to last
		{"99", 5, 4},   // way out of range -> clamps to last
		{"1", 1, 0},    // single slide
		{"3", 3, 2},    // last slide exactly
	}

	for _, tt := range tests {
		name := fmt.Sprintf("buffer=%s_total=%d", tt.buffer, tt.totalSlides)
		t.Run(name, func(t *testing.T) {
			got := navigateToSlide(tt.buffer, tt.totalSlides)
			if got != tt.want {
				t.Errorf("navigateToSlide(%q, %d) = %d, want %d", tt.buffer, tt.totalSlides, got, tt.want)
			}
		})
	}
}
