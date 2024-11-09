package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func repl(cfg *config) {
	out := tview.NewTextView()
	out.SetBorder(true)
	out.SetDynamicColors(true)

	// Create input field
	input := tview.NewInputField()
	input.SetBorder(true)
	input.SetLabel("pokedex> ")
	input.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			words := cleanInput(input.GetText())
			if len(words) == 0 {
				return
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
			input.SetText("")
			out.ScrollToEnd()
			return
		}
	})

	// Create Grid containing the application's widgets
	appGrid := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(out, 0, 5, false).
		AddItem(input, 3, 1, true)

	// Capture user input
	cfg.appConfig.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Anything handled here will be executed on the main thread
		switch event.Key() {
		case tcell.KeyEsc:
			cfg.appConfig.SetRoot(appGrid, true).SetFocus(input)
			return nil
		}
		return event
	})

	// Set the grid as the application root and focus the input field
	cfg.appConfig.SetRoot(appGrid, true).SetFocus(input)
	cfg.windowRoot = appGrid
	cfg.windowCMD = out
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
		"catch": {
			name:        "catch <pokemon>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspect a known pokemon's stats",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display current pokeDex entires",
			callback:    commandPokedex,
		},
		"save": {
			name:        "save",
			description: "Save the current game",
			callback:    commandSave,
		},
		"load": {
			name:        "load",
			description: "Load a saved game",
			callback:    commandLoad,
		},
		"admin": {
			name:        "admin",
			description: "testing",
			callback:    commandAdmin,
		},
	}
}
