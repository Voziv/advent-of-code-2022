package day_3

import (
	"github.com/voziv/aoc-2022/internal/util"
	"strconv"
	"strings"
)

func runPartOne(inputFileName string) string {
	rucksacks := util.GetFileContents(inputFileName)

	var prioritySum = 0

	for _, rucksack := range rucksacks {
		characters := strings.Split(rucksack, "")
		itemMap := map[string]int{}

		for i := 0; i < len(characters)/2; i++ {
			for j := len(characters) / 2; j < len(characters); j++ {
				if characters[i] == characters[j] {
					itemMap[characters[i]] = convertCharacterToPriority(characters[i][0])
				}
			}
		}

		for _, item := range itemMap {
			prioritySum += item
		}
	}

	return strconv.Itoa(prioritySum)
}

func convertCharacterToPriority(character uint8) int {
	if character >= "a"[0] {
		// a = 97
		return int(character) - int("a"[0]) + 1
	} else if character >= "A"[0] {
		// A = 65
		return int(character) - int("A"[0]) + 27
	}

	return 0
}
