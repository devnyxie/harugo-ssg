package main

import (
	"harugo/cmd"
)

func main() {
	config := cmd.Config{
		Theme: "",
		Pages: make(map[string]cmd.Page),
	}
	cmd.StartCMD(&config)
}
