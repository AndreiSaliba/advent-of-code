package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"math"
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	strNumbers := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var calibrationValues []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var calibrationNum []string

		strNumbersFirstIndex := math.MaxInt
		strNumbersFirstNum := ""
		for i, numString := range strNumbers {
			index := strings.Index(line, numString)
			if index < strNumbersFirstIndex && index != -1 {
				strNumbersFirstIndex = index
				strNumbersFirstNum = strconv.Itoa(i + 1)
			}
		}

		strNumbersLastIndex := -1
		strNumbersLastNum := ""
		for i, numString := range strNumbers {
			index := strings.LastIndex(line, numString)
			if index > strNumbersLastIndex && index != -1{
				strNumbersLastIndex = index
				strNumbersLastNum = strconv.Itoa(i + 1)
			}
		}

		numbersFirstIndex := -1
		numbersFirstNum := ""
		for i := 0; i < len(line); i++ {
			if unicode.IsNumber(rune(line[i])) {
				numbersFirstIndex = i
				numbersFirstNum = string(line[i])
				break
			}
		}
		
		numbersLastIndex := -1
		numbersLastNum := ""
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsNumber(rune(line[i])) {
				numbersLastIndex = i
				numbersLastNum = string(line[i])
				break
			}
		}

		if numbersFirstIndex != -1 && strNumbersFirstIndex != - 1 {
			if numbersFirstIndex < strNumbersFirstIndex {
				calibrationNum = append(calibrationNum, numbersFirstNum)
			} else {
				calibrationNum = append(calibrationNum, strNumbersFirstNum)
			}
		} else if numbersFirstIndex == -1 {
			calibrationNum = append(calibrationNum, strNumbersFirstNum)
		} else {
			calibrationNum = append(calibrationNum, numbersFirstNum)
		}

		if numbersLastIndex != -1 && strNumbersLastIndex != - 1 {
			if numbersLastIndex > strNumbersLastIndex {
				calibrationNum = append(calibrationNum, numbersLastNum)
			} else {
				calibrationNum = append(calibrationNum, strNumbersLastNum)
			}
		} else if numbersLastIndex == -1 {
			calibrationNum = append(calibrationNum, strNumbersLastNum)
		} else {
			calibrationNum = append(calibrationNum, numbersLastNum)
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
