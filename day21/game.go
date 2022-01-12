package day21

type Player struct {
	id       int
	position int
	score    int
}

func newPlayer(id, position int) *Player {
	return &Player{
		id:       id,
		position: position,
	}
}

func (p *Player) HasWon(maxScore int) bool {
	return p.score >= maxScore
}

func (p *Player) Play(diceScore int) {
	for i := 0; i < diceScore; i++ {
		p.position++
		if p.position > 10 {
			p.position = 1
		}
	}

	p.score += p.position
}
