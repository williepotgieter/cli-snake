package pkg

import (
	"errors"
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

// MoveSnake moves the snake head one position
func (s *SnakeGame) MoveSnake(c *cli.Context) error {
	move := strings.ToLower(c.String("move"))

	switch move {
	case "up":
		s.moveUp()
	case "down":
		s.moveDown()
	case "left":
		s.moveLeft()
	case "right":
		s.moveRight()
	default:
		return errors.New("invalid move. enter only up, down, left or right")
	}

	return nil
}

func (s *SnakeGame) moveUp() {
	fmt.Println("Moved up!")
}

func (s *SnakeGame) moveDown() {
	fmt.Println("Moved down!")
}

func (s *SnakeGame) moveLeft() {
	fmt.Println("Moved left!")
}

func (s *SnakeGame) moveRight() {
	fmt.Println("Moved right!")
}
