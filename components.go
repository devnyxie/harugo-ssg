package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func askComponents(config *Config, selectedPageName string) error {
	var err error
	// --- all components ---
	var components []Component
	components, err = findAllComponents()
	components = append(components, Component{Name: "Continue"})
	if err != nil {
		panic(err)
	}

	for i := range components {
		name := components[i].Name
		selected := IsSelectedFunc(config, selectedPageName, name)
		if selected {
			fmt.Println("SETTING IsSelected TO TRUE")
			components[i].IsSelected = true
		}
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ .  }}?",
		Active:   `{{ .Name | green }}`,
		Inactive: `{{ if .IsSelected }}{{ .Name | yellow }}{{ else }}{{ .Name | faint }}{{ end }}`,
		Selected: "{{ .Name | faint }}",
		Details: `
--------- Component ----------
{{ "Name:" | faint }}    {{ .Name }}`,
	}

	prompt := promptui.Select{
		Label:        "Select components you would like to have on " + selectedPageName + " page",
		Items:        components,
		Templates:    templates,
		HideSelected: true,
	}
	i, _, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	selectedComponent := components[i]

	if selectedComponent.Name == "Continue" {
		askPages(config)
	} else {
		//to-do: add component to the page

		if IsSelectedFunc(config, selectedPageName, selectedComponent.Name) {
			fmt.Println("Component already exists in the page")

			deleteComponent(config, config.Pages[selectedPageName], selectedComponent.Name)
		} else {
			addComponent(config, config.Pages[selectedPageName], selectedComponent.Name)
		}
		askComponents(config, selectedPageName)
	}

	return nil
}
