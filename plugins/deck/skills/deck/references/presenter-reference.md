# Presenter Reference

Runtime behavior and navigation keys for the `deck` presenter. This is for human reference, not slide authoring.

## Navigation Keys

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

Press `ctrl+e` to execute the last code block on the current slide. Supported languages: Go, Bash, Python, JavaScript, Ruby. Test code blocks before presenting.

## Hot Reload

When presenting a file, deck watches for changes and automatically jumps to the modified slide. Edit the source file while presenting to iterate live.

## Running deck

```bash
# Present a file
deck slides.md

# Pipe content
cat slides.md | deck

# Built-in tutorial
deck
```
