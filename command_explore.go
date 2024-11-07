package main

import "fmt"

func commandExplore(cfg *config, args ...string) string {
	if len(args) != 1 {
		return "[red]error:[white] you must provide a location name"
	}
	name := args[0]
	location, err := cfg.papiClient.GetLocation(name)
	if err != nil {
		return fmt.Sprintf("[red]error:[white] %s\n", err)
	}
	output := fmt.Sprintln()
	output += fmt.Sprintf("Exploring [green]%s[white]...\n", location.Name)
	output += fmt.Sprintln("[yellow]Found Pokemon: [white]")
	for _, enc := range location.PokemonEncounters {
		output += fmt.Sprintf(" - %s\n", enc.Pokemon.Name)
	}
	output += fmt.Sprintln()
	return output
}
