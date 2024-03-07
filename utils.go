package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func IsSelectedFunc(config *Config, selectedPageName string, componentName string) bool {
	_, ok := config.Pages[selectedPageName].Components[componentName]
	return ok
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
			components = append(components, Component{Name: info.Name()})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return components, nil
}
