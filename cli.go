package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bigg215/pokedexcli/internal/papi"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) string
}

type config struct {
	Next       *string
	Previous   *string
	appConfig  *tview.Application
	papiClient papi.Client
}

func repl(cfg *config) {
	out := tview.NewTextView()
	out.SetBorder(true)
	out.SetDynamicColors(true)

	// Create input field
	input := tview.NewInputField()
	input.SetBorder(true)
	input.SetLabel("pokedex> ")

	// Create Grid containing the application's widgets
	appGrid := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(out, 0, 5, false).
		AddItem(input, 3, 1, true)

	// Capture user input
	cfg.appConfig.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Anything handled here will be executed on the main thread
		switch event.Key() {
		case tcell.KeyEnter:
			words := cleanInput(input.GetText())

			if len(words) == 0 {
				return nil
			}
			commandName := words[0]
			args := []string{}

			if len(words) > 1 {
				args = words[1:]
			}

			curr := out.GetText(false)
			command, exists := getCommands()[commandName]
			if exists {
				output := command.callback(cfg, args...)
				out.SetText(curr + output)
			} else {
				out.SetText(curr + fmt.Sprintln("[red]error:[white] unknown command[white]"))
			}
			//output := out.GetText(false) + input.GetText()
			//out.SetText(fmt.Sprintf("%s\n", output))
			input.SetText("")
			out.ScrollToEnd()
			return nil
		case tcell.KeyEsc:
			// Exit the application
			cfg.appConfig.Stop()
			return nil
		}
		return event
	})

	// Set the grid as the application root and focus the input field
	cfg.appConfig.SetRoot(appGrid, true).SetFocus(input)

	// Run the application
	err := cfg.appConfig.Run()
	if err != nil {
		log.Fatal(err)
	}

}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List next 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location>",
			description: "Explore a location",
			callback:    commandExplore,
		},
	}
}
