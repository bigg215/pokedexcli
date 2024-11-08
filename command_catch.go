package main

import "fmt"

func commandCatch(cfg *config, args ...string) string {
	if len(args) != 1 {
		return "[red]error:[white] you must provide a pokemon name"
	}
	name := args[0]
	pokemon, err := cfg.papiClient.GetPokemon(name)
	if err != nil {
		return fmt.Sprintf("[red]error:[white] %s\n", err)
	}
	output := fmt.Sprintln()
	output += fmt.Sprintf("Throwing a [yellow]Pokeball[white] at [green]%s[white]...\n", pokemon.Name)
	output += fmt.Sprintf("[green]%s[white] was caught!\n", pokemon.Name)
	cfg.currentUser.PokeDex[pokemon.Name] = pokemon
	return output
}
