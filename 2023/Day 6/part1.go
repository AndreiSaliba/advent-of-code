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
	file, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	start := time.Now()
	lines := strings.Split(strings.TrimSuffix(string(file), "\n"), "\n")
	times := strings.Fields(lines[0][strings.Index(lines[0], ":")+1:])
	distances := strings.Fields(lines[1][strings.Index(lines[1], ":")+1:])

	var winningTimesRanges []int

	for i, distanceStr := range distances {
		distance, _ := strconv.Atoi(distanceStr)
		time, _ := strconv.Atoi(times[i])

		timeMin := int(math.Floor(float64(time / 2)))
		for distance < timeMin*(time-timeMin) {
			timeMin--
		}
		timeMax := int(math.Floor(float64(time/2))) + 1
		for distance < timeMax*(time-timeMax) {
			timeMax++
		}

		winningTimesRanges = append(winningTimesRanges, timeMax-timeMin-1)
	}

	result := winningTimesRanges[0]
	for i := 1; i < len(winningTimesRanges); i++ {
		result = result * winningTimesRanges[i]
	}

	fmt.Println("Done in", time.Since(start))
	fmt.Println(result)
}
