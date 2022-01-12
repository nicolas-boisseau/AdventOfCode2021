package day21

type DeterministicDice struct {
	currentValue int
	faces        int
	rollCount    int
}

func newDeterministicDice(faces int) *DeterministicDice {
	return &DeterministicDice{
		currentValue: 0,
		faces:        faces,
	}
}

func (dice *DeterministicDice) Roll() int {
	dice.currentValue = dice.currentValue + 1
	if dice.currentValue > 100 {
		dice.currentValue = 1
	}

	dice.rollCount++

	return dice.currentValue
}
