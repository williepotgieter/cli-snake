package pkg

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// SetGameState saves the current game state as json encoded environment variable
func (s *SnakeGame) SetGameState() error {
	jsonStr, err := json.Marshal(s)
	if err != nil {
		return err
	}

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	newpath := filepath.Join(path, "temp")
	if _, err := os.Stat(newpath); os.IsNotExist(err) {
		os.Mkdir(newpath, 0755)
	}

	ioutil.WriteFile("temp/gamestate.json", jsonStr, os.ModePerm)

	return nil
}

// GetGameState gets the current game state from the "CLISNAKE_GAME" environment variable
func GetGameState() (*SnakeGame, error) {
	var sg SnakeGame

	file, err := ioutil.ReadFile("temp/gamestate.json")
	if err != nil {
		return &SnakeGame{}, err
	}

	if err := json.Unmarshal(file, &sg); err != nil {
		return &SnakeGame{}, err
	}

	return &sg, nil
}
