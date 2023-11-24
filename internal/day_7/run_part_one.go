package day_7

import (
	"fmt"
	"github.com/voziv/aoc-2022/internal/util"
	"strconv"
	"strings"
)

// $ = command line
// cd to change dir
// ls to list the contents of the current directory
// dir <name> there's a folder
// 1234 <name> filesize followed by file name

func runPartOne(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	var directorySizes map[string]int = map[string]int{}
	var processedFiles map[string]int = map[string]int{}
	var directory []string

	for _, line := range lines {
		tokens := strings.Split(line, " ")

		if tokens[0] == "$" {
			cmd := tokens[1]

			if cmd == "cd" {
				dir := tokens[2]
				fmt.Printf("COMMAND: %s %s\n", cmd, dir)

				if dir == ".." {
					directory = directory[:len(directory)-1]
				} else if dir == "/" {
					directory = []string{}
				} else {
					directory = append(directory, dir)
				}

				fmt.Printf("PWD: %s\n", getDirectory(directory))
			} else if cmd == "ls" {
				// I don't think we care about this tbh
			}
		} else if line[:3] == "dir" {
			// Don't think we care about directories until we CD into them
		} else if size, err := strconv.Atoi(tokens[0]); err == nil {
			cwd := getDirectory(directory)
			fileName := cwd + "/" + tokens[1]

			if _, alreadyProcessed := processedFiles[fileName]; alreadyProcessed == true {
				continue
			}

			for i := 0; i < len(directory); i++ {
				cwd := getDirectory(directory[:len(directory)-i])
				_, keyExists := directorySizes[cwd]
				if !keyExists {
					directorySizes[cwd] = 0
				}

				directorySizes[cwd] += size
			}

			processedFiles[fileName] = size
		}

	}

	total := 0
	for dir, size := range directorySizes {
		fmt.Println(dir, size)
		if size <= 100000 {
			total += size
		}
	}

	return strconv.Itoa(total)
}

func getDirectory(directory []string) string {
	return "/" + strings.Join(directory, "/")
}
