package main

import "fmt"

func commandPokedex(cfg *config, args ...string) string {
	output := fmt.Sprintln("Your Pokedex:")
	if len(cfg.currentUser.pokeDex) == 0 {
		output += fmt.Sprintln("[red]your pokedex is empty.[white]")
		return output
	}
	for _, val := range cfg.currentUser.pokeDex {
		output += fmt.Sprintf("\t- %s\n", val.Name)
	}

	return output
}
