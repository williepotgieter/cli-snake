package pkg

// NewGame returns a pointer to a new game
func (b *BoardInfo) NewGame() *SnakeGame {
	return &SnakeGame{
		BoardInfo:    *b,
		CurrentRound: 1,
		Score:        0,
		GameOver:     false,
	}
}
