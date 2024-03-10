package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
)

func InitializeProject(config *Config) {
	srcDir := "./foundation"
	var destDir string
	if config.ProjectLocation != "." {
		destDir = filepath.Join(config.ProjectLocation, config.ProjectName)
	} else {
		destDir = config.ProjectName
	}
	// --- Initialization ---
	// 1. Copy nextjs-base to desired location & rename it to config.ProjectName.
	// --- Pages & their components ---
	// 2. Create files for each page in project/pages
	//    2.1. Get "common/page.jsx", rename it and use it.
	//    2.2. For each Page.Component:
	//         2.2.1. Find the component's folder in "components/componentName"
	//         2.2.2. We will have files of different type in these folders. Each one ends differently: _api, _util etc.
	//                This suffix points to where this components belongs. Copy them to their folders inside of the project.
	//         2.2.3: These components share same dir-structure, so they all are already connected theoretically,
	//                But we still have to connect them to their pages! (only _component's)
	//				  Each page has comments specifically for this occasion:
	//                // --- imports start --- ,  --- imports end ---
	//                // --- content start --- , --- content end ---
	//				  Find these comments, import & declare Page's components.
	// --- Theming ---
	// 3. Grab required theme from /themes and put it in project/styles/theme.js (emotionCSS theming, optional dark mode)
	//
	// *Notes:
	// Parse jsx files using goquery.

	// init
	copyPasteInitialStructure(config, srcDir, destDir)
	// LOOP
	for _, page := range config.Pages {
		//1. page srcDir = "./foundation/common/Pa	e.jsx"
		pageSrcDir := "./foundation/common/Page.jsx"

		pageDestDir := "./" + config.ProjectName + "/pages/" + page.Name + ".jsx"
		//config pageName
		// pageName := page.Name;
		//
		copyFile(pageSrcDir, pageDestDir)
	}

	//create func to copy paste ./common/Page.jsx into folder pages and rename the file & inner func name

}

func copyPasteInitialStructure(config *Config, srcDir string, destDir string) {
	exceptions := getExceptions(config)

	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			pterm.Println(pterm.Red(err))
			os.Exit(1)
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			pterm.Println(pterm.Red(err))
			os.Exit(1)
		}

		destPath := filepath.Join(destDir, relPath)

		// Skip exceptions (files and folders)
		for _, exception := range exceptions {
			if strings.Contains(path, exception) {
				return nil
			}
		}

		if info.IsDir() {
			err = os.MkdirAll(destPath, info.Mode())
			if err != nil {
				pterm.Println(pterm.Red(err))
				os.Exit(1)
			}
		} else {
			err = copyFile(path, destPath)
			if err != nil {
				pterm.Println(pterm.Red(err))
				os.Exit(1)
			}
		}

		return nil
	})
	if err != nil {
		pterm.Println(pterm.Red(err))
		os.Exit(1)
	}
	fmt.Println("Base was just created in: " + destDir)
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func getExceptions(config *Config) []string {
	var exceptions []string = []string{"node_modules", "2023-05-05.md", ".next"}
	allComponents, _ := findAllComponents()
	var allChosenComponents []string
	for _, page := range config.Pages {
		for _, component := range config.Pages[page.Name].Components {
			allChosenComponents = append(allChosenComponents, component.Name)
		}
	}

	for _, comp := range allComponents {
		if !stringExistsInSlice(comp.Name, allChosenComponents) {
			exceptions = append(exceptions, comp.Name)
		}
	}

	return exceptions
}
