package cmd

type Component struct {
	Index int    `yaml:"index"`
	Name  string `yaml:"name"`
}

type Page struct {
	Index      int                  `yaml:"index"`
	Name       string               `yaml:"name"`
	Components map[string]Component `yaml:"components"`
}

type Entity interface{}

type Config struct {
	ProjectName     string          `yaml:"projectName"`
	ProjectLocation string          `yaml:"projectLocation"`
	Theme           string          `yaml:"theme"`
	Pages           map[string]Page `yaml:"pages"`
}
