package day19

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process_sample0(t *testing.T) {
	result := Process("sample0.txt")

	assert.Equal(t, 3, result)
}

func Test_Process_sample(t *testing.T) {
	result := Process("sample.txt")

	assert.Equal(t, 79, result)
}
