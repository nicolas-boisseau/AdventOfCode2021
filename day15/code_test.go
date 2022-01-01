package day15

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process_Exo1_sample0(t *testing.T) {
	computed := Process("sample0.txt", 2, false)

	assert.Equal(t, 7, computed)
}

func Test_Process_Exo1_sample(t *testing.T) {
	computed := Process("sample.txt", 10, false)

	assert.Equal(t, 40, computed)
}

func Test_Process_Exo1_input(t *testing.T) {
	computed := Process("input.txt", 100, false)

	assert.Equal(t, 527, computed)
}

func Test_Process_Exo2_sample(t *testing.T) {
	computed := Process("sample.txt", 10, true)

	assert.Equal(t, 315, computed)
}

func Test_Process_Exo2_input(t *testing.T) {
	computed := Process("input.txt", 100, true)

	assert.Equal(t, 315, computed)
}
