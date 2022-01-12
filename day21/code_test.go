package day21

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process_sample(t *testing.T) {
	result := Process("sample.txt")

	assert.Equal(t, 739785, result)
}

func Test_Process_input(t *testing.T) {
	result := Process("input.txt")

	assert.Equal(t, 576600, result)
}

func Test_Process_sample_complex(t *testing.T) {
	result := Process2("sample.txt")

	var expected int64 = 444356092776315

	assert.Equal(t, expected, result)
}

func Test_Process_input_complex(t *testing.T) {
	result := Process2("input.txt")

	var expected int64 = 131888061854776

	assert.Equal(t, expected, result)
}

func Test_Chances(t *testing.T) {
	chancesByNumbers := make(map[int]int)
	for d1 := 1; d1 <= 3; d1++ {
		for d2 := 1; d2 <= 3; d2++ {
			for d3 := 1; d3 <= 3; d3++ {
				chancesByNumbers[d1+d2+d3]++
			}
		}
	}
	fmt.Println(chancesByNumbers)
}
