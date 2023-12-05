package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	sections := strings.Split(strings.ReplaceAll(string(file), "\n\r", "\n\n"), "\n\n")
	seeds := strings.Split(sections[0][strings.Index(sections[0], ":")+2:], " ")
	var seedMappings []int

	for _, i := range seeds {
		seedInt, err := strconv.Atoi(strings.TrimSpace(i))
		if err == nil {
			seedMappings = append(seedMappings, seedInt)
		}
	}

	for _, section := range sections[1:] {
		for i, seed := range seedMappings {
			if i == 0 {
				seedMappings = seedMappings[:0]
			}

			seedMappings = append(seedMappings, GetMapResult(section, seed))
		}
		fmt.Println(seedMappings)
	}

	fmt.Println("Closest Location:", slices.Min(seedMappings))
}

func GetMapResult(almanacMap string, input int) int {
	ranges := strings.Split(almanacMap, "\n")[2:]

	for _, rangeString := range ranges {
		if rangeString != "" {
			splitRangeString := strings.Split(rangeString, " ")

			destination, _ := strconv.Atoi(strings.TrimSpace(splitRangeString[0]))
			source, _ := strconv.Atoi(strings.TrimSpace(splitRangeString[1]))
			r, _ := strconv.Atoi(strings.TrimSpace(splitRangeString[2]))

			if input >= source && input <= source+r {
				return destination + (input - source)
			}
		}
	}

	return input
}
