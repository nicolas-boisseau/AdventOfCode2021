package day21

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RollDeterministicDice(t *testing.T) {

	d := newDeterministicDice(100)

	assert.Equal(t, 1, d.Roll())
	assert.Equal(t, 2, d.Roll())
	assert.Equal(t, 3, d.Roll())
	assert.Equal(t, 4, d.Roll())
	assert.Equal(t, 5, d.Roll())

	for i := 0; i < 93; i++ {
		d.Roll()
	}

	assert.Equal(t, 99, d.Roll())
	assert.Equal(t, 100, d.Roll())
	assert.Equal(t, 1, d.Roll())
}
