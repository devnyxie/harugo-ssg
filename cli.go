package main

import (
	"fmt"           //Go standard lib
	"os"            //Go standard lib
	"path/filepath" //Go standard lib
	"regexp"        //Go standard lib

	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
)

type Component struct {
	ID       string
	Name     string
	Selected bool
}

type Page struct {
	ID         string
	Name       string
	Components map[string]Component //slice
}

type Config struct {
	Theme string
	Pages map[string]Page //map
}

// func to check if component is already selected
func componentExists(page Page, targetComponentName string) bool {
	for _, comp := range page.Components {
		if comp.Name == targetComponentName {
			return true
		}
	}
	return false
}

// func to delete checkbox from a string before adding it to the page
func removeCheckboxes(input string) string {
	pattern := `\s*\[(x| )\]\s*`
	r := regexp.MustCompile(pattern)
	result := r.ReplaceAllString(input, "")
	fmt.Println(result)
	return result
}

func addComponent(config *Config, page Page, targetComponentName string) {
	newComponent := Component{
		ID:   uuid.New().String(),
		Name: targetComponentName,
	}
	fmt.Println(newComponent.Name)
	config.Pages[page.Name].Components[newComponent.Name] = newComponent
}

func deleteComponent(config *Config, page Page, targetComponentName string) {
	delete(config.Pages[page.Name].Components, targetComponentName)
}

func findAllComponents() ([]Component, error) {
	var dir string = "./components"
	var components []Component
	/*
		filepath.Walk() call initiates the walk of the directory tree rooted at dir.
		For each file or directory encountered during the walk, it calls the provided function.
	*/
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != dir {
			components = append(components, Component{ID: info.Name(), Name: info.Name()})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return components, nil
}

func askComponents(config *Config, selectedPageName string) error {
	var err error
	// --- all components ---
	var components []Component
	components, err = findAllComponents()
	if err != nil {
		panic(err)
	}
	var componentsNames []string
	for _, comp := range components {
		if componentExists(config.Pages[selectedPageName], comp.Name) {
			componentsNames = append(componentsNames, comp.Name+" [x]")
			continue
		} else {
			componentsNames = append(componentsNames, comp.Name+" [ ]")
		}
	}
	prompt := promptui.Select{
		Label: "Select components you would like to have on " + selectedPageName + " page",
		Items: append(componentsNames, "continue"),
	}
	_, result, err := prompt.Run()
	if err != nil {
		panic(err)
	}
	if result == "continue" {
		askPages(config)
	} else {
		selectedCompNoCheckbox := removeCheckboxes(result)
		fmt.Printf("\nAdded %s to the list of components of the %s\n", selectedCompNoCheckbox, config.Pages[selectedPageName].Name)
		//to-do: add component to the page
		if componentExists(config.Pages[selectedPageName], selectedCompNoCheckbox) {
			fmt.Println("Component already exists in the page")

			deleteComponent(config, config.Pages[selectedPageName], selectedCompNoCheckbox)
		} else {
			addComponent(config, config.Pages[selectedPageName], selectedCompNoCheckbox)
		}
		askComponents(config, selectedPageName)
	}

	return nil
}

func askPages(config *Config) error {
	var pagesNames []string
	var err error
	for _, page := range config.Pages {
		pagesNames = append(pagesNames, page.Name)
	}
	pagesNames = append(pagesNames, "Add Page")
	pagesNames = append(pagesNames, "Continue")
	prompt := promptui.Select{
		Label: "Select pages you would like to modify, continue when ready",
		Items: pagesNames,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	if result == "Add Page" {
		//add page
		fmt.Println("Add Page")
		prompt := promptui.Prompt{
			Label: "Enter new page's name",
		}
		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}
		newPage := Page{
			ID:         uuid.New().String(),
			Name:       result,
			Components: make(map[string]Component),
		}
		config.Pages[newPage.Name] = newPage
		askPages(config)
	} else if result == "Continue" {
		return nil
	} else {
		askComponents(config, result)
	}
	return nil
}

func main() {
	config := Config{
		Theme: "",
		Pages: make(map[string]Page),
	}

	var err error //generic error
	// step 1: Ask for Pages
	err = askPages(&config)
	if err != nil {
		panic(err)
	}

	// for _, comp := range components {
	// 	fmt.Println(comp.Name)
	// }
	fmt.Println(config)
}
