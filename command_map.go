package main

import (
	"errors"
	"fmt"

	"github.com/bigg215/pokedexcli/internal/papi"
)

func commandMapf(cfg *config) (string, error) {
	resp, err := papi.GetLocations(cfg.Next)
	if err != nil {
		return "", err
	}
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	return "", nil
}

func commandMapb(cfg *config) (string, error) {
	if cfg.Previous == nil {
		return "", errors.New("you're on the first page")
	}
	resp, err := papi.GetLocations(cfg.Previous)
	if err != nil {
		return "", err
	}
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	return "", nil
}
