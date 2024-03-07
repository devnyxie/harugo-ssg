package main

import (
	"fmt"
)

func main() {
	var err error //generic error
	config := Config{
		Theme: "",
		Pages: make(map[string]Page),
	}
	err = askPages(&config)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}
