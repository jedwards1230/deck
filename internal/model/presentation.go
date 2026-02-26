package model

// Presentation is the fully-parsed slide deck.
type Presentation struct {
	Slides      []Slide
	Frontmatter Frontmatter
}

// Frontmatter holds YAML metadata from the slide file header.
type Frontmatter struct {
	Theme  string `yaml:"theme"`
	Author string `yaml:"author"`
	Date   string `yaml:"date"`
	Paging string `yaml:"paging"`
	Footer string `yaml:"footer"`
}
