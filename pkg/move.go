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
		if err := s.moveUp(); err != nil {
			return err
		}
	case "down":
		if err := s.moveDown(); err != nil {
			return err
		}
	case "left":
		if err := s.moveLeft(); err != nil {
			return err
		}
	case "right":
		if err := s.moveRight(); err != nil {
			return err
		}
	default:
		return errors.New("invalid move. enter only up, down, left or right")
	}

	return nil
}

func (s *SnakeGame) moveUp() error {
	currentPos := Position{
		Row:    s.BoardInfo.Snake.Body[0].Row,
		Column: s.BoardInfo.Snake.Body[0].Column,
	}

	nextPos := Position{
		Row:    currentPos.Row - 1,
		Column: currentPos.Column,
	}

	if err := s.makeMove(nextPos); err != nil {
		return err
	}

	return nil
}

func (s *SnakeGame) moveDown() error {
	currentPos := Position{
		Row:    s.BoardInfo.Snake.Body[0].Row,
		Column: s.BoardInfo.Snake.Body[0].Column,
	}

	nextPos := Position{
		Row:    currentPos.Row + 1,
		Column: currentPos.Column,
	}

	if err := s.makeMove(nextPos); err != nil {
		return err
	}

	return nil
}

func (s *SnakeGame) moveLeft() error {
	currentPos := Position{
		Row:    s.BoardInfo.Snake.Body[0].Row,
		Column: s.BoardInfo.Snake.Body[0].Column,
	}

	nextPos := Position{
		Row:    currentPos.Row,
		Column: currentPos.Column - 1,
	}

	if err := s.makeMove(nextPos); err != nil {
		return err
	}

	return nil
}

func (s *SnakeGame) moveRight() error {
	currentPos := Position{
		Row:    s.BoardInfo.Snake.Body[0].Row,
		Column: s.BoardInfo.Snake.Body[0].Column,
	}

	nextPos := Position{
		Row:    currentPos.Row,
		Column: currentPos.Column + 1,
	}

	if err := s.makeMove(nextPos); err != nil {
		return err
	}

	return nil
}

func (s *SnakeGame) makeMove(np Position) error {
	s.MoveNext(np)
	s.CheckGameOver(np)

	if s.GameOver == true {
		if err := s.EndGame(); err != nil {
			return err
		}
		return nil
	}

	//s.MoveNext(np)

	s.CurrentRound++

	s.SetGameState()

	s.RenderBoard()
	return nil
}
