package cmd

import (
	"github.com/pterm/pterm"
)

func askComponents(config *Config, selectedPageName string) error {
	var err error
	var components []Component

	// Define the default options indices
	components, err = findAllComponents()
	if err != nil {
		panic(err)
	}
	componentNames := []string{}
	selectedComponentNames := []string{}
	for i := range components {
		name := components[i].Name
		selected := IsSelectedFunc(config, selectedPageName, name)
		if selected {
			selectedComponentNames = append(selectedComponentNames, name)
		}
		componentNames = append(componentNames, name)
	}

	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions(componentNames).
		WithDefaultOptions(selectedComponentNames).
		WithFilter(false)
	selectedOptions, err := printer.Show()

	if err != nil {
		panic(err)
	}

	// Reset chosen components
	page := config.Pages[selectedPageName]
	page.Components = make(map[string]Component)
	config.Pages[selectedPageName] = page
	// Add selected components to the page
	for i := range selectedOptions {
		var selectedOption string = selectedOptions[i]
		addComponent(config, config.Pages[selectedPageName], selectedOption)
	}

	AskPages(config)

	return nil
}
