package main

import "fmt"

func commandHelp(cfg *config) (string, error) {
	out := fmt.Sprintln()
	out += fmt.Sprintln("Welcome to the pokedex!")
	out += fmt.Sprintln("Usage:")
	out += fmt.Sprintln()

	for _, cmd := range getCommands() {
		out += fmt.Sprintf("%s: %s\n", cmd.name, cmd.description)
	}

	out += fmt.Sprintln()
	return out, nil
}
