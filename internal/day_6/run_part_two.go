package day_6

import (
	"github.com/voziv/aoc-2022/internal/util"
	"strconv"
	"strings"
)

func runPartTwo(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	buffer := strings.Split(lines[0], "")

	result := findPacketStart(buffer, 14)

	return strconv.Itoa(result)
}

func findPacketStart(buffer []string, lengthOfMarker int) int {
	var i int
	var foundPacket bool
	for i = 0; i < len(buffer)-4; i++ {

		foundPacket = true
		for j := 0; j < lengthOfMarker; j++ {
			for k := j + 1; k < lengthOfMarker; k++ {
				if buffer[i+j] == buffer[i+k] {
					foundPacket = false
					break
				}
			}
			if !foundPacket {
				break
			}
		}

		if foundPacket {
			break
		}
	}

	return i + lengthOfMarker
}
