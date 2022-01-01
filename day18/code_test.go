package day18

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process_sample(t *testing.T) {
	result := Process("sample.txt")

	assert.Equal(t, 4140, result)
}

func Test_Process_input(t *testing.T) {
	result := Process("input.txt")

	fmt.Println("Result:", result)
}

func Test_Process2_sample(t *testing.T) {
	result := Process2("sample.txt")

	assert.Equal(t, 3993, result)
}

func Test_Process2_input(t *testing.T) {
	result := Process2("input.txt")

	fmt.Println("Result:", result)
}
