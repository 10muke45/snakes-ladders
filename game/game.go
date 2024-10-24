package game

import (
	"fmt"
	"snake-ladder/board"
	"snake-ladder/dice"
	"snake-ladder/player"
)

type Game struct {
	Board         board.Board
	Players       []*player.Player
	IsGameOver    bool
	CurrentPlayer int
}

func NewGame(board board.Board, players []*player.Player) *Game {
	return &Game{
		Board:         board,
		Players:       players,
		IsGameOver:    false,
		CurrentPlayer: 0,
	}
}

func (g *Game) PlayTurn() {
	var x string
	diceNum := dice.Roll()
	x = fmt.Sprintf("%d player rolled %d", g.CurrentPlayer, diceNum)
	fmt.Println(x)
	g.Players[g.CurrentPlayer].Move(diceNum, g.Board.Size)
	newPosition := g.Board.GetNewPosition(g.Players[g.CurrentPlayer].Position)
	g.Players[g.CurrentPlayer].Position = newPosition
	x = fmt.Sprintf("player %d in %d position", g.CurrentPlayer, g.Players[g.CurrentPlayer].Position)
	fmt.Println(x)
	if g.Players[g.CurrentPlayer].Position == g.Board.Size {
		g.IsGameOver = true
		x = fmt.Sprintf("player %d won the game", g.CurrentPlayer)
		fmt.Println(x)
		return
	}
	g.PlayNext(g.CurrentPlayer, len(g.Players))
}

func (g *Game) PlayNext(currentPlayer int, noOfPlayers int) {
	currentPlayer = (currentPlayer + 1) % noOfPlayers
	g.CurrentPlayer = currentPlayer
}
