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

	start := time.Now()
	lines := strings.Split(strings.TrimSuffix(string(file), "\n"), "\n")
	timeStr := strings.ReplaceAll(lines[0][strings.Index(lines[0], ":")+1:], " ", "")
	distanceStr := strings.ReplaceAll(lines[1][strings.Index(lines[1], ":")+1:], " ", "")

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

	fmt.Println("Done in", time.Since(start))
	fmt.Println(timeMax - timeMin - 1)
}
