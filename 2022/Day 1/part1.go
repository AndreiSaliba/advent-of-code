package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	var elfTotalCalories []int
	var currentElf []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			sum := 0
			for _, num := range currentElf {
				sum += num
			}

			elfTotalCalories = append(elfTotalCalories, sum)
			currentElf = currentElf[:0]
		}

		num, err := strconv.Atoi(line)
		if err == nil {
			currentElf = append(currentElf, num)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	max := 0
	for _, num := range elfTotalCalories {
		if max < num {
			max = num
		}
	}

	fmt.Println(max)
}
