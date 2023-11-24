package day_6

import (
	"github.com/voziv/aoc-2022/internal/util"
	"strconv"
	"strings"
)

func runPartOne(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	buffer := strings.Split(lines[0], "")

	result := useStupidNestedIf(buffer)

	return strconv.Itoa(result)
}

func useStupidNestedIf(buffer []string) int {
	var i int
	for i = 0; i < len(buffer)-4; i++ {

		if buffer[i] != buffer[i+1] && buffer[i] != buffer[i+2] && buffer[i] != buffer[i+3] {
			if buffer[i+1] != buffer[i+2] && buffer[i+1] != buffer[i+3] {
				if buffer[i+2] != buffer[i+3] {
					break
				}
			}
		}
	}

	return i + 4
}
