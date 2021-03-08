package pkg

import (
	"errors"
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
	currentPos := Position{
		Row:    s.BoardInfo.Snake.Body[0].Row,
		Column: s.BoardInfo.Snake.Body[0].Column,
	}

	nextPos := Position{
		Row:    currentPos.Row - 1,
		Column: currentPos.Column,
	}

	s.makeMove(nextPos)
}

func (s *SnakeGame) moveDown() {
	currentPos := Position{
		Row:    s.BoardInfo.Snake.Body[0].Row,
		Column: s.BoardInfo.Snake.Body[0].Column,
	}

	nextPos := Position{
		Row:    currentPos.Row + 1,
		Column: currentPos.Column,
	}

	s.makeMove(nextPos)
}

func (s *SnakeGame) moveLeft() {
	currentPos := Position{
		Row:    s.BoardInfo.Snake.Body[0].Row,
		Column: s.BoardInfo.Snake.Body[0].Column,
	}

	nextPos := Position{
		Row:    currentPos.Row,
		Column: currentPos.Column - 1,
	}

	s.makeMove(nextPos)
}

func (s *SnakeGame) moveRight() {
	currentPos := Position{
		Row:    s.BoardInfo.Snake.Body[0].Row,
		Column: s.BoardInfo.Snake.Body[0].Column,
	}

	nextPos := Position{
		Row:    currentPos.Row,
		Column: currentPos.Column + 1,
	}

	s.makeMove(nextPos)
}

func (s *SnakeGame) makeMove(np Position) {
	s.CheckGameOver(np)

	if s.GameOver == true {
		s.EndGame()
		return
	}

	s.MoveNext(np)

	s.CurrentRound++

	s.SetGameState()

	s.RenderBoard()
}
