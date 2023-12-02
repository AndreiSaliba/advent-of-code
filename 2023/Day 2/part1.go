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

	var possibleGames []int
	colors := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), ";", ""), ",", "")
		gameId, _ := strconv.Atoi(line[strings.Index(line, " ")+1 : strings.Index(line, ":")])
		possible := true

		for color, _ := range colors {
			splitLine := strings.Split(line, " ")
			for i, str := range splitLine {
				if strings.Contains(str, color) {
					if num, _ := strconv.Atoi(splitLine[i-1]); num > colors[str] {
						possible = false
						break
					}
				}
			}
		}

		if possible {
			possibleGames = append(possibleGames, gameId)
		}
	}

	total := 0
	for _, num := range possibleGames {
		total += num
	}

	fmt.Println("Total:", total)
}
