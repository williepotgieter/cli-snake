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

	sg.RenderBoard()

	return nil
}

// SetFood maps the position of the food to the game board
func (s *SnakeGame) SetFood(p Position) {
	s.BoardInfo.Board[p.Row][p.Column] = "f"
}

// SetSnake maps the position of the snake to the game board
func (s *SnakeGame) SetSnake() {
	sn := s.BoardInfo.Snake

	// Clear existing snake
	for i := 0; i < s.BoardInfo.Size.Rows; i++ {
		for j := 0; j < s.BoardInfo.Size.Columns; j++ {
			if s.BoardInfo.Board[i][j] != "f" {
				s.BoardInfo.Board[i][j] = ""
			}
		}
	}

	// Set updated snake
	for i := range sn.Body {
		if i == 0 {
			s.BoardInfo.Board[sn.Body[i].Row][sn.Body[i].Column] = "sh"
		} else {
			s.BoardInfo.Board[sn.Body[i].Row][sn.Body[i].Column] = "sb"
		}
	}
}

// CheckGameOver ends the game
func (s *SnakeGame) CheckGameOver(p Position) {
	// Check whether snake passes horizontal borders
	if (p.Row < 0) || (p.Row == s.BoardInfo.Size.Rows) {
		s.GameOver = true
	}
	// Check whether snake passes vertical borders
	if (p.Column < 0) || (p.Column == s.BoardInfo.Size.Columns) {
		s.GameOver = true
	}

	// Check whether snake head touches snake body
	for i := 1; i < len(s.BoardInfo.Snake.Body); i++ {
		if s.BoardInfo.Snake.Body[0] == s.BoardInfo.Snake.Body[i] {
			s.GameOver = true
		}
	}
}

// CheckFoundFood checks whether the snake has found food and adds length to the snake
func (s *SnakeGame) CheckFoundFood(np Position) error {

	if np == s.BoardInfo.FoodPos {
		s.BoardInfo.Snake.Body = append([]Position{np}, s.BoardInfo.Snake.Body...)

		pos := Position{
			Row:    RealRandom(s.BoardInfo.Size.Rows),
			Column: RealRandom(s.BoardInfo.Size.Columns),
		}
		// Check whether food position falls on snake body
		for HasPosition(s.BoardInfo.Snake.Body, pos) == true {
			pos = Position{
				Row:    RealRandom(s.BoardInfo.Size.Rows),
				Column: RealRandom(s.BoardInfo.Size.Columns),
			}
		}

		s.BoardInfo.FoodPos = pos

		s.Score++
		if err := s.SetGameState(); err != nil {
			return err
		}
	} else {
		s.BoardInfo.Snake.Body = []Position{np}
	}

	s.SetSnake()

	return nil
}

// EndGame ends the current came and removes temporary game state directory
func (s *SnakeGame) EndGame() {
	fmt.Println("GAME OVER!")
	fmt.Println("Score: ", s.Score)
	os.Remove("temp/gamestate.json")
}

// MoveNext moves the snake's head to the next position
func (s *SnakeGame) MoveNext(np Position) {
	// Check if game over
	s.CheckGameOver(np)
	if s.GameOver == true {
		s.EndGame()
	}

	// Check if next pos is food
	if np == s.BoardInfo.FoodPos {
		s.BoardInfo.Snake.Body = append([]Position{np}, s.BoardInfo.Snake.Body...)
	} else {
		s.BoardInfo.Snake.Body = []Position{np}
	}

	s.SetSnake()
}
