package cmd

func StartCMD(config *Config) {
	// Init Full Screen PTERM area (temporary disabled)
	// area, _ := pterm.DefaultArea.WithFullscreen().Start()
	// - Project Name -
	AskProjectName(config)
	// - Project Pages -
	AskPages(config)
	// - Project Theme -
	AskThemes(config)
	// - Project Location -
	AskProjectLocation(config)
	// - Project Confirmation -
	AskConfirmation(config)
	// - Project Creation -
	InitializeProject(config)
	// Terminate Full Screen PTERM area
	// area.Stop()
}
