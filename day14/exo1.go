package day14

import (
	"AdventOfCode2021/common"
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"strings"
)

func matchRule(template, pair string, toInsert string) ([]int, string) {
	//m, _ := regexp.MatchString(pair, template)
	//re := regexp.MustCompile(pair)
	//indexes := re.FindAllStringIndex(template, len(template))

	indexes2 := make([]int, 0)
	str := template
	ind := strings.Index(str, pair)
	if ind != -1 {
		indexes2 = append(indexes2, ind+1)
		delta := ind + 1
		for ind != -1 {
			str = template[delta:]

			ind = strings.Index(str, pair)
			if ind != -1 {
				indexes2 = append(indexes2, delta+ind+1)
				delta += ind + 1
			}
		}
	}
	//fmt.Println(indexes2)

	//fmt.Println(indexes)
	//if indexes == nil {
	//	indexes = make([][]int, 0)
	//}

	//re.FindAllIndex(template)

	return indexes2, toInsert
}

type Item struct {
	index    int
	toInsert string
}

func Exo1(fileName string) {
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

	fmt.Println("Template: ", template)
	fmt.Println("rules: ", rules)

	newTemplate := template
	for step := 1; step <= 40; step++ {
		work := make([]Item, 0)
		for k, v := range rules {
			indexes, toInsert := matchRule(newTemplate, k, v)
			for _, ind := range indexes {
				work = append(work, Item{index: ind, toInsert: toInsert})
			}
		}

		sortedWork := make([]Item, len(work))
		From(work).OrderByDescendingT(func(item Item) int { return item.index }).ToSlice(&sortedWork)

		//fmt.Println("work:", sortedWork)

		// Perform replaces
		replaced := newTemplate
		for _, workItem := range sortedWork {
			replaced = replaced[:workItem.index] + workItem.toInsert + replaced[workItem.index:]

		}

		fmt.Println("step", step, ":", replaced)

		fmt.Println("step", step)

		//if step == 40 {
		countsByLetter := make(map[string]int, 0)
		for _, v := range rules {
			_, isPresent := countsByLetter[v]
			if !isPresent {
				countsByLetter[v] = strings.Count(replaced, v)
			}
		}

		fmt.Println(countsByLetter)
		max := From(countsByLetter).SelectT(func(kv KeyValue) int { return kv.Value.(int) }).Max().(int)
		min := From(countsByLetter).SelectT(func(kv KeyValue) int { return kv.Value.(int) }).Min().(int)
		fmt.Println(max - min)
		//}

		newTemplate = replaced
	}
}

const colorReset = "\033[0m"
const colorGreen = "\033[32m"
