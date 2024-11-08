package main

import (
	"encoding/json"
	"os"
)

func commandLoad(cfg *config, args ...string) string {
	jsonData, err := os.ReadFile("saves/saveFile.json")

	if err != nil {
		return check(err)
	}

	loadUser := user{}

	err = json.Unmarshal(jsonData, &loadUser)

	if err != nil {
		return check(err)
	}
	cfg.currentUser = loadUser
	return "Save file loaded\n"
}
