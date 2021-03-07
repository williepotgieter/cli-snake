package pkg

import (
	"math/rand"

	"github.com/urfave/cli/v2"
)

// CreateBoard creates a new snake board
func CreateBoard(c *cli.Context) (BoardInfo, error) {
	bs, err := ValidateBoardSize(c)
	if err != nil {
		return BoardInfo{}, err
	}

	b := make([][]string, bs.Rows, bs.Columns)

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			b[i][j] = ""
		}
	}

	// Generate random food position
	fp := Position{
		Row:    rand.Intn(bs.Rows),
		Column: rand.Intn(bs.Columns),
	}

	// Generate random position for snake's head
	sh := Position{
		Row:    rand.Intn(bs.Rows),
		Column: rand.Intn(bs.Columns),
	}

	// Ensure that snake head position and food position is different
	for sh == fp {
		sh = Position{
			Row:    rand.Intn(bs.Rows),
			Column: rand.Intn(bs.Columns),
		}
	}

	sn := Snake{
		Body:   []Position{sh},
		Length: 1,
	}

	return BoardInfo{
		Size:    bs,
		Board:   b,
		Snake:   sn,
		FoodPos: fp,
	}, nil
}
