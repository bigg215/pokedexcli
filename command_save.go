package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/rivo/tview"
)

func commandSave(cfg *config, args ...string) string {
	curr := cfg.windowCMD.GetText(false)
	saveData := saveFile{
		SavedGameData: make(map[int]gameData),
	}

	err := loadGameFile(&saveData)

	if err != nil {
		cfg.windowCMD.SetText(curr + fmt.Sprintf("%s\n", err))
		cfg.windowCMD.ScrollToEnd()
		cfg.appConfig.SetRoot(cfg.windowRoot, true)
	}

	saveList := tview.NewList()

	for i := 0; i < len(ALPHA); i++ {
		title := fmt.Sprintf("[%d] New Save", i)
		desc := ""
		saveCreatedTime := time.Now().UTC()

		if val, ok := saveData.SavedGameData[i]; ok {
			title = fmt.Sprintf("[%d] - user: %s", i, val.User.UserID)
			desc = fmt.Sprintf("Created: %s - Updated: %s", val.CreatedAt.Local().Format("2006-01-02 15:04:05"), val.UpdatedAt.Local().Format("2006-01-02 15:04:05"))
			saveCreatedTime = val.CreatedAt
		}

		saveList.AddItem(title, desc, rune(ALPHA[i]), func() {
			saveData.SavedGameData[i] = gameData{
				GameID:    i,
				CreatedAt: saveCreatedTime,
				UpdatedAt: time.Now().UTC(),
				User:      cfg.currentUser,
			}
			err := saveGameFile(saveData)
			if err != nil {
				cfg.windowCMD.SetText(curr + fmt.Sprintf("%s\n", err))
			}
			cfg.windowCMD.SetText(curr + "game saved\n")
			cfg.windowCMD.ScrollToEnd()
			cfg.appConfig.SetRoot(cfg.windowRoot, true)
		})
	}

	centerGrid := centeredGrid(saveList, 65, 0)

	cfg.appConfig.SetRoot(centerGrid, true).SetFocus(saveList)
	return ""
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
