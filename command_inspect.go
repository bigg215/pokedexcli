package main

import "fmt"

func commandInspect(cfg *config, args ...string) string {
	if len(args) != 1 {
		return "[red]error:[white] you must provide a pokemon name\n"
	}
	name := args[0]
	if val, ok := cfg.currentUser.pokeDex[name]; ok {
		output := formatStats("Name", val.Name)
		output += formatStats("Height", fmt.Sprintf("%d", val.Height))
		output += formatStats("Weight", fmt.Sprintf("%d", val.Weight))
		output += formatStats("Stats", "")

		for _, stat := range val.Stats {
			output += formatStats("\t-"+stat.Stat.Name, fmt.Sprintf("%d", stat.BaseStat))
		}

		output += formatStats("Types", "")

		for _, types := range val.Types {
			output += fmt.Sprintf("\t- %s\n", types.Type.Name)
		}

		return output
	}
	return "[red]error:[white] unknown pokemon\n"
}

func formatStats(pre, post string) string {
	return fmt.Sprintf("[yellow]%s:[white] %s\n", pre, post)
}
