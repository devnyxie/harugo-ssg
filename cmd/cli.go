package cmd

import (
	"github.com/pterm/pterm"
)

func StartCMD(config *Config) {
	// Initial config

	// Init Full Screen PTERM area
	area, _ := pterm.DefaultArea.WithFullscreen().Start()
	// - Project Name -
	AskProjectName(config)
	// - Project Pages -
	AskPages(config)
	// - Project Theme -
	AskThemes(config)
	// - Project Location -
	AskProjectLocation(config)
	// - Project Confirmation -
	AskConfirmation(config)
	// - Project Creation -
	InitializeProject(config)
	// Terminate Full Screen PTERM area
	area.Stop()
}
