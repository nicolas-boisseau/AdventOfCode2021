package day21

import (
	"AdventOfCode2021/common"
	"fmt"
	"log"
	"strings"
)

func Process(fileName string) int {
	lines := common.ReadLinesFromFile(fileName)

	reader := strings.NewReader(lines[0])
	var p1Position, p2Position int
	_, err := fmt.Fscanf(reader, "Player 1 starting position: %d", &p1Position)
	if err != nil {
		log.Fatal(err)
	}
	reader = strings.NewReader(lines[1])
	_, err = fmt.Fscanf(reader, "Player 2 starting position: %d", &p2Position)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Player 1 start at", p1Position)
	fmt.Println("Player 2 start at", p2Position)

	// Game loop
	d := newDeterministicDice(100)
	p1 := newPlayer(1, p1Position)
	p2 := newPlayer(2, p2Position)

	for !p1.HasWon(1000) && !p2.HasWon(1000) {

		// p1 plays
		p1.Play(d.Roll() + d.Roll() + d.Roll())
		fmt.Println("p1: pos=", p1.position, ", score=", p1.score)

		if p1.HasWon(1000) {
			break
		}

		p2.Play(d.Roll() + d.Roll() + d.Roll())
		fmt.Println("p2: pos=", p2.position, ", score=", p2.score)
	}

	fmt.Println("FINAL :")
	fmt.Println("p1: pos=", p1.position, ", score=", p1.score)
	fmt.Println("p2: pos=", p2.position, ", score=", p2.score)
	fmt.Println("dice: count=", d.rollCount)

	if p1.HasWon(1000) {
		return p2.score * d.rollCount
	}

	return p1.score * d.rollCount
}

func Process2(fileName string) int64 {
	lines := common.ReadLinesFromFile(fileName)

	reader := strings.NewReader(lines[0])
	var p1Position, p2Position int
	_, err := fmt.Fscanf(reader, "Player 1 starting position: %d", &p1Position)
	if err != nil {
		log.Fatal(err)
	}
	reader = strings.NewReader(lines[1])
	_, err = fmt.Fscanf(reader, "Player 2 starting position: %d", &p2Position)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Player 1 start at", p1Position)
	fmt.Println("Player 2 start at", p2Position)

	// Game loop
	p1 := newPlayer(1, p1Position)
	p2 := newPlayer(2, p2Position)
	universes := make([]*Universe, 0)
	u := &Universe{
		players:         make(map[int]*Player),
		multiple:        1,
		currentPlayerId: 1,
	}
	u.players[1] = p1
	u.players[2] = p2
	universes = append(universes, u)

	winners := make(map[int]int64)

	count := 0
	for len(universes) > 0 {
		u := universes[0]
		universes = universes[1:]

		nextUniverses := make([]*Universe, 0)

		//fmt.Println("Compute next universes...")
		nextUniverses = append(nextUniverses, u.Play(3, 1))
		nextUniverses = append(nextUniverses, u.Play(4, 3))
		nextUniverses = append(nextUniverses, u.Play(5, 6))
		nextUniverses = append(nextUniverses, u.Play(6, 7))
		nextUniverses = append(nextUniverses, u.Play(7, 6))
		nextUniverses = append(nextUniverses, u.Play(8, 3))
		nextUniverses = append(nextUniverses, u.Play(9, 1))

		//fmt.Println("Next universes computed.")

		//fmt.Println("Now searching winners.")
		for _, nextU := range nextUniverses {
			if nextU.winnerId != 0 {
				winners[nextU.winnerId] += nextU.multiple
			} else {
				universes = append(universes, nextU)
			}
		}

		count++

		//fmt.Println("universes:", len(universes))
		if count%100000 == 0 {
			fmt.Println("winners:", winners)
			fmt.Println("universes:", len(universes))
		}
	}

	fmt.Println("FINAL :")
	fmt.Println("winners:", winners)

	if winners[1] > winners[2] {
		return winners[1]
	}

	return winners[2]
}
