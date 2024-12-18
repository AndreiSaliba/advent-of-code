package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	powersTotal := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		colors := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		line := strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), ";", ""), ",", "")

		for color, _ := range colors {
			splitLine := strings.Split(line, " ")
			for i, str := range splitLine {
				if strings.Contains(str, color) {
					if num, _ := strconv.Atoi(splitLine[i-1]); num > colors[str] {
						colors[str] = num
					}
				}
			}
		}

		powersTotal += colors["red"] * colors["green"] * colors["blue"]
	}

	fmt.Println("Total:", powersTotal)
}
