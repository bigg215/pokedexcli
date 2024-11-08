package main

import (
	"time"

	"github.com/bigg215/pokedexcli/internal/papi"
	"github.com/rivo/tview"
)

func main() {
	// Initialize application

	app := tview.NewApplication()
	papiClient := papi.NewClient(5*time.Second, 5*time.Minute)
	newUser := user{
		PokeDex: make(map[string]papi.Pokemon),
	}
	cfg := &config{
		appConfig:   app,
		papiClient:  papiClient,
		currentUser: newUser,
	}
	repl(cfg)

}

//func main() {
//cfg := &config{}
//startRepl(cfg)
//}
