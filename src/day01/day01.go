package day01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/erietz/aoc2022/src"
)

func Solve(input string) {
	caloriesPerElf := getCalsPerElf(input)
	caloriesPart1 := part1(caloriesPerElf)
	caloriesPart2 := part2(caloriesPerElf)

	fmt.Printf("Calories Part 1: %v\n", caloriesPart1)
	fmt.Printf("Calories Part 2: %v\n", caloriesPart2)
}

func part1(calsPerElf []int) int {
	_, max := aoc.MinMax(calsPerElf)
	return max
}

func getCalsPerElf(input string) []int {
	calsPerElf := []int{}
	for _, chunk := range strings.Split(strings.TrimSpace(input), "\n\n") {
		caloriesSum := 0
		for _, calString := range strings.Split(chunk, "\n") {
			calories, err := strconv.Atoi(calString)
			if err != nil {
				panic(err)
			}
			caloriesSum += calories
		}
		calsPerElf = append(calsPerElf, caloriesSum)
	}
	return calsPerElf
}

func part2(calsPerElf []int) int {
	calsCopy := append([]int{}, calsPerElf...)
	sort.Ints(calsCopy)

	sum := 0
	for _, v := range calsCopy[len(calsCopy)-3:] {
		sum += v
	}
	return sum
}
