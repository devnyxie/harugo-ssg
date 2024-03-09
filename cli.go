package main

import (
	"fmt"

	"harugo/cmd"

	"github.com/pterm/pterm"
)

func main() {
	// Initial config
	config := cmd.Config{
		Theme: "",
		Pages: make(map[string]cmd.Page),
	}
	// Init Full Screen PTERM area
	area, _ := pterm.DefaultArea.WithFullscreen().Start()
	// - Project Name -
	cmd.AskProjectName(&config)
	// - Project Pages -
	cmd.AskPages(&config)
	// - Project Theme -
	cmd.AskThemes(&config)
	// - Project Location -
	cmd.AskProjectLocation(&config)
	// - Project Confirmation -
	cmd.AskConfirmation(&config)
	// - Overall config check -
	fmt.Println(config)
	// Terminate Full Screen PTERM area
	area.Stop()
}
