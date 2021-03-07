package pkg

import (
	"encoding/json"
	"errors"
	"os"
)

// SetGameState saves the current game state as json encoded environment variable
func (s *SnakeGame) SetGameState() {
	bytes, _ := json.Marshal(s)
	jsonStr := string(bytes)
	os.Setenv("CLISNAKE_GAME", jsonStr)
}

// GetGameState gets the current game state from the "CLISNAKE_GAME" environment variable
func GetGameState() (*SnakeGame, error) {
	var sg SnakeGame
	g, err := os.LookupEnv("CLISNAKE_GAME")
	if err == false {
		return &SnakeGame{}, errors.New("no saved current game")
	}

	if err := json.Unmarshal([]byte(g), &sg); err != nil {
		return &SnakeGame{}, errors.New("corrupted game data. start new game to resolve")
	}

	return &sg, nil
}
