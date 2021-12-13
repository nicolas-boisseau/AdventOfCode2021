package day4

import (
	"AdventOfCode2021/common"
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"strconv"
	"strings"
)

// https://pkg.go.dev/github.com/ahmetb/go-linq

func Exo2() {
	lines := common.ReadLinesFromFile("day4/input.txt")

	fmt.Println("lines[0]:", lines[0])

	// Extract bingo random numbers
	bingoNumbersStr := strings.Split(lines[0], ",")
	var bingoNumbers []int = make([]int, 0)
	for _, bingoNumberStr := range bingoNumbersStr {
		bingoNumber, _ := strconv.Atoi(bingoNumberStr)
		bingoNumbers = append(bingoNumbers, bingoNumber)
	}

	boards := make([]*BoardSet, 0)
	for i := 2; i < len(lines); i += 6 {
		boardLines := lines[i : i+5]
		boardMatrix := ExtractMatrix(boardLines, 5, 5)
		board := newBoardSet(boardMatrix, 5, 5)
		boards = append(boards, board)
	}

	fmt.Println("Read all board!")

	fmt.Println("Play !")

	// Launch game !
	boardAlreadyWons := make([]int, 0)
	for _, bingoNumber := range bingoNumbers {
		for i, b := range boards {
			b.play(bingoNumber)

			won, serie := b.won()
			if won && !From(boardAlreadyWons).AnyWithT(func(bb int) bool { return bb == i }) {
				fmt.Println("Gagné !! board n°", i, "avec une suite:", serie)
				sum := b.unmarkedNumbersSum()
				fmt.Println("Résulat = ", bingoNumber*sum)
				boardAlreadyWons = append(boardAlreadyWons, i)
			}
		}
	}
}
