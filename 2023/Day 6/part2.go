package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	lines := strings.Split(strings.TrimSuffix(string(file), "\n"), "\n")
	timeStr := strings.ReplaceAll(lines[0][strings.Index(lines[0], ":")+1:], " ", "")
	distanceStr := strings.ReplaceAll(lines[1][strings.Index(lines[1], ":")+1:], " ", "")
	start := time.Now()

	var winningTimesCounts []int

	timeInt, _ := strconv.Atoi(strings.TrimSpace(timeStr))
	distance, _ := strconv.Atoi(strings.TrimSpace(distanceStr))

	timeMin := int(math.Floor(float64(timeInt / 2)))
	for distance < timeMin*(timeInt-timeMin) {
		timeMin--
	}
	timeMax := int(math.Floor(float64(timeInt/2))) + 1
	for distance < timeMax*(timeInt-timeMax) {
		timeMax++
	}

	winningTimesCounts = append(winningTimesCounts, timeMax-timeMin-1)

	result := winningTimesCounts[0]
	for i := 1; i < len(winningTimesCounts); i++ {
		result = result * winningTimesCounts[i]
	}

	fmt.Println("Done in", time.Since(start))
	fmt.Println(result)
}
