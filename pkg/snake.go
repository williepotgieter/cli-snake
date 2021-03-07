package pkg

// BoardInfo type to store information about board
type BoardInfo struct {
	Size    BoardSize
	Board   [][]string
	Snake   Snake
	FoodPos Position
}

// BoardSize type to store board size
type BoardSize struct {
	Rows    int
	Columns int
}

// Position type to store a position on the board
type Position struct {
	Row    int
	Column int
}

// Snake type stores details about the current snake
type Snake struct {
	Body []Position
}

// SnakeGame to store state of snake game
type SnakeGame struct {
	BoardInfo    BoardInfo
	CurrentRound int
	Score        int
	GameOver     bool
}
