package day14

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func Test_Exo1_sample(t *testing.T) {
	Exo1("sample.txt")
}

func Test_Exo1_sample_nico(t *testing.T) {
	Exo1("sample_nico.txt")
}

func Test_Regex(t *testing.T) {
	re := regexp.MustCompile("AB")
	indexes := re.FindStringIndex("EEEABEEABGG")
	fmt.Println(indexes)
}

func Test_IndexStr(t *testing.T) {
	fmt.Println(strings.Index("BBB", "BB"))

}

func Test_MatchRule(t *testing.T) {
	indexes, toInsert := matchRule("NBBBCNCCNBBNBNBBCHBHHBCHB", "HB", "C")

	// Expected : 18, 21, 24

	fmt.Println("Insert ", toInsert, "at indexes:", indexes)
}

func Test_Exo1_input(t *testing.T) {
	Exo1("input.txt")
}
