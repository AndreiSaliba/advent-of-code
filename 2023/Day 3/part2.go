package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	lines := strings.Split(strings.ReplaceAll(string(file), ".", " "), "\n")
	total := 0

	for i, line := range lines {
		for j, char := range line {
			if char == '*' {
				var surroundings []string

				lPad := j - 1
				if j <= 0 {
					lPad = j
				}
				rPad := j + 2
				if j >= len(line)-1 {
					rPad = j + 1
				}

				if i > 0 {
					surroundings = append((surroundings), lines[i-1][lPad:rPad])
				}

				surroundings = append((surroundings), lines[i][lPad:rPad])

				if i < len(lines)-1 && len(lines[i]) == len(lines[i+1]) {
					surroundings = append((surroundings), lines[i+1][lPad:rPad])
				}

				var surroundingPartNums []int
				for k, level := range surroundings {
					for l, char := range level {
						if unicode.IsNumber(char) {
							line := i + (k - 1)
							lIndex, rIndex := j+(l-1), j+(l-1)

							for lIndex >= 0 && unicode.IsNumber(rune(lines[line][lIndex])) {
								lIndex -= 1
							}
							for rIndex < len(lines[line]) && unicode.IsNumber(rune(lines[line][rIndex])) {
								rIndex += 1
							}

							num, _ := strconv.Atoi(lines[line][lIndex+1 : rIndex])

							containsNum := false
							for _, partNum := range surroundingPartNums {
								if partNum == num {
									containsNum = true
								}
							}

							if !containsNum {
								surroundingPartNums = append(surroundingPartNums, num)
							}
						}
					}
				}

				if len(surroundingPartNums) == 2 {
					total += surroundingPartNums[0] * surroundingPartNums[1]
				}
			}
		}
	}

	fmt.Println("Total", total)
}
