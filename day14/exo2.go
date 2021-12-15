package day14

import (
	"AdventOfCode2021/common"
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"strings"
)

func Exo2(fileName string) {
	lines := common.ReadLinesFromFile(fileName)

	template := lines[0]
	rules := make(map[string]string)
	for i, line := range lines {
		if i >= 2 {
			var pair, toInsert string
			reader := strings.NewReader(line)
			fmt.Fscanf(reader, "%s -> %s", &pair, &toInsert)
			rules[pair] = toInsert
		}
	}

	pairs := make(map[string]int)
	for k, _ := range rules {
		pairs[k] = strings.Count(template, k)
	}

	fmt.Println(pairs)
	countsByLetter := make(map[string]int, 0)
	for _, v := range rules {
		countsByLetter[v] = strings.Count(template, v)
	}
	fmt.Println(countsByLetter)

	for step := 1; step <= 40; step++ {
		newPairs := make(map[string]int)

		for k, v := range rules {
			newPairs[string(k[0])+v] += pairs[k]
			newPairs[v+string(k[1])] += pairs[k]
			countsByLetter[v] += pairs[k]
		}

		fmt.Println("Step", step)
		fmt.Println(newPairs)

		fmt.Println(countsByLetter)
		max := From(countsByLetter).SelectT(func(kv KeyValue) int { return kv.Value.(int) }).Max().(int)
		min := From(countsByLetter).SelectT(func(kv KeyValue) int { return kv.Value.(int) }).Min().(int)
		fmt.Println(max - min)

		pairs = newPairs
	}

}
