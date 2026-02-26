package parse

import (
	"testing"

	"github.com/jedwards1230/deck/internal/model"
)

func TestParseFrontmatter(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		wantFM        model.Frontmatter
		wantRemaining string
	}{
		{
			name: "valid frontmatter with all fields",
			input: `---
author: Jane Doe
date: 2025-01-15
paging: Page %d of %d
footer: My Presentation
---
# Hello World`,
			wantFM: model.Frontmatter{
				Author: "Jane Doe",
				Date:   "2025-01-15",
				Paging: "Page %d of %d",
				Footer: "My Presentation",
			},
			wantRemaining: "# Hello World",
		},
		{
			name:  "content without frontmatter",
			input: "# Hello World\n\nSome body text.",
			wantFM: model.Frontmatter{
				// All zero values, no defaults applied.
			},
			wantRemaining: "# Hello World\n\nSome body text.",
		},
		{
			name:  "only opening delimiter no closing",
			input: "---\nauthor: Test\n# Slide content",
			wantFM: model.Frontmatter{
				// No closing ---, so frontmatter is not parsed.
			},
			wantRemaining: "---\nauthor: Test\n# Slide content",
		},
		{
			name:  "empty frontmatter block",
			input: "---\n---\n# Slide content",
			wantFM: model.Frontmatter{
				Paging: "Slide %d / %d", // default applied
			},
			wantRemaining: "# Slide content",
		},
		{
			name: "default paging when not specified",
			input: `---
author: Test
---
Body`,
			wantFM: model.Frontmatter{
				Author: "Test",
				Paging: "Slide %d / %d",
			},
			wantRemaining: "Body",
		},
		{
			name: "custom paging overrides default",
			input: `---
paging: "%d / %d"
---
Content`,
			wantFM: model.Frontmatter{
				Paging: "%d / %d",
			},
			wantRemaining: "Content",
		},
		{
			name: "leading whitespace before frontmatter",
			input: `  ---
author: Jane
---
After`,
			wantFM: model.Frontmatter{
				Author: "Jane",
				Paging: "Slide %d / %d",
			},
			wantRemaining: "After",
		},
		{
			name:  "empty content",
			input: "",
			wantFM: model.Frontmatter{
				// No frontmatter, no defaults.
			},
			wantRemaining: "",
		},
		{
			name: "frontmatter with no remaining content",
			input: `---
author: Test
---`,
			wantFM: model.Frontmatter{
				Author: "Test",
				Paging: "Slide %d / %d",
			},
			wantRemaining: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFM, gotRemaining := ParseFrontmatter(tt.input)

			if gotFM.Author != tt.wantFM.Author {
				t.Errorf("Author = %q, want %q", gotFM.Author, tt.wantFM.Author)
			}
			if gotFM.Date != tt.wantFM.Date {
				t.Errorf("Date = %q, want %q", gotFM.Date, tt.wantFM.Date)
			}
			if gotFM.Paging != tt.wantFM.Paging {
				t.Errorf("Paging = %q, want %q", gotFM.Paging, tt.wantFM.Paging)
			}
			if gotFM.Footer != tt.wantFM.Footer {
				t.Errorf("Footer = %q, want %q", gotFM.Footer, tt.wantFM.Footer)
			}
			if gotRemaining != tt.wantRemaining {
				t.Errorf("remaining = %q, want %q", gotRemaining, tt.wantRemaining)
			}
		})
	}
}
