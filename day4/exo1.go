package day4

import (
	"AdventOfCode2021/common"
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"strconv"
	"strings"
)

// https://pkg.go.dev/github.com/ahmetb/go-linq

func ExtractMatrix(lines []string, h int, w int) [][]int {
	// init matrix H x W
	matrixOutput := make([][]int, h)
	//for i := range matrixOutput {
	//	matrixOutput[i] = make([]int, w)
	//}

	// Extract numbers and populate the matrix
	for i := 0; i < h; i++ {
		bingoNumbersStr := strings.Split(lines[i], " ")
		boardLineNumbers := make([]int, 0)
		for _, bingoNumberStr := range bingoNumbersStr {
			if bingoNumberStr != "" {
				bingoNumber, _ := strconv.Atoi(bingoNumberStr)
				boardLineNumbers = append(boardLineNumbers, bingoNumber)
			}
		}
		matrixOutput[i] = boardLineNumbers
	}

	return matrixOutput
}

type BoardSet struct {
	board  [][]int
	checks [][]bool
	h      int
	w      int
}

func newBoardSet(board [][]int, h int, w int) *BoardSet {
	b := BoardSet{board: board}
	b.h = h
	b.w = w
	b.checks = make([][]bool, h)
	for i := range b.checks {
		b.checks[i] = make([]bool, w)
	}
	return &b
}

func (b *BoardSet) play(bingoNumber int) {
	for i, _ := range b.board {
		for j, _ := range b.board[i] {
			if b.board[i][j] == bingoNumber {
				b.checks[i][j] = true
			}
		}
	}
}

func (b *BoardSet) unmarkedNumbersSum() int {
	sum := 0
	for i, _ := range b.board {
		for j, _ := range b.board[i] {
			if !b.checks[i][j] {
				sum += b.board[i][j]
			}
		}
	}
	return sum
}

func (b *BoardSet) won() (bool, []int) {
	for i, _ := range b.checks {

		// row
		if From(b.checks[i]).AllT(func(check bool) bool { return check }) {
			return true, b.board[i]
		}

		// col
		col := make([]int, 0)
		colChecks := make([]bool, 0)
		for j := 0; j < 5; j++ {
			col = append(col, b.board[j][i])
			colChecks = append(colChecks, b.checks[j][i])
		}
		if From(colChecks).AllT(func(check bool) bool { return check }) {
			return true, col
		}
	}

	return false, nil
}

func Exo1() {
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
	finished := false

	// Launch game !
	for _, bingoNumber := range bingoNumbers {
		for i, b := range boards {
			b.play(bingoNumber)

			won, serie := b.won()
			if won {
				fmt.Println("Gagné !! board n°", i, "avec une suite:", serie)
				sum := b.unmarkedNumbersSum()
				fmt.Println("Résulat = ", bingoNumber*sum)
				finished = true
				break
			}
		}

		if finished {
			break
		}
	}

	//fmt.Println("after play :")
	//for _, b := range boards {
	//	fmt.Println(b)
	//}
}
