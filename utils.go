package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
)

func IsSelectedFunc(config *Config, selectedPageName string, componentName string) bool {
	_, ok := config.Pages[selectedPageName].Components[componentName]
	return ok
}

func addPage(config *Config) {
	newPageIndex := len(config.Pages) + 1
	newPageName, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Name of the page").Show()
	if newPageName == "" {
		pterm.Println(pterm.Red("Page name cannot be empty"))
		addPage(config)
	}
	newPage := Page{
		Index:      newPageIndex,
		Name:       newPageName,
		Components: make(map[string]Component),
	}
	config.Pages[newPage.Name] = newPage
	askPages(config)
}

func addComponent(config *Config, page Page, targetComponentName string) {
	newComponentIndex := len(config.Pages[page.Name].Components) + 1
	newComponent := Component{
		Index: newComponentIndex,
		Name:  targetComponentName,
	}
	fmt.Println(newComponent.Name)
	config.Pages[page.Name].Components[newComponent.Name] = newComponent
}

func findAllComponents() ([]Component, error) {
	var dir string = "./components"
	var components []Component
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != dir {
			components = append(components, Component{Name: info.Name()})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return components, nil
}

func findAllThemes() ([]string, error) {
	var dir string = "./themes"
	var themes []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != dir {
			themes = append(themes, info.Name())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return themes, nil
}

func configToPtermTree(config *Config) pterm.TreeNode {
	tree := pterm.TreeNode{
		Text: config.ProjectName,
		Children: []pterm.TreeNode{
			{Text: fmt.Sprintf("Theme: %s", config.Theme)},
		},
	}

	for pageName, page := range config.Pages {
		pageNode := pterm.TreeNode{
			Text: fmt.Sprintf("Page: %s", pageName),
		}

		for componentName, component := range page.Components {
			componentNode := pterm.TreeNode{
				Text: fmt.Sprintf("Component: %s", componentName),
				Children: []pterm.TreeNode{
					{Text: fmt.Sprintf("Name: %s", component.Name)},
				},
			}
			pageNode.Children = append(pageNode.Children, componentNode)
		}

		tree.Children = append(tree.Children, pageNode)
	}

	return tree
}
