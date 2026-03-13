# Claude Plugin Design

**Date**: 2026-03-13
**Branch**: feat/claude-plugin

## Goal

Ship a self-contained Claude Code plugin from the `deck` repo so users can install it via `/plugin marketplace add jedwards1230/deck` and get AI-assisted presentation creation.

## Components

### Marketplace Registry

`.claude-plugin/marketplace.json` — registers `deck` as the single plugin in this repo.

### Plugin: `deck`

`plugins/deck/.claude-plugin/plugin.json` — name, version, description, author.

### Skill: `deck`

`plugins/deck/skills/deck/SKILL.md`

The authoritative reference for creating deck presentations. Covers:
- Slide format (`---` separators, YAML frontmatter fields)
- All comment directives: `<!-- pause -->`, `<!-- column_layout: [w,w] -->`, `<!-- column: N -->`, `<!-- reset_layout -->`, `<!-- speaker_note: ... -->`
- Footer template variables
- Navigation keys (for speaker reference)
- Code execution blocks
- Opinionated guidance: slide count, content density, when to use columns vs progressive reveal
- Workflow: draft outline → write slides → review structure → write to file

Triggers: "create a presentation", "make slides", "write a deck", "deck presentation", "make me a slide deck".

### Agent: `deck-creator`

`plugins/deck/agents/deck-creator.md`

Lightweight agent for net-new presentations. Takes a topic + context from the user, drafts the full `.md` file, writes it to disk. Imports the `deck` skill for full syntax reference. Isolated context keeps presentation work out of the user's main conversation.

### Agent: `deck-editor`

`plugins/deck/agents/deck-editor.md`

Lightweight agent for iterating on an existing deck file. Reads the file, improves structure, adds progressive reveals where useful, tightens content density. Imports the `deck` skill. Triggered when a `.md` file already exists.

## GHA: Version Check

`.github/workflows/plugin-version-check.yml`

Fires on push touching `plugins/**` or `.claude-plugin/marketplace.json`. Validates that each plugin's `plugin.json` version matches its entry in `marketplace.json`. Fails if they diverge, blocking merge.

## File Tree

```
deck/
├── .claude-plugin/
│   └── marketplace.json
├── plugins/
│   └── deck/
│       ├── .claude-plugin/
│       │   └── plugin.json
│       ├── skills/
│       │   └── deck/
│       │       └── SKILL.md
│       └── agents/
│           ├── deck-creator.md
│           └── deck-editor.md
└── .github/
    └── workflows/
        └── plugin-version-check.yml
```

## Installation

```bash
/plugin marketplace add jedwards1230/deck
/plugin install deck@jedwards1230-deck
```
