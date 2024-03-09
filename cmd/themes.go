package cmd

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

func AskThemes(config *Config) {
	options, _ := findAllThemes()
	selectedOption, err := pterm.DefaultInteractiveSelect.WithOptions(options).WithFilter(false).WithDefaultText("Select a theme").Show()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	config.Theme = selectedOption
}
