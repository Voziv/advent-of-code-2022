package main

import (
	"fmt"
	"github.com/voziv/aoc-2022/internal/rucksack"
)

const (
	BASE_DIR = "./docs/03-rucksack/"
)

func main() {

	runPartOne("example.txt", 157)

	runPartOne("input.txt", 7848)

}

func runPartOne(inputFileName string, expectedResult int) {
	fmt.Printf("\nRucksack Part One (%s)\n", inputFileName)
	result := rucksack.RunPartOne(BASE_DIR + inputFileName)

	fmt.Printf("Expecting: %d\n", expectedResult)
	fmt.Printf("Result: %d\n", result)

	if result != expectedResult {
		fmt.Println("Didn't get expected result")
	}
}
