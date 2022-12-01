package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	caloriesPart1 := part1(string(bytes))

	fmt.Printf("Calories Part 1: %v\n", caloriesPart1)
}

func part1(input string) int {
	data := strings.Split(input, "\n")
	calsPerElf := []int{}

	cals := 0
	for _, d := range data {
		if d == "" {
			calsPerElf = append(calsPerElf, cals)
			cals = 0
			continue
		}

		numCals, err := strconv.Atoi(d)
		if err != nil {
			panic(err)
		}
		cals += numCals
	}
	_, max := minMax(calsPerElf)
	return max
}

func minMax(array []int) (int, int) {
	min := array[0]
	max := array[0]
	for _, v := range array {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}
