package board

import (
	"fmt"
	ladder "snake-ladder/ladders"
	snake "snake-ladder/snakes"
)

type Board struct {
	Size          int
	IsGameOver    bool
	Snakes        []snake.Snake
	Ladders       []ladder.Ladder
	CurrentPlayer int
}

func (b *Board) GetNewPosition(currPos int) int {
	var x string
	for _, snake := range b.Snakes {
		if snake.Start == currPos {
			x = fmt.Sprintf("player moved down from %d to %d", snake.Start, snake.End)
			fmt.Println(x)
			currPos = snake.End
		}
	}
	for _, ladder := range b.Ladders {
		if ladder.Start == currPos {
			x = fmt.Sprintf("player moved up from %d to %d", ladder.Start, ladder.End)
			fmt.Println(x)
			currPos = ladder.End
		}
	}
	return currPos
}
