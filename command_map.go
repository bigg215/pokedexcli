package main

import (
	"fmt"
	"strings"
)

func commandMapf(cfg *config, args ...string) string {
	resp, err := cfg.papiClient.ListLocations(cfg.Next)
	out := ""
	if err != nil {
		return fmt.Sprintf("[red]error:[white] %s\n", err)
	}
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	for _, loc := range resp.Results {
		out += formatMapOutput(loc.URL, loc.Name)
	}
	return out
}

func commandMapb(cfg *config, args ...string) string {
	if cfg.Previous == nil {
		return "[red]error:[white] you're on the first page\n"
	}
	resp, err := cfg.papiClient.ListLocations(cfg.Previous)
	out := ""
	if err != nil {
		return fmt.Sprintf("[red]error:[white] %s\n", err)
	}
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	for _, loc := range resp.Results {
		out += formatMapOutput(loc.URL, loc.Name)
	}
	return out
}

func formatMapOutput(loc, name string) string {
	stripped := strings.TrimPrefix(loc, "https://pokeapi.co/api/v2/location-area/")
	trimmed := strings.Trim(stripped, "/")
	out := fmt.Sprintf("[yellow]%s[white]\t%s\n", trimmed, name)
	return out
}
