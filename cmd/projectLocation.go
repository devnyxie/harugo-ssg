package cmd

import (
	"os"

	"github.com/pterm/pterm"
)

func AskProjectLocation(config *Config) {
	projectDir, promptErr := pterm.DefaultInteractiveTextInput.WithDefaultText("Where to create your project?\n*Leave blank in order to create it in current directory").Show()
	if promptErr != nil {
		pterm.Error.Println(promptErr)
		os.Exit(1)
	}
	if projectDir == "" {
		projectDir = "."
	}
	config.ProjectLocation = projectDir
}
