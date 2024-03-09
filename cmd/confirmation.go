package cmd

import (
	"os"

	"github.com/pterm/pterm"
)

func AskConfirmation(config *Config) {
	// pterm.Println()
	pterm.DefaultTree.WithRoot(configToPtermTree(config)).Render()
	result, err := pterm.DefaultInteractiveConfirm.WithDefaultText("Confirm if this is the project you want").Show()
	if err != nil {
		pterm.Println(pterm.Red("An unexpected error occured"))
		os.Exit(1)
	}
	if result {
		pterm.Println(pterm.Green("Confirmed."))
		return
	} else {
		pterm.Println(pterm.Red("Not confirmed."))
		AskPages(config)
	}

}
