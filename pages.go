package main

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

func askPages(config *Config) error {
	var options []string
	var err error
	for _, page := range config.Pages {
		options = append(options, page.Name)
	}
	options = append(options, "Add Page", "Continue", "Exit")
	selectedOption, err := pterm.DefaultInteractiveSelect.WithOptions(options).WithFilter(false).Show()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	if selectedOption == "Add Page" {
		addPage(config)
	} else if selectedOption == "Continue" {
		if len(config.Pages) == 0 {
			pterm.Warning.Println("No pages found. Please add a page.")
			askPages(config)
		}
		return nil
	} else if selectedOption == "Exit" {
		pterm.Println(pterm.Red("Exiting..."))
		os.Exit(0)
	} else {
		askComponents(config, selectedOption)
	}
	return nil
}
