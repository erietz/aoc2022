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
	return calsPerElf
}

func part2(calsPerElf []int) int {
	tmp := append([]int{}, calsPerElf...)
	sort.Ints(tmp)

	sum := 0
	for _, v := range tmp[len(tmp)-3:] {
		sum += v
	}
	return sum
}
