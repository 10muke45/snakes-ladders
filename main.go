package main

import (
	board "snake-ladder/board"
	"snake-ladder/game"
	ladder "snake-ladder/ladders"
	"snake-ladder/player"
	snake "snake-ladder/snakes"
)

func main() {
	player1 := &player.Player{Position: 0}
	player2 := &player.Player{Position: 0}

	snakes := []snake.Snake{
		{
			Start: 11,
			End:   7,
		},
		{
			Start: 31,
			End:   25,
		},
		{
			Start: 67,
			End:   44,
		},
		{
			Start: 87,
			End:   70,
		},
		{
			Start: 92,
			End:   33,
		},
	}
	ladders := []ladder.Ladder{
		{
			Start: 10,
			End:   50,
		},
		{
			Start: 23,
			End:   86,
		},
		{
			Start: 45,
			End:   49,
		},
		{
			Start: 61,
			End:   79,
		},
		{
			Start: 77,
			End:   91,
		},
	}
	gameBoard := board.Board{
		Size:    100,
		Snakes:  snakes,
		Ladders: ladders,
	}
	players := []*player.Player{
		player1, player2,
	}
	game := game.NewGame(gameBoard, players)
	for !game.IsGameOver {
		game.PlayTurn()
	}
}
