package main

import (
	"time"

	"github.com/bigg215/pokedexcli/internal/papi"
	"github.com/google/uuid"
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
	windowRoot  *tview.Flex
	windowCMD   *tview.TextView
	papiClient  papi.Client
	currentUser user
}

type user struct {
	UserID      uuid.UUID               `json:"user_id"`
	PlayerName  string                  `json:"player_name"`
	EquipedBall ballType                `json:"equiped_pokeball"`
	PokeDex     map[string]papi.Pokemon `json:"pokedex"`
}

type gameData struct {
	GameID    int       `json:"game_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      user      `json:"user"`
}

type saveFile struct {
	SavedGameData map[int]gameData `json:"saved_game_data"`
}
