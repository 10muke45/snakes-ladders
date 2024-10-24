package player

type Player struct {
	Position int
}

func (p *Player) Move(diceNum int, boardSize int) {
	currPos := p.Position + diceNum
	if currPos > 100 {
		return
	}
	p.Position += diceNum
}
