package day_2

import (
	"github.com/voziv/aoc-2022/internal/util"
)

func Run() {
	util.AssertResult("Part 1 example.txt", runPartOne("./internal/day_2/example.txt"), "15")
	util.AssertResult("Part 1 input.txt", runPartOne("./internal/day_2/input.txt"), "11873")
	util.AssertResult("Part 2 example.txt", runPartTwo("./internal/day_2/example.txt"), "12")
	util.AssertResult("Part 2 input.txt", runPartTwo("./internal/day_2/input.txt"), "12014")
	util.AssertResult("Part 2 Fast example.txt", runPartTwoCheeky("./internal/day_2/example.txt"), "12")
	util.AssertResult("Part 2 Fast input.txt", runPartTwoCheeky("./internal/day_2/input.txt"), "12014")

}
