package day_7

import "github.com/voziv/aoc-2022/internal/util"

func Run() {
	util.AssertResult("Part 1 example.txt", runPartOne("./inputs/7/example.txt"), "95437")
	util.AssertResult("Part 1 input.txt", runPartOne("./inputs/7/input.txt"), "1367870")
	//
	//util.AssertResult("Part 2 example.txt", runPartTwo("./inputs/7/example.txt"), "0")
	//util.AssertResult("Part 2 input.txt", runPartTwo("./inputs/7/input.txt"), "0")
}
