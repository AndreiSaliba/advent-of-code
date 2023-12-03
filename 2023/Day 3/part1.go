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
	currentNumIsPart := false
	total := 0

	for i, line := range lines {
		for j, char := range line {
			if unicode.IsNumber(rune(char)) {
				surroundings := ""

				lPad := j - 1
				if j <= 0 {
					lPad = j
				}
				rPad := j + 2
				if j >= len(line)-1 {
					rPad = j + 1
				}

				if i > 0 {
					surroundings += lines[i-1][lPad:rPad]
				}

				surroundings += lines[i][lPad:rPad]

				if i < len(lines)-1 && len(lines[i]) == len(lines[i+1]) {
					surroundings += lines[i+1][lPad:rPad]
				}

				for _, char := range surroundings {
					if char != ' ' && !unicode.IsNumber(char) {
						lIndex, rIndex := j, j
						for lIndex >= 0 && unicode.IsNumber(rune(line[lIndex])) {
							lIndex -= 1
						}
						for rIndex < len(line) && unicode.IsNumber(rune(line[rIndex])) {
							rIndex += 1
						}

						num, _ := strconv.Atoi(line[lIndex+1 : rIndex])
						if currentNumIsPart == false {
							total += num
						}

						currentNumIsPart = true
						break
					}
				}
			} else {
				currentNumIsPart = false
			}
		}
	}

	fmt.Println("Total", total)
}
