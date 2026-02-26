---
author: deck
date: 2025
paging: Slide %d / %d
permalink: services/deck/tutorial
---

# Welcome to Deck

A terminal slide presenter built with Bubbletea v2.

---

## Navigation

| Key | Action |
|-----|--------|
| `l` `space` `right` | Next slide |
| `h` `left` | Previous slide |
| `j` / `k` | Next / Previous |
| `gg` | First slide |
| `G` | Last slide |
| `3G` | Go to slide 3 |
| `/` | Search |
| `q` | Quit |

---

## Markdown Support

Deck renders **full markdown** via Glamour:

- **Bold** and *italic* text
- `inline code` and code blocks
- Lists, tables, and blockquotes

> "The terminal is a canvas." — Someone, probably

---

## Code Blocks

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello from deck!")
}
```

Press `ctrl+e` to execute code blocks.

---

## Progressive Reveal

Content can be revealed incrementally:

<!-- pause -->

First, this appears...

<!-- pause -->

Then this!

<!-- pause -->

And finally, the conclusion.

---

## Column Layouts

<!-- column_layout: [1, 1] -->
<!-- column: 0 -->

### Left Column

This content appears on the left side of the slide.

<!-- column: 1 -->

### Right Column

And this content appears on the right.

<!-- reset_layout -->

---

## Speaker Notes

Speaker notes are hidden from the audience:

<!-- speaker_note: Remember to mention the hot reload feature! -->

Notes are parsed and stored but not displayed.

---

## Hot Reload

Edit your slides file and deck will:

1. Detect the change instantly
2. Re-parse the presentation
3. Jump to the modified slide

No restart needed!

---

## Getting Started

```bash
# Present a file
deck slides.md

# Or pipe content
cat slides.md | deck
```

---

# Thank You

**deck** — presentations in the terminal

`github.com/jedwards1230/deck`