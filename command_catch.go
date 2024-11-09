package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) string {
	if len(args) != 1 {
		return "[red]error:[white] you must provide a pokemon name"
	}
	name := args[0]
	pokemon, err := cfg.papiClient.GetPokemon(name)
	if err != nil {
		return check(err)
	}
	species, err := cfg.papiClient.GetSpecies(pokemon.Species.URL)
	if err != nil {
		return check(err)
	}
	output := fmt.Sprintln()
	output += fmt.Sprintf("Throwing a [yellow]Pokeball[white] at [yellow]%s[white]...\n", pokemon.Name)

	if rand.IntN(255) > 255-species.CaptureRate {
		output += fmt.Sprintf("[yellow]%s[white] was [green]caught[white]!\n", pokemon.Name)
		cfg.currentUser.PokeDex[pokemon.Name] = pokemon
	} else {
		output += fmt.Sprintf("[yellow]%s[white] ran [red]away[white]!\n", pokemon.Name)
	}

	return output
}
