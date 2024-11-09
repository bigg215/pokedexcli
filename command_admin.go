package main

import "github.com/rivo/tview"

func commandAdmin(cfg *config, args ...string) string {
	list := tview.NewList().
		AddItem("List item 1", "Some explanatory text", 'a', nil).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("List item 3", "Some explanatory text", 'c', nil).
		AddItem("Main Screen", "Return to the root window", 'd', func() {
			cfg.appConfig.SetRoot(cfg.windowRoot, true)
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			cfg.appConfig.Stop()
		})
	cfg.appConfig.SetRoot(list, true)
	return ""
}
