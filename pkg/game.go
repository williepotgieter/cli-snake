package pkg

import (
	"fmt"
	"os"
)

// NewGame returns a pointer to a new game
func (b *BoardInfo) NewGame() error {
	sg := SnakeGame{
		BoardInfo:    *b,
		CurrentRound: 1,
		Score:        0,
		GameOver:     false,
	}

	// Set food position
	sg.SetFood(sg.BoardInfo.FoodPos)

	// Set snake head position
	sg.SetSnake()

	if err := sg.SetGameState(); err != nil {
		return err
	}

	fmt.Printf("Created new cli-snake game with board size of %v rows and %v colums.\n", b.Size.Rows, b.Size.Columns)

	return nil
}

// SetFood maps the position of the food to the game board
func (s *SnakeGame) SetFood(p Position) {
	s.BoardInfo.Board[p.Row][p.Column] = "f"
}

// SetSnake maps the position of the snake to the game board
func (s *SnakeGame) SetSnake() {
	sn := s.BoardInfo.Snake
	for i := range sn.Body {
		// Set snake head
		if i != 0 {
			s.BoardInfo.Board[sn.Body[i].Row][sn.Body[i].Column] = "sb"
		} else {
			s.BoardInfo.Board[sn.Body[i].Row][sn.Body[i].Column] = "sh"
		}
	}
}

// CheckGameOver ends the game
func (s *SnakeGame) CheckGameOver() {
	// Check whether snake passes horizontal borders
	if (s.BoardInfo.Snake.Body[0].Row < 0) || (s.BoardInfo.Snake.Body[0].Row == s.BoardInfo.Size.Rows) {
		s.GameOver = true
	}
	// Check whether snake passes vertical borders
	if (s.BoardInfo.Snake.Body[0].Column < 0) || (s.BoardInfo.Snake.Body[0].Column == s.BoardInfo.Size.Columns) {
		s.GameOver = true
	}

	// Check whether snake head touches snake body
	for i := 1; i < len(s.BoardInfo.Snake.Body); i++ {
		if s.BoardInfo.Snake.Body[0] == s.BoardInfo.Snake.Body[i] {
			s.GameOver = true
		}
	}

	if s.GameOver == true {
		fmt.Println("GAME OVER!")
		fmt.Println("Score: ", s.Score)
		os.Remove("temp/gamestate.json")
		return
	}
}

// CheckFoundFood checks whether the snake has found food and adds length to the snake
func (s *SnakeGame) CheckFoundFood() error {
	if s.BoardInfo.Snake.Body[0] == s.BoardInfo.FoodPos {
		s.BoardInfo.Snake.Body = append([]Position{s.BoardInfo.FoodPos}, s.BoardInfo.Snake.Body...)
		s.SetSnake()
		pos := Position{
			Row:    realRandom(s.BoardInfo.Size.Rows),
			Column: realRandom(s.BoardInfo.Size.Columns),
		}
		// Check whether food position falls on snake body
		for hasPosition(s.BoardInfo.Snake.Body, pos) == true {
			pos = Position{
				Row:    realRandom(s.BoardInfo.Size.Rows),
				Column: realRandom(s.BoardInfo.Size.Columns),
			}
		}

		s.BoardInfo.FoodPos = pos

		s.Score++
		if err := s.SetGameState(); err != nil {
			return err
		}
	}

	return nil
}

func hasPosition(b []Position, p Position) bool {
	for _, pos := range b {
		if pos == p {
			return true
		}
	}

	return false
}
