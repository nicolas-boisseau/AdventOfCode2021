package day17

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process_sample(t *testing.T) {
	maxY, nbShots := Process("sample.txt")

	assert.Equal(t, 45, maxY)
	assert.Equal(t, 112, nbShots)
}

func Test_Process_input(t *testing.T) {
	x, y := Process("input.txt")

	fmt.Println("Result:", x, ",", y)
}
