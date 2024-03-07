package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	var err error //generic error
	area, _ := pterm.DefaultArea.WithFullscreen().Start()
	config := Config{
		Theme: "",
		Pages: make(map[string]Page),
	}
	err = askPages(&config)
	if err != nil {
		panic(err)
	}
	err = askThemes(&config)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
	area.Stop()
}
