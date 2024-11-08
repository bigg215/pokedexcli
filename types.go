package main

import (
	"github.com/bigg215/pokedexcli/internal/papi"
	"github.com/rivo/tview"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) string
}

type config struct {
	Next        *string
	Previous    *string
	appConfig   *tview.Application
	papiClient  papi.Client
	currentUser user
}

type user struct {
	PokeDex map[string]papi.Pokemon `json:"pokedex"`
}
