package day_3

import (
	"github.com/voziv/aoc-2022/internal/util"
	"strconv"
	"strings"
)

func runPartTwo(inputFileName string) string {
	rucksacks := util.GetFileContents(inputFileName)

	var prioritySum = 0

	for i := 0; i < len(rucksacks); i += 3 {
		items := map[string]int{}
		items = getUniqueItems(rucksacks[i], items)
		items = getUniqueItems(rucksacks[i+1], items)
		items = getUniqueItems(rucksacks[i+2], items)

		for item, count := range items {
			if count == 3 {
				prioritySum += convertCharacterToPriority(item[0])
				break
			}
		}
	}

	return strconv.Itoa(prioritySum)
}

func getUniqueItems(rucksack string, items map[string]int) map[string]int {
	itemMap := map[string]int{}

	for _, item := range strings.Split(rucksack, "") {
		itemMap[item] = 1
	}

	for i, _ := range itemMap {
		_, ok := items[i]
		if ok {
			items[i]++
		} else {
			items[i] = 1
		}
	}

	return items
}
