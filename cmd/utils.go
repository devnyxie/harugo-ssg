package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pterm/pterm"
)

func IsSelectedFunc(config *Config, selectedPageName string, componentName string) bool {
	_, ok := config.Pages[selectedPageName].Components[componentName]
	return ok
}

func addPage(config *Config) {
	var newPageIndex int
	if len(config.Pages) == 0 {
		newPageIndex = 0
	} else {
		newPageIndex = len(config.Pages)
	}
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
	AskPages(config)
}

func addComponent(config *Config, page Page, targetComponentName string) {
	var newComponentIndex int
	if len(config.Pages[page.Name].Components) == 0 {
		newComponentIndex = 0
	} else {
		newComponentIndex = len(config.Pages[page.Name].Components)
	}
	newComponent := Component{
		Index: newComponentIndex,
		Name:  targetComponentName,
	}
	fmt.Println(newComponent.Name)
	config.Pages[page.Name].Components[newComponent.Name] = newComponent
}

func findAllComponents() ([]Component, error) {
	var dir string = "./foundation/components"
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
	var dir string = "./foundation/themes"
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
			{Text: fmt.Sprintf("Location: %s", config.ProjectLocation)},
			{Text: fmt.Sprintf("Theme: %s", config.Theme)},
		},
	}

	for pageName, page := range config.Pages {
		pageNode := pterm.TreeNode{
			Text: fmt.Sprintf("[%d] Page: %s", page.Index, pageName),
		}

		for componentName, component := range page.Components {
			componentNode := pterm.TreeNode{
				Text: fmt.Sprintf("[%d] Component: %s", component.Index, componentName),
			}
			pageNode.Children = append(pageNode.Children, componentNode)
		}

		tree.Children = append(tree.Children, pageNode)
	}

	return tree
}

func stringExistsInSlice(target string, slice []string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}
func sortMapByIndex(entity Entity) []string {
	switch v := entity.(type) {
	case map[string]Component:

		keys := make([]string, 0, len(v))
		for key := range v {
			keys = append(keys, key)
		}
		sort.SliceStable(keys, func(i, j int) bool {
			return v[keys[i]].Index < v[keys[j]].Index
		})

		return keys
	case map[string]Page:
		keys := make([]string, 0, len(v))
		for key := range v {
			keys = append(keys, key)
		}
		sort.SliceStable(keys, func(i, j int) bool {
			return v[keys[i]].Index < v[keys[j]].Index
		})
		return keys

	default:
		fmt.Println("Unsupported type")
		return nil
	}
}

func removeFileExtension(filename string) string {
	extension := filepath.Ext(filename)
	return strings.TrimSuffix(filename, extension)
}

func extractFileExtension(filename string) string {
	return filepath.Ext(filename)
}
