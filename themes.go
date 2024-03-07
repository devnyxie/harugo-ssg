package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func askThemes(config *Config) error {
	options, _ := findAllThemes()
	selectedOption, err := pterm.DefaultInteractiveSelect.WithOptions(options).WithFilter(false).WithDefaultText("Select a theme").Show()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}
	config.Theme = selectedOption
	return nil
}
