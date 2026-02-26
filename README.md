# deck

A terminal slide presenter built with [Bubbletea v2](https://charm.land/bubbletea).

Present markdown files as navigable slide decks in the terminal with live hot reload, progressive reveal, column layouts, code execution, and vim-style navigation.

## Install

```bash
go install github.com/jedwards1230/deck@latest
```

Or download a binary from [releases](https://github.com/jedwards1230/deck/releases).

## Usage

```bash
# Present a markdown file
deck slides.md

# Pipe content
cat slides.md | deck

# Built-in tutorial
deck
```

## Slide Format

Slides are separated by `---` on its own line. YAML frontmatter is optional:

```markdown
---
theme: dark
author: Your Name
date: 2025
paging: "%d / %d"
---

# First Slide

Content here.

---

# Second Slide

More content.
```

## Features

### Navigation

| Key | Action |
|-----|--------|
| `l` `space` `right` `enter` | Next |
| `h` `left` | Previous |
| `j` / `k` | Forward / Back |
| `gg` | First slide |
| `G` | Last slide |
| `3G` | Go to slide 3 |
| `/` | Search |
| `ctrl+n` / `N` | Next / previous match |
| `ctrl+e` | Execute code block |
| `y` | Copy code to clipboard |
| `q` | Quit |

### Progressive Reveal

```markdown
First point appears immediately.

<!-- pause -->

This appears on the next advance.

<!-- pause -->

And this on the advance after that.
```

### Column Layouts

```markdown
<!-- column_layout: [1, 1] -->
<!-- column: 0 -->

### Left

Left column content.

<!-- column: 1 -->

### Right

Right column content.

<!-- reset_layout -->
```

### Speaker Notes

```markdown
<!-- speaker_note: This is hidden from display. -->
```

### Hot Reload

When presenting a file, deck watches for changes and automatically jumps to the modified slide.

### Code Execution

Press `ctrl+e` to execute the last code block on the current slide. Supports Go, Bash, Python, JavaScript, and Ruby.

### Custom Footer

```yaml
---
footer: "{author} | {current_slide}/{total_slides}"
---
```

## License

MIT
