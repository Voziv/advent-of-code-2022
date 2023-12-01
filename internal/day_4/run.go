package day_4

import "github.com/voziv/aoc-2022/internal/util"

func Run() {
	util.AssertResult("Part 1 example.txt", runPartOne("./internal/day_4/example.txt"), "2")
	util.AssertResult("Part 1 input.txt", runPartOne("./internal/day_4/input.txt"), "651")
	util.AssertResult("Part 2 example.txt", runPartTwo("./internal/day_4/example.txt"), "4")
	util.AssertResult("Part 2 input.txt", runPartTwo("./internal/day_4/input.txt"), "956")
}
