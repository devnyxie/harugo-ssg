package cmd

import (
	"os"

	"github.com/pterm/pterm"
)

func AskProjectName(config *Config) {
	newPageName, promptErr := pterm.DefaultInteractiveTextInput.WithDefaultText("Name your project").Show()
	if promptErr != nil {
		pterm.Println(pterm.Red("An unexpected error occured"))
		os.Exit(1)
	}
	if newPageName == "" {
		pterm.Println(pterm.Red("Project name cannot be empty"))
		AskProjectName(config)
	} else {
		config.ProjectName = newPageName
	}
}
