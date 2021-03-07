package pkg

import "fmt"

// RenderBoard renders a snake board in the command line
func (s *SnakeGame) RenderBoard() {
	var sb = [10][10]string{
		{"", "", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", "f", "", ""},
		{"", "", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "sh", "", "", "", ""},
		{"", "", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", "", "", ""},
	}

	// Render top border
	fmt.Printf("%c", '┏')
	for i := 0; i < len(sb); i++ {
		fmt.Printf("%c", '━')
	}
	fmt.Printf("%c\n", '┓')

	// Render vertical borders and snake
	for i := 0; i < len(sb); i++ {
		fmt.Printf("%c", '┃')
		for j := 0; j < len(sb[0]); j++ {
			switch sb[i][j] {
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
	for i := 0; i < len(sb); i++ {
		fmt.Printf("%c", '━')
	}
	fmt.Printf("%c\n", '┛')

}

// RenderGameInfo renders round and score info about a snake game
func RenderGameInfo() {
	s := 0
	r := 1
	fmt.Printf("Round: %v	Score: %v\n", r, s)
}
