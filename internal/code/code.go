package code

import "regexp"

var codeBlockRegex = regexp.MustCompile("(?s)```(\\w+)\\n(.*?)\\n```")

// Block represents an extracted code block.
type Block struct {
	Language string
	Code     string
}

// ExtractBlocks finds all code blocks in the given content.
func ExtractBlocks(content string) []Block {
	matches := codeBlockRegex.FindAllStringSubmatch(content, -1)
	blocks := make([]Block, 0, len(matches))
	for _, match := range matches {
		if len(match) >= 3 {
			blocks = append(blocks, Block{
				Language: match[1],
				Code:     match[2],
			})
		}
	}
	return blocks
}
