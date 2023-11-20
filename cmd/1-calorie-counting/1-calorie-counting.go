package main

import (
	"fmt"
	"github.com/voziv/aoc-2022/internal/calories"
)

func main() {
	run("./docs/01-calorie-counting/input1.txt")
	run("./docs/01-calorie-counting/input2.txt")
}

func run(inputFileName string) {
	fmt.Printf("\nCalorie Counting App: '%s'\n", inputFileName)
	calories.Count(inputFileName)
}
