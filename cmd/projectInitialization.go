package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

// --- main function to initialize the project ---
func InitializeProject(config *Config) {
	srcDir := "./base"
	destDir := config.ProjectName
	// - init base structure -
	copyPasteInitialStructure(config, srcDir, destDir)
	// - create pages, import & declare components -
	initPages(config, srcDir, destDir)
	// - create site_config.yaml -
	createSiteConfigYaml(config, destDir)
	// - setup theme -
	setupTheme(config, srcDir, destDir)
}

// --- utils ---
func copyPasteInitialStructure(config *Config, srcDir string, destDir string) {
	requiredPaths := getRequiredPaths(config)

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
		// for _, exception := range exceptions {
		// 	if strings.Contains(path, exception) {                                    to-do: enable before release
		// 		return nil
		// 	}
		// }

		// Skip files and folders that are not required
		found := false
		for _, requiredPath := range requiredPaths {
			if strings.HasPrefix(relPath, requiredPath) {
				found = true
				break
			}
		}
		if !found {
			return nil
		}

		if info.IsDir() {
			fmt.Println("Creating directory: " + destPath)
			err = os.MkdirAll(destPath, info.Mode())
			if err != nil {
				pterm.Println(pterm.Red(err))
				os.Exit(1)
			}
		} else {
			fmt.Println("Copying file from: " + path + " to: " + destPath)
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
	fmt.Printf("Project %s has been initialized successfully.\n", config.ProjectName)
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

func getRequiredPaths(config *Config) []string {
	var allowedPaths = []string{}
	// 1.
	for _, Page := range config.Pages {
		for _, Component := range Page.Components {
			ComponentPath := "components/" + Component.Name
			allowedPaths = append(allowedPaths, ComponentPath)
			// 2.
			if Component.Name == "blog" {
				allowedPaths = append(allowedPaths, "_posts")
			}
		}
	}
	// 3.
	allowedPaths = append(allowedPaths, "config")
	// 4.
	allowedPaths = append(allowedPaths, "public")
	// 5. theme setup is managed elsewhere
	// 6.
	allowedPaths = append(allowedPaths, "utils")
	// 7.
	allowedPaths = append(allowedPaths, "layouts")
	// 8.
	allowedPaths = append(allowedPaths, "pages", "pages/_app.js")
	// 9.
	allowedPaths = append(allowedPaths, "pages/defaultStyles")
	// 10.
	allowedPaths = append(allowedPaths, "package.json", "package-lock.json", ".gitignore", "next.config.js", "next.config.mjs", "jsconfig.json", ".")
	// 11.
	allowedPaths = append(allowedPaths, "defaultComponents")
	return allowedPaths
}

func initPages(config *Config, srcDir string, destDir string) {
	sortedPages := sortMapByIndex(config.Pages)
	for i, pageName := range sortedPages {
		page := config.Pages[pageName]
		pageSrcDir := srcDir + "/common/Page.js"
		var pageDestDir string = destDir + "/pages/"
		if i == 0 {
			pageDestDir = pageDestDir + "index.js"
		} else {
			pageDestDir = pageDestDir + strings.ToLower(page.Name) + ".js"
		}
		copyFile(pageSrcDir, pageDestDir)
		changeMainFunctionName(pageDestDir, "function Page", "function "+page.Name)
		for _, component := range page.Components {
			// capitalizedComponentName := strings.ToUpper(component.Name[:1]) + component.Name[1:]
			// 1.
			componentSrcDir := srcDir + "/components/" + component.Name
			// 2.
			err := filepath.Walk(componentSrcDir, func(path string, info os.FileInfo, err error) error {
				// exclude base path
				if path == componentSrcDir {
					return nil
				}
				//
				fmt.Println("walking over path: ", path)
				//
				compPageDestDir := pageDestDir
				filename := filepath.Base(path)
				fileNameWithoutExtension := removeFileExtension(filename)
				fmt.Println("fileNameWithoutExtension: ", fileNameWithoutExtension)
				JSComponentIdentifier := transformToJSIdentifier(fileNameWithoutExtension)
				CapitalizedJSComponentIdentifier := strings.ToUpper(JSComponentIdentifier[:1]) + JSComponentIdentifier[1:]

				if err != nil {
					pterm.Println(pterm.Red(err))
					os.Exit(1)
				}
				// =--

				if strings.Contains(fileNameWithoutExtension, "_api") {
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
					htmlToInsert := fmt.Sprintf("\ndata.%s = await %s();\n", funcName, funcName)
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
				} else if strings.Contains(fileNameWithoutExtension, "_slug") {
					fileNameWithoutSlug := strings.Replace(fileNameWithoutExtension, "_slug", "", 1)
					os.MkdirAll(destDir+"/pages/"+strings.ToLower(fileNameWithoutSlug), 0755)
					_ = copyFile(path, destDir+"/pages/"+strings.ToLower(fileNameWithoutSlug)+"/"+"["+fileNameWithoutSlug+"]"+extractFileExtension(filename))
				} else {
					// just components
					// --- 2.1, 2.2, 2.3
					content, err := os.ReadFile(compPageDestDir)
					if err != nil {
						fmt.Println(err)
						return err
					}
					modifiedContent := string(content)
					// is component (js, jsx) or a css module etc?
					isComponent := strings.Contains(filename, ".js") || strings.Contains(filename, ".jsx")
					// import all files - no harm here
					var importStatement string
					if isComponent {
						importStatement = fmt.Sprintf("\n import %s from '@/components/%s/%s'; \n", CapitalizedJSComponentIdentifier, component.Name, filename)
					} else if strings.Contains(filename, ".css") {
						importStatement = fmt.Sprintf("\n import '@/components/%s/%s'; \n", component.Name, filename)
					}
					startIndexImport := strings.Index(modifiedContent, "// IMPORTS START")
					endIndexImport := strings.Index(modifiedContent, "// IMPORTS END")
					if startIndexImport != -1 && endIndexImport != -1 {
						modifiedContent = modifiedContent[:endIndexImport] +
							importStatement +
							modifiedContent[endIndexImport:]
					} else {
						fmt.Println("Start or end comment not found")
					}

					// --- 2.4, 2.5, 2.6
					// declarations
					if isComponent {
						htmlToInsert := fmt.Sprintf("\n<%s {...props}/>\n", CapitalizedJSComponentIdentifier)
						startIndexContent := strings.Index(modifiedContent, "{/* CONTENT START */}")
						endIndexContent := strings.Index(modifiedContent, "{/* CONTENT END */}")
						if startIndexContent != -1 && endIndexContent != -1 {
							modifiedContent = modifiedContent[:startIndexContent+len("{/* CONTENT START */}")] +
								htmlToInsert +
								modifiedContent[endIndexContent:]
						} else {
							fmt.Println("Start or end comment not found")
						}
					}
					err = os.WriteFile(compPageDestDir, []byte(modifiedContent), 0644)
					if err != nil {
						return err
					}
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
	}
}

func changeMainFunctionName(path string, oldName string, newName string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	modifiedContent := string(content)
	modifiedContent = strings.Replace(modifiedContent, oldName, newName, -1)
	err = os.WriteFile(path, []byte(modifiedContent), 0644)
	if err != nil {
		return err
	}
	return nil
}

func createSiteConfigYaml(config *Config, destDir string) {
	// Marshal the config struct into YAML format
	ymlBytes, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println("Error marshaling YAML:", err)
		return
	}

	// Create the file
	filePath := destDir + "/config/site_config.yaml"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write YAML content to file
	_, err = file.Write(ymlBytes)
	if err != nil {
		fmt.Println("Error writing YAML to file:", err)
		return
	}

	fmt.Println("Config YAML file created successfully!")
}

func setupTheme(config *Config, srcDir string, destDir string) {
	// 1. copy theme files
	// 2. add import statements to _app.js (only for css files)
	themeDir := srcDir + "/themes/" + config.Theme
	destThemeDir := destDir + "/theme/" + config.Theme
	// copy all files from themeDir to destThemeDir
	err := filepath.Walk(themeDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			pterm.Println(pterm.Red(err))
			os.Exit(1)
		}
		relPath, err := filepath.Rel(themeDir, path)
		if err != nil {
			pterm.Println(pterm.Red(err))
			os.Exit(1)
		}
		destPath := filepath.Join(destThemeDir, relPath)

		if info.IsDir() {
			fmt.Println("Creating directory: " + destPath)
			err = os.MkdirAll(destPath, info.Mode())
			if err != nil {
				pterm.Println(pterm.Red(err))
				os.Exit(1)
			}
		} else {
			fmt.Println("Copying file from: " + path + " to: " + destPath)
			err = copyFile(path, destPath)
			if err != nil {
				pterm.Println(pterm.Red(err))
				os.Exit(1)
			}
			// add import statements to _app.js
			if strings.Contains(relPath, ".css") {
				fmt.Println("Adding css import statement to _app.js")
				err = insertContentBetweenComments(destDir+"/pages/_app.js", "// THEME IMPORT START", "// THEME IMPORT END", fmt.Sprintf("\n import '@/theme/%s/%s'; \n", config.Theme, relPath), false)
				if err != nil {
					pterm.Println(pterm.Red(err))
					os.Exit(1)
				}
			}
		}

		return nil
	},
	)
	if err != nil {
		pterm.Println(pterm.Red(err))
		os.Exit(1)
	}

}

func insertContentBetweenComments(filePath string, startComment string, endComment string, contentToInsert string, eraseExistingContent bool) error {
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

	if !eraseExistingContent {
		// Preserve existing content between the comments
		existingContent := content[startIndex:endIndex]
		contentToInsert = existingContent + contentToInsert
	}

	// Insert the content
	modifiedContent := content[:startIndex] + contentToInsert + content[endIndex:]

	// Write the modified content back to the file
	err = os.WriteFile(filePath, []byte(modifiedContent), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
