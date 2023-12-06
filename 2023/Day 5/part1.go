package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	sections := strings.Split(strings.ReplaceAll(strings.TrimSuffix(string(file), "\n"), "\n\r", "\n\n"), "\n\n")
	strSeeds := strings.Split(sections[0][strings.Index(sections[0], ":")+2:], " ")

	var seeds []int
	for _, i := range strSeeds {
		seedInt, err := strconv.Atoi(strings.TrimSpace(i))
		if err == nil {
			seeds = append(seeds, seedInt)
		}
	}

	for _, section := range sections[1:] {
		for i, seed := range seeds {
			if i == 0 {
				seeds = seeds[:0]
			}

			seeds = append(seeds, GetMapResult(section, seed))
		}
	}

	fmt.Println(slices.Min(seeds))
}

func GetMapResult(almanacMap string, input int) int {
	ranges := strings.Split(almanacMap, "\n")[2:]

	for _, rangeStr := range ranges {
		var dest, src, r int
		fmt.Sscanf(strings.TrimSpace(rangeStr), "%d %d %d", &dest, &src, &r)

		if input >= src && input <= src+r {
			return dest + (input - src)
		}
	}

	return input
}
