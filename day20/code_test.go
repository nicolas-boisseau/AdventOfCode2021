package day20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process_sample(t *testing.T) {
	result := Process("sample.txt", 2)

	assert.Equal(t, 35, result)
}

func Test_Process_input(t *testing.T) {
	result := Process("input.txt", 2)

	assert.Equal(t, 5583, result)
}

func Test_Process_sample_complex(t *testing.T) {
	result := Process("sample.txt", 50)

	assert.Equal(t, 3351, result)
}

func Test_Process_input_complex(t *testing.T) {
	result := Process("input.txt", 50)

	assert.Equal(t, 19592, result)
}
