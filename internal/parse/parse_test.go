package parse

import (
	"strings"
	"testing"

	"github.com/jedwards1230/deck/internal/model"
)

func TestParsePresentation(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		wantSlideCount int
		wantFM         model.Frontmatter
		check          func(t *testing.T, p *model.Presentation)
	}{
		{
			name: "full presentation with frontmatter and slides",
			input: `---
author: Test Author
paging: "%d / %d"
---
# Title Slide

Welcome!
---
# Slide Two

Details here.
---
# Conclusion

The end.`,
			wantSlideCount: 3,
			wantFM: model.Frontmatter{
				Author: "Test Author",
				Paging: "%d / %d",
			},
			check: func(t *testing.T, p *model.Presentation) {
				t.Helper()

				s0Content := p.Slides[0].VisibleContent(len(p.Slides[0].Chunks) - 1)
				if !strings.Contains(s0Content, "Title Slide") {
					t.Errorf("slide 0 should contain 'Title Slide', got %q", s0Content)
				}

				s1Content := p.Slides[1].VisibleContent(len(p.Slides[1].Chunks) - 1)
				if !strings.Contains(s1Content, "Slide Two") {
					t.Errorf("slide 1 should contain 'Slide Two', got %q", s1Content)
				}

				s2Content := p.Slides[2].VisibleContent(len(p.Slides[2].Chunks) - 1)
				if !strings.Contains(s2Content, "Conclusion") {
					t.Errorf("slide 2 should contain 'Conclusion', got %q", s2Content)
				}
			},
		},
		{
			name: "pause commands create multiple chunks",
			input: `# Progressive Reveal

First point
<!-- pause -->
Second point
<!-- pause -->
Third point`,
			wantSlideCount: 1,
			wantFM:         model.Frontmatter{},
			check: func(t *testing.T, p *model.Presentation) {
				t.Helper()

				slide := p.Slides[0]
				if len(slide.Chunks) != 3 {
					t.Fatalf("expected 3 chunks, got %d", len(slide.Chunks))
				}

				if !strings.Contains(slide.Chunks[0].Content, "First point") {
					t.Errorf("chunk 0 should contain 'First point', got %q", slide.Chunks[0].Content)
				}

				if !strings.Contains(slide.Chunks[1].Content, "Second point") {
					t.Errorf("chunk 1 should contain 'Second point', got %q", slide.Chunks[1].Content)
				}

				if !strings.Contains(slide.Chunks[2].Content, "Third point") {
					t.Errorf("chunk 2 should contain 'Third point', got %q", slide.Chunks[2].Content)
				}

				v0 := slide.VisibleContent(0)
				if strings.Contains(v0, "Second point") {
					t.Errorf("VisibleContent(0) should not contain 'Second point'")
				}

				v1 := slide.VisibleContent(1)
				if !strings.Contains(v1, "First point") || !strings.Contains(v1, "Second point") {
					t.Errorf("VisibleContent(1) should contain first and second points, got %q", v1)
				}
			},
		},
		{
			name: "speaker notes extracted",
			input: `# My Slide

Content here.
<!-- speaker_note: Do not forget to mention X -->
<!-- speaker_note: Also mention Y -->`,
			wantSlideCount: 1,
			wantFM:         model.Frontmatter{},
			check: func(t *testing.T, p *model.Presentation) {
				t.Helper()

				slide := p.Slides[0]
				if len(slide.SpeakerNotes) != 2 {
					t.Fatalf("expected 2 speaker notes, got %d: %v",
						len(slide.SpeakerNotes), slide.SpeakerNotes)
				}
				if slide.SpeakerNotes[0] != "Do not forget to mention X" {
					t.Errorf("note[0] = %q, want %q",
						slide.SpeakerNotes[0], "Do not forget to mention X")
				}
				if slide.SpeakerNotes[1] != "Also mention Y" {
					t.Errorf("note[1] = %q, want %q",
						slide.SpeakerNotes[1], "Also mention Y")
				}

				visible := slide.VisibleContent(len(slide.Chunks) - 1)
				if strings.Contains(visible, "speaker_note") {
					t.Errorf("visible content should not contain speaker_note command, got %q", visible)
				}
			},
		},
		{
			name:           "empty content returns empty presentation",
			input:          "",
			wantSlideCount: 0,
			wantFM:         model.Frontmatter{},
			check:          nil,
		},
		{
			name: "slides with mixed commands",
			input: `---
author: Mixer
---
# Intro
---
<!-- column_layout: [1, 1] -->
<!-- speaker_note: Two column layout -->
<!-- column: 0 -->
Left side
<!-- pause -->
<!-- column: 1 -->
Right side
---
<!-- reset_layout -->
# Full Width Again`,
			wantSlideCount: 3,
			wantFM: model.Frontmatter{
				Author: "Mixer",
				Paging: "Slide %d / %d",
			},
			check: func(t *testing.T, p *model.Presentation) {
				t.Helper()

				// Slide 1 (index 1): should have column layout.
				slide1 := p.Slides[1]
				if slide1.Layout == nil {
					t.Fatal("slide 1 should have a column layout")
				}
				if !intSliceEqual(slide1.Layout.Ratios, []int{1, 1}) {
					t.Errorf("slide 1 layout ratios = %v, want [1, 1]", slide1.Layout.Ratios)
				}

				// Columns should be populated
				if len(slide1.Columns) != 2 {
					t.Fatalf("slide 1 expected 2 columns, got %d", len(slide1.Columns))
				}
				if !strings.Contains(slide1.Columns[0], "Left side") {
					t.Errorf("column 0 should contain 'Left side', got %q", slide1.Columns[0])
				}
				if !strings.Contains(slide1.Columns[1], "Right side") {
					t.Errorf("column 1 should contain 'Right side', got %q", slide1.Columns[1])
				}

				if len(slide1.SpeakerNotes) != 1 {
					t.Fatalf("slide 1 expected 1 speaker note, got %d", len(slide1.SpeakerNotes))
				}
				if slide1.SpeakerNotes[0] != "Two column layout" {
					t.Errorf("note = %q, want %q", slide1.SpeakerNotes[0], "Two column layout")
				}

				// Slide 1 should have 2 chunks (split by pause).
				if len(slide1.Chunks) != 2 {
					t.Fatalf("slide 1 expected 2 chunks, got %d", len(slide1.Chunks))
				}

				// Slide 2 (index 2): full width, no layout, no columns.
				slide2 := p.Slides[2]
				if slide2.Layout != nil {
					t.Errorf("slide 2 should not have a layout")
				}
				if len(slide2.Columns) != 0 {
					t.Errorf("slide 2 should not have columns, got %d", len(slide2.Columns))
				}
				visible := slide2.VisibleContent(len(slide2.Chunks) - 1)
				if !strings.Contains(visible, "Full Width Again") {
					t.Errorf("slide 2 should contain 'Full Width Again', got %q", visible)
				}
			},
		},
		{
			name: "frontmatter only no slides after",
			input: `---
author: minimal
---`,
			wantSlideCount: 0,
			wantFM: model.Frontmatter{
				Author: "minimal",
				Paging: "Slide %d / %d",
			},
			check: nil,
		},
		{
			name: "no frontmatter defaults paging to empty",
			input: `# Just Content

No frontmatter here.`,
			wantSlideCount: 1,
			wantFM: model.Frontmatter{
				Paging: "",
			},
			check: func(t *testing.T, p *model.Presentation) {
				t.Helper()

				visible := p.Slides[0].VisibleContent(0)
				if !strings.Contains(visible, "Just Content") {
					t.Errorf("slide should contain 'Just Content', got %q", visible)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ParsePresentation(tt.input)

			if len(p.Slides) != tt.wantSlideCount {
				t.Fatalf("slide count = %d, want %d", len(p.Slides), tt.wantSlideCount)
			}

			if p.Frontmatter.Author != tt.wantFM.Author {
				t.Errorf("Frontmatter.Author = %q, want %q", p.Frontmatter.Author, tt.wantFM.Author)
			}
			if p.Frontmatter.Paging != tt.wantFM.Paging {
				t.Errorf("Frontmatter.Paging = %q, want %q", p.Frontmatter.Paging, tt.wantFM.Paging)
			}

			if tt.check != nil {
				tt.check(t, p)
			}
		})
	}
}
