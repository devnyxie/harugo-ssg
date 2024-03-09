package cmd

type Component struct {
	Index      int
	Name       string
	IsSelected bool
}

type Page struct {
	Index      int
	Name       string
	Components map[string]Component
}

type Config struct {
	ProjectName     string
	ProjectLocation string
	Theme           string
	Pages           map[string]Page
}
