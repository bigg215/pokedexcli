package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	// Initialize application

	app := tview.NewApplication()
	cfg := &config{
		appConfig: app,
	}
	// Create label
	out := tview.NewTextView()

	// Create input field
	input := tview.NewInputField()

	// Create Grid containing the application's widgets
	appGrid := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(out, 0, 1, false).
		AddItem(input, 5, 1, true)

	// Capture user input
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Anything handled here will be executed on the main thread
		switch event.Key() {
		case tcell.KeyEnter:
			words := cleanInput(input.GetText())

			if len(words) == 0 {
				return nil
			}
			commandName := words[0]

			command, exists := getCommands()[commandName]
			if exists {
				output, err := command.callback(cfg)
				if err != nil {
					out.SetText(fmt.Sprintf("%s\n", err))
				}
				out.SetText(output)
			} else {
				out.SetText(fmt.Sprintln("Unknown command"))
			}
			//output := out.GetText(false) + input.GetText()
			//out.SetText(fmt.Sprintf("%s\n", output))
			input.SetText("")
			return nil
		case tcell.KeyEsc:
			// Exit the application
			app.Stop()
			return nil
		}
		return event
	})

	// Set the grid as the application root and focus the input field
	app.SetRoot(appGrid, true).SetFocus(input)

	// Run the application
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}

}

//func main() {
//cfg := &config{}
//startRepl(cfg)
//}
