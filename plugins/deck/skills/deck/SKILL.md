---
name: deck
description: Use when creating, editing, or improving terminal slide presentations
  with the deck CLI. Trigger phrases include "create a presentation", "make slides",
  "write a deck", "deck presentation", "make me a slide deck", "add slides", "edit
  my slides", "add progressive reveal", "add speaker notes", "column layout", "fix
  my deck". Covers the full deck markdown format, all comment directives, frontmatter
  options, and opinionated guidance on presentation structure.
---

# deck Skill

`deck` is a terminal slide presenter. Slides are written in Markdown, separated by `---`, and rendered in the terminal via `deck slides.md`.

## Quick Start

```bash
# Install
go install github.com/jedwards1230/deck@latest

# Present a file
deck slides.md

# Pipe content
cat slides.md | deck

# Built-in tutorial
deck
```

## File Format

A deck file is a Markdown file. Slides are separated by `---` on its own line. An optional YAML frontmatter block at the top of the file sets document-level metadata.

```markdown
---
author: Jane Smith
date: 2025
title: My Presentation
paging: "%d / %d"
footer: "{author} | {current_slide}/{total_slides}"
---

# First Slide

Content here.

---

# Second Slide

More content.
```

### Frontmatter Fields

| Field | Description | Example |
|-------|-------------|---------|
| `author` | Author name (usable in footer) | `author: Jane Smith` |
| `date` | Date or year | `date: 2025` |
| `title` | Presentation title | `title: My Talk` |
| `paging` | Slide counter format (`%d / %d`) | `paging: "%d / %d"` |
| `footer` | Footer template string | `footer: "{author} \| {current_slide}/{total_slides}"` |

### Footer Template Variables

- `{author}` — from frontmatter `author`
- `{current_slide}` — current slide number
- `{total_slides}` — total slide count

## Comment Directives

All directives are HTML comments and must be on their own line.

### Progressive Reveal — `<!-- pause -->`

Splits a slide's content into steps. Each advance reveals the next section. Use to build up arguments, lists, or diagrams incrementally.

```markdown
# Key Points

First point is visible immediately.

<!-- pause -->

Second point appears on next advance.

<!-- pause -->

Third point appears after that.
```

**When to use**: Complex arguments, step-by-step processes, numbered lists you want to walk through, before/after comparisons.

**When to avoid**: Simple slides with 1-2 points, visual layouts (column pauses get complex).

### Column Layouts

Split a slide into side-by-side columns. Columns are defined by a width ratio array.

```markdown
# Comparison

<!-- column_layout: [1, 1] -->

<!-- column: 0 -->

### Option A

- Fast
- Simple
- Well understood

<!-- column: 1 -->

### Option B

- Flexible
- Extensible
- More complex

<!-- reset_layout -->
```

**Directives:**
- `<!-- column_layout: [w, w, ...] -->` — define columns with relative widths (e.g. `[1, 1]` for equal, `[2, 1]` for 2:1)
- `<!-- column: N -->` — switch to column N (zero-indexed)
- `<!-- reset_layout -->` — end the column layout, return to full width

**When to use**: Side-by-side comparisons, pros/cons, before/after code, two concepts that benefit from visual separation.

**Common ratios:**
- `[1, 1]` — equal halves
- `[2, 1]` — wider left, narrow right (e.g. code + callout)
- `[1, 2]` — narrow left, wider right
- `[1, 1, 1]` — three equal columns

### Speaker Notes — `<!-- speaker_note: ... -->`

Hidden from the rendered slide. For presenter reminders, talking points, or time cues.

```markdown
# Architecture Overview

Key components: ingress, services, storage.

<!-- speaker_note: Mention the migration from NFS to Longhorn here. Expect questions about failover. -->
```

## Navigation Keys

For reference when presenting:

| Key | Action |
|-----|--------|
| `l` `space` `→` `enter` | Next |
| `h` `←` | Previous |
| `j` / `k` | Forward / back |
| `gg` | First slide |
| `G` | Last slide |
| `3G` | Jump to slide 3 |
| `/` | Search |
| `ctrl+n` / `N` | Next / previous search match |
| `ctrl+e` | Execute code block |
| `y` | Copy code to clipboard |
| `q` | Quit |

## Code Execution

Code blocks can be executed in-presentation with `ctrl+e`. Supported languages: Go, Bash, Python, JavaScript, Ruby.

```markdown
# Live Demo

\```bash
echo "Hello from deck!"
\```
```

## Hot Reload

When presenting a file, deck watches for changes and jumps to the modified slide automatically. Edit while presenting to iterate live.

---

## Presentation Design Guidance

### Structure

A good deck presentation follows a clear arc:

1. **Title slide** — topic, author, date
2. **Agenda / overview** — what you'll cover (optional for short talks)
3. **Content slides** — one idea per slide
4. **Summary / conclusion** — key takeaways
5. **Q&A or contact** — closing slide

**Ideal slide count**: 10–20 slides for a 20-minute talk. Fewer, denser slides are better than many thin ones.

### Content Density

- **One idea per slide** — if you find yourself using `<!-- pause -->` more than 3 times on a slide, it should probably be two slides.
- **Short bullets** — 4–6 words per bullet. No full sentences in lists.
- **Code blocks** — show only the relevant part. Full files belong in a repo, not a slide.
- **Headers as titles** — every slide should have an `# H1` or `## H2` header.

### When to Use Each Feature

| Feature | Use When |
|---------|----------|
| `<!-- pause -->` | Walking through steps, building an argument, numbered sequences |
| Column layout | Side-by-side comparison, before/after, two equal concepts |
| Speaker notes | Timing cues, stats to cite, anticipated questions |
| Code blocks | Live demos, showing syntax, before/after refactors |
| Footer | Multi-section talks, conference slides, when branding matters |

### Common Patterns

**Step-by-step process:**
```markdown
# Deployment Process

1. Build the image

<!-- pause -->

2. Push to registry

<!-- pause -->

3. Update the manifest

<!-- pause -->

4. ArgoCD syncs automatically
```

**Before / after code:**
```markdown
# Simplified Error Handling

<!-- column_layout: [1, 1] -->

<!-- column: 0 -->

**Before**

\```go
if err != nil {
    log.Println(err)
    return
}
\```

<!-- column: 1 -->

**After**

\```go
if err != nil {
    return fmt.Errorf("context: %w", err)
}
\```

<!-- reset_layout -->
```

**Closing slide:**
```markdown
# Thanks

Questions?

github.com/jedwards1230/deck
```

---

## Workflow: Creating a Presentation

1. **Get the brief** — topic, audience, approximate length, any key points to cover
2. **Draft an outline** — list slide titles first, confirm structure before writing content
3. **Write slides** — fill in content, add directives where appropriate
4. **Review** — check: one idea per slide? Consistent depth? Pauses feel natural?
5. **Write to file** — save as `<topic>.md` in the working directory
6. **Verify** — run `deck <file>.md` to confirm it renders correctly
