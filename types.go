package main

type Component struct {
	ID         string
	Name       string
	IsSelected bool
}

type Page struct {
	ID         string
	Name       string
	Components map[string]Component
}

type Config struct {
	Theme string
	Pages map[string]Page
}
