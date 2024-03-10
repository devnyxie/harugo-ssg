package cmd

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

func AskPages(config *Config) {
	var options []string
	for _, page := range config.Pages {
		options = append(options, page.Name)
	}
	options = append(options, "Add Page", "Continue", "Exit")
	selectedOption, selectErr := pterm.DefaultInteractiveSelect.WithOptions(options).WithFilter(false).Show()
	if selectErr != nil {
		fmt.Printf("Prompt failed %v\n", selectErr)
		os.Exit(1)
	}

	if selectedOption == "Add Page" {
		addPage(config)
	} else if selectedOption == "Continue" {
		if len(config.Pages) == 0 {
			pterm.Warning.Println("No pages found. Please add a page.")
			AskPages(config)
		}
	} else if selectedOption == "Exit" {
		pterm.Println(pterm.Red("Exiting..."))
		os.Exit(0)
	} else {
		askComponents(config, selectedOption)
	}
}
