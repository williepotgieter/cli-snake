package pkg

import "fmt"

// MoveUp moves the snake head one position up
func (s *SnakeGame) MoveUp() {
	fmt.Println("Moved up!")
}

// MoveDown moves the snake head one position down
func (s *SnakeGame) MoveDown() {
	fmt.Println("Moved down!")
}

// MoveLeft moves the snake head one position left
func (s *SnakeGame) MoveLeft() {
	fmt.Println("Moved left!")
}

// MoveRight moves the snake head one position right
func (s *SnakeGame) MoveRight() {
	fmt.Println("Moved right!")
}
