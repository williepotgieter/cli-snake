package pkg

import "fmt"

// RenderBoard renders a snake board in the command line
func (s *SnakeGame) RenderBoard() {
	r := s.BoardInfo.Size.Rows
	c := s.BoardInfo.Size.Columns
	b := s.BoardInfo.Board

	// Render top border
	fmt.Printf("%c", '┏')
	for i := 0; i < c; i++ {
		fmt.Printf("%c", '━')
	}
	fmt.Printf("%c\n", '┓')

	// Render vertical borders and snake
	for i := 0; i < r; i++ {
		fmt.Printf("%c", '┃')
		for j := 0; j < c; j++ {
			switch b[i][j] {
			case "sh":
				fmt.Printf("%c", '◯')
			case "sb":
				fmt.Printf("%c", '▫')
			case "f":
				fmt.Printf("%c", '◈')
			default:
				fmt.Printf(" ")
			}
		}
		fmt.Printf("%c\n", '┃')
	}

	// Render bottom border
	fmt.Printf("%c", '┗')
	for i := 0; i < c; i++ {
		fmt.Printf("%c", '━')
	}
	fmt.Printf("%c\n", '┛')

	s.RenderGameInfo()
}

// RenderGameInfo renders round and score info about a snake game
func (s *SnakeGame) RenderGameInfo() {
	fmt.Printf("Round: %v	Score: %v\n", s.CurrentRound, s.Score)
}
