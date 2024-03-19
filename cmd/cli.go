package cmd

func StartCMD(config *Config) {
	AskProjectName(config)
	// - Project Pages -
	AskPages(config)
	// - Project Theme -
	AskThemes(config)
	// - Project Confirmation -
	AskConfirmation(config)
	// - Project Creation -
	InitializeProject(config)
}
