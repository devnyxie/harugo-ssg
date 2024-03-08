package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	// Initial config
	config := Config{
		Theme: "",
		Pages: make(map[string]Page),
	}
	// Init Full Screen PTERM area
	area, _ := pterm.DefaultArea.WithFullscreen().Start()
	// - Project Name -
	// askProjectName(&config)
	// - Project Pages -
	askPages(&config)
	// - Project Theme -
	askThemes(&config)
	// - Project Confirmation -
	askConfirmation(&config)
	// - Overall config check -
	fmt.Println(config)
	// Terminate Full Screen PTERM area
	area.Stop()
}
