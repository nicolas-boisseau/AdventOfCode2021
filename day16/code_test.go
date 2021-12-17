package day16

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Exo1_sample(t *testing.T) {
	versionSum, computed := Process("sample.txt")

	assert.Equal(t, 12, versionSum)
	assert.Equal(t, int64(46), computed)
}

func Test_Exo1_input(t *testing.T) {
	versionSum, computed := Process("input.txt")

	fmt.Println("Version sum:", versionSum, ", Computed:", computed)
}

func Test_binToHexa(t *testing.T) {
	fmt.Println(binToHexa("110100101111111000101000"))
}
