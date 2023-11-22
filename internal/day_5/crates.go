package day_5

import "fmt"

func GetTopCrates(crates [][]string) string {
	result := ""
	for i := 0; i < len(crates); i++ {
		result += crates[i][len(crates[i])-1]
	}
	return result
}

func PrintCrates(crates [][]string) {
	fmt.Println("Crates")
	highest := 0
	for i := 0; i < len(crates); i++ {
		if len(crates[i]) > highest {
			highest = len(crates[i])
		}
	}

	for j := highest; j >= 0; j-- {
		for i := 0; i < len(crates); i++ {
			if len(crates[i]) > j {
				fmt.Printf("(%s) ", crates[i][j])
			} else {
				fmt.Printf("    ")
			}
		}
		fmt.Println()
	}

	for i := 0; i < len(crates); i++ {
		fmt.Printf(" %-2d ", i)
	}
	fmt.Println()
	fmt.Println("End of Crates")
}
