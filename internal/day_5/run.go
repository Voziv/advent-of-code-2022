package day_5

import "github.com/voziv/aoc-2022/internal/util"

func Run() {
	util.AssertResult("Part 1 example.txt", runPartOne("./inputs/5/example.txt"), "CMZ")
	util.AssertResult("Part 1 input.txt", runPartOne("./inputs/5/input.txt"), "BZLVHBWQF")
	util.AssertResult("Part 2 example.txt", runPartTwo("./inputs/5/example.txt"), "MCD")
	util.AssertResult("Part 2 input.txt", runPartTwo("./inputs/5/input.txt"), "TDGJQTZSL")
}
