package main

import "fmt"

func commandPokedex(cfg *config, args ...string) string {
	output := fmt.Sprintln("Your Pokedex:")
	if len(cfg.currentUser.PokeDex) == 0 {
		output += fmt.Sprintln("[red]your pokedex is empty.[white]")
		return output
	}
	for _, val := range cfg.currentUser.PokeDex {
		output += fmt.Sprintf("\t- %s\n", val.Name)
	}

	return output
}
