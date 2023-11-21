package main

import (
	"fmt"
	"github.com/voziv/aoc-2022/internal/camp_cleanup"
)

const (
	BASE_DIR = "./docs/04-camp-cleanup/"
)

func main() {

	runPartOne("example.txt", 2)
	runPartOne("input.txt", 651)

	runPartTwo("example.txt", 4)
	runPartTwo("input.txt", 956)

}

func runPartOne(inputFileName string, expectedResult int) {
	fmt.Printf("\nCamp Cleanup Part One (%s)\n", inputFileName)
	result := camp_cleanup.RunPartOne(BASE_DIR + inputFileName)

	fmt.Printf("Expecting: %d\n", expectedResult)
	fmt.Printf("Result: %d\n", result)

	if result != expectedResult {
		fmt.Println("Didn't get expected result")
	}
}

func runPartTwo(inputFileName string, expectedResult int) {
	fmt.Printf("\nCamp Cleanup Part Two (%s)\n", inputFileName)
	result := camp_cleanup.RunPartTwo(BASE_DIR + inputFileName)

	fmt.Printf("Expecting: %d\n", expectedResult)
	fmt.Printf("Result: %d\n", result)

	if result != expectedResult {
		fmt.Println("Didn't get expected result")
	}
}
