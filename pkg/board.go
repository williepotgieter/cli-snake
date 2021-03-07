package pkg

import (
	"math/rand"
	"time"

	"github.com/urfave/cli/v2"
)

// CreateBoard creates a new snake board
func CreateBoard(c *cli.Context) (BoardInfo, error) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

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
		Row:    r1.Intn(bs.Rows),
		Column: r1.Intn(bs.Columns),
	}

	// Generate random position for snake's head
	sh := Position{
		Row:    r1.Intn(bs.Rows),
		Column: r1.Intn(bs.Columns),
	}

	// Ensure that snake head position and food position is different
	for sh == fp {
		sh = Position{
			Row:    r1.Intn(bs.Rows),
			Column: r1.Intn(bs.Columns),
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
