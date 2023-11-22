package day_5

import (
	"fmt"
	"github.com/voziv/aoc-2022/internal/util"
)

func runPartTwo(inputFileName string) string {
	result := ""

	lines := util.GetFileContents(inputFileName)

	for _, line := range lines {
		fmt.Println(line)
	}

	return result
}
