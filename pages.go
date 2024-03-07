package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pterm/pterm"
)

func askPages(config *Config) error {
	var options []string
	var err error
	for _, page := range config.Pages {
		options = append(options, page.Name)
	}
	options = append(options, "Add Page", "Continue")
	selectedOption, err := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	if selectedOption == "Add Page" {
		newPageName, err := pterm.DefaultInteractiveTextInput.Show()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}
		newPage := Page{
			ID:         uuid.New().String(),
			Name:       newPageName,
			Components: make(map[string]Component),
		}
		config.Pages[newPage.Name] = newPage
		askPages(config)
	} else if selectedOption == "Continue" {
		return nil
	} else {
		askComponents(config, selectedOption)
	}
	return nil
}
