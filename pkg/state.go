package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// SetGameState saves the current game state as json encoded environment variable
func (s *SnakeGame) SetGameState() error {
	bytes, err := json.Marshal(s)
	if err != nil {
		return err
	}
	jsonStr := string(bytes)
	fmt.Println(jsonStr)
	if err := os.Setenv("CLISNAKE_GAME", jsonStr); err != nil {
		return err
	}

	return nil
}

// GetGameState gets the current game state from the "CLISNAKE_GAME" environment variable
func GetGameState() (*SnakeGame, error) {
	var sg SnakeGame
	g, err := os.LookupEnv("CLISNAKE_GAME")
	if err == false {
		return &SnakeGame{}, errors.New("no saved game data available")
	}

	if err := json.Unmarshal([]byte(g), &sg); err != nil {
		return &SnakeGame{}, errors.New("corrupted game data. start new game to resolve")
	}

	return &sg, nil
}
