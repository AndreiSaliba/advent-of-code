package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Seed struct {
	Start int
	End   int
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	sections := strings.Split(strings.ReplaceAll(strings.TrimSuffix(string(file), "\n"), "\n\r", "\n\n"), "\n\n")
	strSeeds := strings.Split(sections[0][strings.Index(sections[0], ":")+2:], " ")
	start := time.Now()

	var seeds []Seed
	for i := 0; i < len(strSeeds); i += 2 {
		min, _ := strconv.Atoi(strings.TrimSpace(strSeeds[i]))
		max, _ := strconv.Atoi(strings.TrimSpace(strSeeds[i+1]))

		seeds = append(seeds, Seed{min, min + max})
	}

	for _, section := range sections[1:] {
		ranges := strings.Split(section, "\n")[1:]
		var newSeeds []Seed

		for len(seeds) > 0 {
			seed := seeds[len(seeds)-1:][0]
			seeds = seeds[:len(seeds)-1]

			check := true
			for _, rangeStr := range ranges[1:] {
				var dest, src, r int
				fmt.Sscanf(strings.TrimSpace(rangeStr), "%d %d %d", &dest, &src, &r)

				overlapStart := int(math.Max(float64(seed.Start), float64(src)))
				overlapEnd := int(math.Min(float64(seed.End), float64(src+r)))

				if overlapStart < overlapEnd {
					newSeeds = append(newSeeds, Seed{overlapStart - src + dest, overlapEnd - src + dest})
					if overlapStart > seed.Start {
						seeds = append(seeds, Seed{seed.Start, overlapStart})
					}
					if seed.End > overlapEnd {
						seeds = append(seeds, Seed{overlapEnd, seed.End})
					}
					check = false
					break
				}
			}
			if check {
				newSeeds = append(newSeeds, Seed{seed.Start, seed.End})
			}
		}
		seeds = newSeeds
	}

	min := seeds[0].Start
	for _, s := range seeds {
		if min > s.Start {
			min = s.Start
		}
	}
	fmt.Println("Done in", time.Since(start))
	fmt.Println(min)
}
