package pkg

import (
	"github.com/urfave/cli/v2"
)

// CreateBoard creates a new snake board
func CreateBoard(c *cli.Context) (BoardInfo, error) {
	bs, err := ValidateBoardSize(c)
	if err != nil {
		return BoardInfo{}, err
	}

	// Initialize empty board
	b := make([][]string, bs.Rows)
	for i := range b {
		b[i] = make([]string, bs.Columns)
	}

	// Generate random food position
	fp := Position{
		Row:    realRandom(bs.Rows),
		Column: realRandom(bs.Columns),
	}

	// Generate random position for snake's head
	sh := Position{
		Row:    realRandom(bs.Rows),
		Column: realRandom(bs.Columns),
	}

	// Ensure that snake head position and food position is different
	for sh == fp {
		sh = Position{
			Row:    realRandom(bs.Rows),
			Column: realRandom(bs.Columns),
		}
	}

	sn := Snake{
		Body: []Position{sh},
	}

	return BoardInfo{
		Size:    bs,
		Board:   b,
		Snake:   sn,
		FoodPos: fp,
	}, nil
}
