package day1

import (
	"AdventOfCode2021/common"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Exo1() {
	//data, err := os.ReadFile("day1/input.txt")
	//check(err)
	//
	//fmt.Print(string(data))
	//
	//f, err := os.Open("/tmp/dat")
	//check(err)
	//

	file, err := os.Open("day1/input.txt")
	common.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var prev_n int64 = -1
	nb_inc := 0
	for scanner.Scan() {
		n, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		//fmt.Println(scanner.Text())
		if prev_n != -1 {
			if prev_n < n {
				fmt.Println(n, "-> increase")
				nb_inc++
			} else if prev_n > n {
				fmt.Println(n, "-> decrease")
			}
		}
		prev_n = n
	}

	fmt.Println("nb_inc =", nb_inc)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
