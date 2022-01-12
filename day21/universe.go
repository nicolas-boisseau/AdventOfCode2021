package day21

type Universe struct {
	players         map[int]*Player
	multiple        int64
	currentPlayerId int
	winnerId        int
}

func (u *Universe) Play(diceTotalScore int, multiple int64) *Universe {

	pos := u.players[u.currentPlayerId].position
	score := u.players[u.currentPlayerId].score

	for i := 0; i < diceTotalScore; i++ {
		pos++
		if pos > 10 {
			pos = 1
		}
	}

	score += pos

	outputU := &Universe{
		players:  make(map[int]*Player),
		multiple: u.multiple * multiple,
	}

	otherPlayerId := 2
	if u.currentPlayerId == 2 {
		otherPlayerId = 1
	}

	outputU.players[u.currentPlayerId] = &Player{
		id:       u.currentPlayerId,
		position: pos,
		score:    score,
	}
	outputU.players[otherPlayerId] = u.players[otherPlayerId]

	if outputU.players[u.currentPlayerId].HasWon(21) {
		outputU.winnerId = u.currentPlayerId
	}

	outputU.currentPlayerId = otherPlayerId

	return outputU
}
