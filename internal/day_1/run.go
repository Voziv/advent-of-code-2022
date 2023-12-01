package day_1

import (
	"github.com/voziv/aoc-2022/internal/util"
	"sort"
	"strconv"
)

func Run() {
	util.AssertResult("Part 1 example.txt", runPartOne("./internal/day_1/example.txt"), "24000")
	util.AssertResult("Part 1 input.txt", runPartOne("./internal/day_1/input.txt"), "65912")
	util.AssertResult("Part 2 example.txt", runPartTwo("./internal/day_1/example.txt"), "45000")
	util.AssertResult("Part 2 input.txt", runPartTwo("./internal/day_1/input.txt"), "195625")
}

func runPartOne(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	elfCalorieCounts := CountCalories(lines)

	// Sort
	sort.Slice(elfCalorieCounts, func(i, j int) bool {
		return elfCalorieCounts[i] > elfCalorieCounts[j]
	})

	// Single highest
	return strconv.Itoa(elfCalorieCounts[0])
}

func runPartTwo(inputFileName string) string {

	lines := util.GetFileContents(inputFileName)

	elfCalorieCounts := CountCalories(lines)

	// Sort
	sort.Slice(elfCalorieCounts, func(i, j int) bool {
		return elfCalorieCounts[i] > elfCalorieCounts[j]
	})

	// Top three summed
	var totalCalories = 0
	for i := 0; i < 3; i++ {
		totalCalories += elfCalorieCounts[i]
	}

	return strconv.Itoa(totalCalories)
}

func CountCalories(lines []string) []int {

	// All the problem wants is the highest total amount, as well as the top three totals summed.
	// We don't actually need to know which elf has what.
	var elfCalorieCounts []int

	var currentTotal = 0

	for _, line := range lines {
		// If we encounter a blank line it means that we're done with the calorie counts for the current elf.
		if line == "" {
			elfCalorieCounts = append(elfCalorieCounts, currentTotal)
			currentTotal = 0
			continue
		}

		amount, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		currentTotal += amount
	}

	// If the file doesn't have a blank line at the end the final elf won't be accounted for.
	if currentTotal > 0 {
		elfCalorieCounts = append(elfCalorieCounts, currentTotal)
	}

	return elfCalorieCounts
}
