package main

import (
	"encoding/json"
	"os"
)

func commandSave(cfg *config, args ...string) string {
	jsonData, err := json.MarshalIndent(cfg.currentUser, "", "  ")

	if err != nil {
		return check(err)
	}

	file, err := os.Create("saves/saveFile.json")

	if err != nil {
		return check(err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)

	if err != nil {
		return check(err)
	}

	file.Sync()

	return "game saved\n"
}
