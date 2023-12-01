package day_3

import "github.com/voziv/aoc-2022/internal/util"

func Run() {
	util.AssertResult("Part 1 example.txt", runPartOne("./internal/day_3/example.txt"), "157")
	util.AssertResult("Part 1 input.txt", runPartOne("./internal/day_3/input.txt"), "7848")
	util.AssertResult("Part 2 example.txt", runPartTwo("./internal/day_3/example.txt"), "70")
	util.AssertResult("Part 2 input.txt", runPartTwo("./internal/day_3/input.txt"), "2616")
}
