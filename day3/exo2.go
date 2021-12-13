package day3

import (
	"AdventOfCode2021/common"
	"fmt"
	"strconv"
)

func filter(lines []string, index int, ones []int, zeros []int, eval func(string, int, []int, []int) bool) []string {
	output := make([]string, 0)
	for _, line := range lines {
		if eval(line, index, ones, zeros) {
			output = append(output, line)
		}
	}
	return output
}

func oxygen_filter(line string, index int, ones []int, zeros []int) bool {
	var mostCommon uint8 = '1'
	if zeros[index] > ones[index] {
		mostCommon = '0'
	}
	return line[index] == mostCommon
}

func co2_scrubber_filter(line string, index int, ones []int, zeros []int) bool {
	var leastCommon uint8 = '0'
	if ones[index] < zeros[index] {
		leastCommon = '1'
	}
	return line[index] == leastCommon
}

func computeOnesAndZeros(lines []string) ([]int, []int) {
	ones := make([]int, len(lines[0]))
	zeros := make([]int, len(lines[0]))
	for _, line := range lines {
		for j, c := range line {
			if c == '1' {
				ones[j]++
			} else {
				zeros[j]++
			}
		}
	}

	return ones, zeros
}

func Exo2() {
	lines := common.ReadLinesFromFile("day3/input.txt")

	ones, zeros := computeOnesAndZeros(lines)

	index := 0
	oxygen_lines := make([]string, len(lines))
	copy(oxygen_lines, lines)

	for len(oxygen_lines) > 1 {
		oxygen_lines = filter(oxygen_lines, index, ones, zeros, oxygen_filter)
		fmt.Println("===================")
		fmt.Println(oxygen_lines)
		ones, zeros = computeOnesAndZeros(oxygen_lines)
		index++
	}
	fmt.Println("debug :", oxygen_lines, "index=", index)

	oxygen_rating, _ := strconv.ParseInt(oxygen_lines[0], 2, 0)
	fmt.Println("oxygen generator rating =", oxygen_rating)

	index = 0
	co2_scrubber_lines := make([]string, len(lines))
	copy(co2_scrubber_lines, lines)

	for len(co2_scrubber_lines) > 1 {
		co2_scrubber_lines = filter(co2_scrubber_lines, index, ones, zeros, co2_scrubber_filter)
		ones, zeros = computeOnesAndZeros(co2_scrubber_lines)
		index++
	}
	fmt.Println("debug :", co2_scrubber_lines, ", index=", index)

	co2_scrubber_rating, _ := strconv.ParseInt(co2_scrubber_lines[0], 2, 0)
	fmt.Println("co2_scrubber rating =", co2_scrubber_rating)

	fmt.Println("Life support rating =", oxygen_rating*co2_scrubber_rating)

	//index = 0
	//co2_scrubber_lines := make([]string, len(lines))
	//copy(co2_scrubber_lines, lines)
	//for len(co2_scrubber_lines) > 1 {
	//	co2_scrubber_lines = filter(co2_scrubber_lines, index, ones, zeros, co2_scrubber_filter)
	//	index++
	//}
	//
	//fmt.Println("CO2 scrubber rating =", lines[0])

	//gamma := ""
	//epsilon := ""
	//for i, _ := range ones {
	//	if ones[i] > zeros[i] {
	//		gamma += "1"
	//		epsilon += "0"
	//	} else {
	//		gamma += "0"
	//		epsilon += "1"
	//	}
	//}
	//
	//gammaNumber, _ := strconv.ParseInt(gamma, 2, 0)
	//epsilonNumber, _ := strconv.ParseInt(epsilon, 2, 0)
	//fmt.Println(gammaNumber)
	//fmt.Println(epsilonNumber)
	//fmt.Println("Power consumption:", gammaNumber * epsilonNumber)
}
