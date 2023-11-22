package day_4

import "github.com/voziv/aoc-2022/internal/util"

func Run() {
	util.AssertResult("Part 1 example.txt", runPartOne("./inputs/4/example.txt"), "2")
	util.AssertResult("Part 1 input.txt", runPartOne("./inputs/4/input.txt"), "651")
	util.AssertResult("Part 2 example.txt", runPartTwo("./inputs/4/example.txt"), "4")
	util.AssertResult("Part 2 input.txt", runPartTwo("./inputs/4/input.txt"), "956")
}
