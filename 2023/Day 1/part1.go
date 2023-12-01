package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	var calibrationValues []int

	for scanner.Scan() {
		line := scanner.Text()
		var calibrationNum []string

		for i := 0; i < len(line); i++ {
			if unicode.IsNumber(rune(line[i])) {
				calibrationNum = append(calibrationNum, string(line[i]))
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsNumber(rune(line[i])) {
				calibrationNum = append(calibrationNum, string(line[i]))
				break
			}
		}

		num, err := strconv.Atoi(strings.Join(calibrationNum[:], ""))
		if err == nil {
			calibrationValues = append(calibrationValues, num)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	total := 0
	for _, num := range calibrationValues {
		total = total + num
	}

	fmt.Println("Total: ", total)
}
