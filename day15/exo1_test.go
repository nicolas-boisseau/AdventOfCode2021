package day15

import (
	"testing"
)

func Test_Exo1_sample(t *testing.T) {
	Exo1("sample.txt", 10, "0_0", "9_9", false)
}

func Test_Exo2_sample(t *testing.T) {
	Exo1("sample.txt", 10, "0_0", "49_49", true)
}

func Test_Exo1_sample0(t *testing.T) {
	Exo1("sample0.txt", 2, "0_0", "1_1", false)
}

func Test_Exo1_input(t *testing.T) {
	Exo1("input.txt", 100, "0_0", "99_99", false)
}

func Test_Exo2_input(t *testing.T) {
	Exo1("input.txt", 100, "0_0", "499_499", true)
}
