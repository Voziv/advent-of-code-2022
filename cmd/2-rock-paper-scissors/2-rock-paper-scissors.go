package main

import (
	"fmt"
	"github.com/voziv/aoc-2022/internal/rock_paper_scissors"
)

func main() {
	run("./docs/02-rock-paper-scissors/example.txt")
	run("./docs/02-rock-paper-scissors/input.txt")
}

func run(inputFileName string) {
	fmt.Printf("\nRock Paper Scissors: '%s'\n", inputFileName)
	rock_paper_scissors.Play(inputFileName)
}
