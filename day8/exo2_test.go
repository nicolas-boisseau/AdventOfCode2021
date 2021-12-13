package day8

import (
	"fmt"
	"testing"
)

func Test_FindCommonChars(t *testing.T) {

	chars, n := findCommonChars("be", "cfbegad")
	fmt.Println(string(chars))
	fmt.Println(n)

}

func Test_pass1(t *testing.T) {
	fmt.Println("should be 1:", pass1("ab"))
	fmt.Println("should be 7:", pass1("abd"))
	fmt.Println("should be 4:", pass1("abeg"))
	fmt.Println("should be 8:", pass1("abcdefg"))
}
