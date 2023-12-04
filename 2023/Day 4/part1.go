package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	lines := strings.Split(strings.ReplaceAll(string(file), ".", " "), "\n")
	total := 0

	for _, line := range lines {
		if line != "" {
			winningNumbers := strings.Split(line[strings.Index(line, ":")+2:strings.Index(line, "|")-1], " ")
			myNumbers := strings.Split(line[strings.Index(line, "|")+2:], " ")
			var myWinningNumbers []int

			for _, number := range myNumbers {
				intNumber, err := strconv.Atoi(strings.TrimSpace(number))

				if err == nil && Contains(winningNumbers, intNumber) {
					myWinningNumbers = append(myWinningNumbers, intNumber)
				}
			}

			if len(myWinningNumbers) > 0 {
				total += int(math.Floor(math.Pow(2, float64(len(myWinningNumbers)-1))))
			}
		}
	}

	fmt.Println("Total:", total)
}

func Contains(array []string, number int) bool {
	for _, item := range array {
		intItem, err := strconv.Atoi(strings.TrimSpace(item))

		if err == nil && intItem == number {
			return true
		}
	}

	return false
}
