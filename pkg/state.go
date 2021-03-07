package pkg

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// SetGameState saves the current game state as json encoded environment variable
func (s *SnakeGame) SetGameState() error {
	jsonStr, err := json.Marshal(s)
	if err != nil {
		return err
	}

	ioutil.WriteFile("gamestate.json", jsonStr, os.ModePerm)

	return nil
}

// GetGameState gets the current game state from the "CLISNAKE_GAME" environment variable
func GetGameState() (*SnakeGame, error) {
	var sg SnakeGame

	file, err := ioutil.ReadFile("gamestate.json")
	if err != nil {
		return &SnakeGame{}, err
	}

	if err := json.Unmarshal(file, &sg); err != nil {
		return &SnakeGame{}, err
	}

	return &sg, nil
}
