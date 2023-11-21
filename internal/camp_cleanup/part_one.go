package camp_cleanup

import (
	"github.com/voziv/aoc-2022/internal/util"
	"strconv"
	"strings"
)

func RunPartOne(inputFileName string) int {
	result := 0

	lines := util.GetFileContents(inputFileName)

	for _, line := range lines {
		ranges := strings.Split(line, ",")

		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")

		range1Low := castToInt(range1[0])
		range1High := castToInt(range1[1])
		range2Low := castToInt(range2[0])
		range2High := castToInt(range2[1])

		if range1Low >= range2Low && range1High <= range2High {
			result++
		} else if range2Low >= range1Low && range2High <= range1High {
			result++
		}
	}

	return result
}

func castToInt(digits string) int {
	number, err := strconv.Atoi(digits)
	if err != nil {
		panic(err)
	}

	return number
}
