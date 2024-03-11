package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/pterm/pterm"
)

func AskPages(config *Config) {
	var options []string
	sortedPages := sortMapByIndex(config.Pages)
	if sortedPages != nil {
		options = append(options, sortedPages...)
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
			return
		} else {
			return
		}
	} else if selectedOption == "Exit" {
		pterm.Println(pterm.Red("Exiting..."))
		os.Exit(0)
	} else {
		askComponents(config, selectedOption)
	}
}

func addPage(config *Config) {
	var newPageIndex int
	if len(config.Pages) == 0 {
		newPageIndex = 0
	} else {
		newPageIndex = len(config.Pages)
	}
	newPageName, err := pterm.DefaultInteractiveTextInput.WithDefaultText("Name of the page").Show()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	newPageName = strings.TrimSpace(newPageName)
	if newPageName == "" {
		pterm.Println(pterm.Red("Page name cannot be empty"))
		return
	}
	newPage := Page{
		Index:      newPageIndex,
		Name:       newPageName,
		Components: make(map[string]Component),
	}
	config.Pages[newPage.Name] = newPage
	AskPages(config)
}
