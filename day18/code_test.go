package day18

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process_sample(t *testing.T) {
	result := Process("sample.txt")

	assert.Equal(t, 143, result)
}

func Test_Process_input(t *testing.T) {
	result := Process("input.txt")

	fmt.Println("Result:", result)
}
