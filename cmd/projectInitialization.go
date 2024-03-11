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

	// init
	copyPasteInitialStructure(config, srcDir, destDir)
	//sort pages in order to start from index [0]
	sortedPages := sortMapByIndex(config.Pages)
	for i, pageName := range sortedPages {
		page := config.Pages[pageName]
		pageSrcDir := "./foundation/common/Page.jsx"
		var pageDestDir string = destDir + "/pages/"
		var pagePath string
		if i == 0 {
			pageDestDir = pageDestDir + "index.jsx"
			pagePath = "/"
		} else {
			pageDestDir = pageDestDir + strings.ToLower(page.Name) + ".jsx"
			pagePath = fmt.Sprintf("/%s", strings.ToLower(page.Name))
		}
		copyFile(pageSrcDir, pageDestDir)
		for _, component := range page.Components {
			capitalizedComponentName := strings.ToUpper(component.Name[:1]) + component.Name[1:]
			// 1.
			componentSrcDir := "./foundation/components/" + component.Name
			// 2.
			err := filepath.Walk(componentSrcDir, func(path string, info os.FileInfo, err error) error {
				filename := filepath.Base(path)
				if err != nil {
					pterm.Println(pterm.Red(err))
					os.Exit(1)
				}
				if strings.Contains(path, "_component") {
					// --- 2.1, 2.2, 2.3
					content, err := os.ReadFile(pageDestDir)
					if err != nil {
						fmt.Println(err)
						return err
					}
					modifiedContent := string(content)
					htmlToInsertImport := fmt.Sprintf("\n import %s from '@/components/%s/%s'; \n", capitalizedComponentName, component.Name, filename)
					startIndexImport := strings.Index(modifiedContent, "// IMPORTS START")
					endIndexImport := strings.Index(modifiedContent, "// IMPORTS END")
					if startIndexImport != -1 && endIndexImport != -1 {
						modifiedContent = modifiedContent[:endIndexImport] +
							htmlToInsertImport +
							modifiedContent[endIndexImport:]
					} else {
						fmt.Println("Start or end comment not found")
					}
					// --- 2.4, 2.5, 2.6
					htmlToInsert := fmt.Sprintf("\n<%s {...props}/>\n", capitalizedComponentName)
					startIndex := strings.Index(modifiedContent, "{/* CONTENT START */}")
					endIndex := strings.Index(modifiedContent, "{/* CONTENT END */}")
					if startIndex != -1 && endIndex != -1 {
						modifiedContent = modifiedContent[:endIndex] +
							htmlToInsert +
							modifiedContent[endIndex:]
					} else {
						fmt.Println("Start or end comment not found")
					}

					err = os.WriteFile(pageDestDir, []byte(modifiedContent), 0644)
					if err != nil {
						return err
					}

				}
				if strings.Contains(path, "_api") {
					//we need to add both import and getStaticProps !!!
					// part 1:
					// add import to the page (  )

					// part 2:
					//  Find the comment // STATIC PROPS START
					//  Find the comment // STATIC PROPS END
					//  Add getStaticProps to the page
					//  Use the file's default exported func in getStaticProps and pass props to the page, and then to the component
					//
					// err = copyFile(path, destPath)
					content, err := os.ReadFile(pageDestDir)
					modifiedContent := string(content)
					funcName := removeFileExtension(filename)
					if err != nil {
						fmt.Println(err)
						return err
					}
					// part 1
					htmlToInsertImport := fmt.Sprintf("\n import %s from '@/components/%s/%s'; \n", funcName, component.Name, filename)
					startIndexImport := strings.Index(modifiedContent, "// IMPORTS START")
					endIndexImport := strings.Index(modifiedContent, "// IMPORTS END")
					if startIndexImport != -1 && endIndexImport != -1 {
						modifiedContent = modifiedContent[:endIndexImport] +
							htmlToInsertImport +
							modifiedContent[endIndexImport:]
					} else {
						fmt.Println("Start or end comment not found")
					}

					// getStaticProps done
					htmlToInsert := fmt.Sprintf("\n export async function getStaticProps() { \n const data = await %s(); \n return { props: { %s: data } }; \n } \n", funcName, funcName)
					startIndex := strings.Index(modifiedContent, "// STATIC PROPS START")
					endIndex := strings.Index(modifiedContent, "// STATIC PROPS END")
					if startIndex != -1 && endIndex != -1 {
						modifiedContent = modifiedContent[:startIndex+len("// STATIC PROPS START")] +
							htmlToInsert +
							modifiedContent[endIndex:]
					} else {
						fmt.Println("Start or end comment not found")
					}
					//find already existing components between // CONTENT START and // CONTENT END and add props of props.funcName to them
					// --- 2.4, 2.5, 2.6

					// save
					err = os.WriteFile(pageDestDir, []byte(modifiedContent), 0644)
					if err != nil {
						return err
					}
				}
				if strings.Contains(path, "_slug") {
					fileNameWithoutExtension := removeFileExtension(filename)
					fileNameWithoutSlug := strings.Replace(fileNameWithoutExtension, "_slug", "", 1)
					os.MkdirAll(destDir+"/pages/"+strings.ToLower(fileNameWithoutSlug), 0755)
					_ = copyFile(path, destDir+"/pages/"+strings.ToLower(fileNameWithoutSlug)+"/"+"["+fileNameWithoutSlug+"]"+extractFileExtension(filename))

				}
				return nil
			})
			if err != nil {
				pterm.Println(pterm.Red(err))
				os.Exit(1)
			}

			if err != nil {
				pterm.Println(pterm.Red(err))
				os.Exit(1)
			}
		}
		// add pages to the config.yaml
		ymlToInsert := fmt.Sprintf("\n- pageName: %s\n  pagePath: %s\n", page.Name, pagePath)
		if i == 0 {
			ymlToInsert = "\npages:" + ymlToInsert
		}
		// insertContentBetweenComments implementation
		err := insertContentBetweenComments(destDir+"/config/config.yaml", "# --- Pages start ---", "# --- Pages end ---", ymlToInsert)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	// --- add website's title and description to the config.yaml ---
	ymlToInsert := fmt.Sprintf("\n- site_title: %s\n  site_description: %s\n", config.ProjectName, config.ProjectName+" website's description.")
	err := insertContentBetweenComments(destDir+"/config/config.yaml", "# --- Website settings start ---", "# --- Website settings end ---", ymlToInsert)
	if err != nil {
		fmt.Println(err)
		return
	}
	// TO-DO
	// --- Theming ---
	// 3. Grab required theme from /themes and put it in project/styles/theme.js (emotionCSS theming, optional dark mode)
	//
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
	var exceptions []string = []string{"node_modules", "2023-05-05.md", ".next", "common", "themes"}
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

func insertContentBetweenComments(filePath string, startComment string, endComment string, contentToInsert string) error {
	// Read the file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Convert to string
	content := string(fileContent)

	// Find the start and end indices
	startIndex := strings.Index(content, startComment)
	endIndex := strings.Index(content, endComment)

	// Check if the start and end comments were found
	if startIndex == -1 || endIndex == -1 {
		return fmt.Errorf("start or end comment not found")
	}

	// Adjust the start index to point to the end of the start comment
	startIndex += len(startComment)

	// Insert the content
	modifiedContent := content[:startIndex] + contentToInsert + content[endIndex:]

	// Write the modified content back to the file
	err = os.WriteFile(filePath, []byte(modifiedContent), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
