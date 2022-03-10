package day22

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process_sample0(t *testing.T) {
	result := Process("sample0.txt", false)

	assert.Equal(t, 39, result)
}

func Test_Process_sample(t *testing.T) {
	result := Process("sample.txt", false)

	assert.Equal(t, 590784, result)
}

func Test_Process_input(t *testing.T) {
	result := Process("input.txt", false)

	assert.Equal(t, int64(582644), result)
}

func Test_Process_sample0_complex(t *testing.T) {
	result := Process("sample0.txt", true)

	var expected int64 = 39

	assert.Equal(t, expected, result)
}

func Test_Process_sample_complex(t *testing.T) {
	result := Process("sample_complex.txt", true)

	var expected int64 = 2758514936282235

	assert.Equal(t, expected, result)
}

func Test_Process_input_complex(t *testing.T) {
	result := Process("input.txt", true)

	var expected int64 = 131888061854776

	assert.Equal(t, expected, result)
}
