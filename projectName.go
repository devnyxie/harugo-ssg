package main

import (
	"os"

	"github.com/pterm/pterm"
)

func askProjectName(config *Config) {
	newPageName, promptErr := pterm.DefaultInteractiveTextInput.WithDefaultText("Name your project").Show()
	if promptErr != nil {
		pterm.Println(pterm.Red("An unexpected error occured"))
		os.Exit(1)
	}
	if newPageName == "" {
		pterm.Println(pterm.Red("Project name cannot be empty"))
		askProjectName(config)
	} else {
		config.ProjectName = newPageName
	}
}
