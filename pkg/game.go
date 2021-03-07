package pkg

import (
	"fmt"
)

// NewGame returns a pointer to a new game
func (b *BoardInfo) NewGame() error {
	sg := SnakeGame{
		BoardInfo:    *b,
		CurrentRound: 1,
		Score:        0,
		GameOver:     false,
	}

	if err := sg.SetGameState(); err != nil {
		return err
	}

	fmt.Printf("Created new cli-snake game with board size of %v rows and %v colums.\n", b.Size.Rows, b.Size.Columns)

	return nil
}
