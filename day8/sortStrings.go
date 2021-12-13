package day8

import (
	"sort"
)

func SortString(input string) string {
	runeArray := []rune(input)
	sort.Sort(sortRuneString(runeArray))
	//fmt.Println(string(runeArray))
	return string(runeArray)
}

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}
