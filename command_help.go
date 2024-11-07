package main

import "fmt"

func commandHelp(cfg *config, args ...string) string {

	logo := "   /$$$$$$$           /$$                       /$$\n"
	logo += "  | $$__  $$         | $$                      | $$\n"
	logo += "  | $$  | $$ /$$$$$$ | $$   /$$  /$$$$$$   /$$$$$$$  /$$$$$$  /$$   /$$\n"
	logo += "  | $$$$$$$//$$__  $$| $$  /$$/ /$$__  $$ /$$__  $$ /$$__  $$|  $$ /$$/\n"
	logo += "  | $$____/| $$  | $$| $$$$$$/ | $$$$$$$$| $$  | $$| $$$$$$$$ |  $$$$/\n"
	logo += "  | $$     | $$  | $$| $$_  $$ | $$_____/| $$  | $$| $$_____/  >$$  $$\n"
	logo += "  | $$     |  $$$$$$/| $$ |  $$|  $$$$$$$|  $$$$$$$|  $$$$$$$ /$$/|  $$\n"
	logo += "  |__/      |______/ |__/  |__/ |_______/ |_______/ |_______/|__/  |__/\n"

	out := fmt.Sprintln(logo)
	out += fmt.Sprintln("[yellow]Welcome to the pokedex![white]")
	out += fmt.Sprintln("Usage:")
	out += fmt.Sprintln()

	for _, cmd := range getCommands() {
		out += fmt.Sprintf("[yellow]%s:[white] %s\n", cmd.name, cmd.description)
	}

	out += fmt.Sprintln()
	return out
}
