package day1

import (
	"AdventOfCode2021/common"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Exo2() {
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
	//var prev_n int64 = -1
	lines := make([]int64, 0)
	nb_inc := 0
	for scanner.Scan() {
		n, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		lines = append(lines, n)
		//fmt.Println(scanner.Text())
		//if prev_n != -1 {
		//	if prev_n < n {
		//		fmt.Println(n, "-> increase")
		//		nb_inc++
		//	} else if prev_n > n {
		//		fmt.Println(n, "-> decrease")
		//	}
		//}
		//prev_n = n
	}

	var sum int64 = 0
	var prev_sum int64 = -1
	for i := 0; i < len(lines); i++ {
		if i < len(lines)-2 {
			sum = lines[i] + lines[i+1] + lines[i+2]
		}

		if prev_sum != -1 {
			if prev_sum < sum {
				fmt.Println(sum, "-> increase")
				nb_inc++
			}
		}
		prev_sum = sum
	}

	fmt.Println("nb_inc =", nb_inc)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
