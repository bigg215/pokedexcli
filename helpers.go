package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/rivo/tview"
)

const (
	ALPHA string = "abcde"
)

func check(e error) string {
	return fmt.Sprintf("%s", e)
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func centeredGrid(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
}

func saveGameFile(T interface{}) error {
	jsonData, err := json.MarshalIndent(T, "", "  ")

	if err != nil {
		return err
	}

	file, err := os.Create("saves/saveFile.json")

	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)

	if err != nil {
		return err
	}

	file.Sync()

	return nil
}

func loadGameFile(T interface{}) error {
	jsonData, err := os.ReadFile("saves/saveFile.json")

	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, &T)

	if err != nil {
		return err
	}
	return nil
}
