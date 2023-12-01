package day_6

import "github.com/voziv/aoc-2022/internal/util"

func Run() {
	util.AssertResult("Part 1 example.txt", runPartOne("./internal/day_6/example.txt"), "7")
	util.AssertResult("Part 1 example2.txt", runPartOne("./internal/day_6/example2.txt"), "5")
	util.AssertResult("Part 1 example3.txt", runPartOne("./internal/day_6/example3.txt"), "6")
	util.AssertResult("Part 1 example4.txt", runPartOne("./internal/day_6/example4.txt"), "10")
	util.AssertResult("Part 1 example5.txt", runPartOne("./internal/day_6/example5.txt"), "11")

	util.AssertResult("Part 1 input.txt", runPartOne("./internal/day_6/input.txt"), "1912")

	util.AssertResult("Part 2 example.txt", runPartTwo("./internal/day_6/example.txt"), "19")
	util.AssertResult("Part 2 example2.txt", runPartTwo("./internal/day_6/example2.txt"), "23")
	util.AssertResult("Part 2 example3.txt", runPartTwo("./internal/day_6/example3.txt"), "23")
	util.AssertResult("Part 2 example4.txt", runPartTwo("./internal/day_6/example4.txt"), "29")
	util.AssertResult("Part 2 example5.txt", runPartTwo("./internal/day_6/example5.txt"), "26")

	util.AssertResult("Part 2 input.txt", runPartTwo("./internal/day_6/input.txt"), "0")
}
