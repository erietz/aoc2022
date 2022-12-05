package day03

import (
	"errors"
	"fmt"
	"strings"
)

func Solve(input string) {
	sum1 := part1(input)
	// Part 1 sum: 7737
	fmt.Println("Part 1 sum:", sum1)

	// Part 2 sum: 2697
	sum2 := part2(input)
	fmt.Println("Part 2 sum:", sum2)
}

func part1(input string) int {
	sum := 0
	for _, items := range strings.Split(strings.TrimSpace(input), "\n") {
		comp1, comp2 := buildCompartmentCounts(items)
		for item := range comp1 {
			if _, ok := comp2[item]; ok {
				char := []rune(item)[0]
				priority := rune2Priority(char)
				sum += priority
			}
		}
	}
	return sum
}

func part2(input string) int {
	sum := 0
	rucksacks := make([]string, 0)
	// NOTE: Not trimming off the last newline for one extra iteration here
	for _, items := range strings.Split(input, "\n") {
		if len(rucksacks) == 3 {
			commonItem, err := getCommonItem(rucksacks)
			if err != nil {
				panic(err)
			}
			sum += rune2Priority(commonItem)
			rucksacks = make([]string, 0)
		}
		rucksacks = append(rucksacks, items)
	}
	return sum
}

func getCommonItem(rucksacks []string) (rune, error) {
	sack0Counts := buildCounts(rucksacks[0])
	sack1Counts := buildCounts(rucksacks[1])
	sack2Counts := buildCounts(rucksacks[2])

	for item := range sack0Counts {
		if _, ok1 := sack1Counts[item]; ok1 {
			if _, ok2 := sack2Counts[item]; ok2 {
				return item, nil
			}
		}
	}
	return '0', errors.New("No common item found")
}

func buildCounts(items string) map[rune]int {
	counts := make(map[rune]int)
	for _, item := range items {
		counts[item] += 1
	}
	return counts
}


// convert a-z -> 1-26 and A-Z -> 27 -> 52
func rune2Priority(c rune) int {
	ascii := int(c)
	// char is A-Z
	if ascii - 'a' < 0 {
		return ascii - 'A' + 1 + 26
	}
	// char is a-z
	return ascii - 'a' + 1
}

func buildCompartmentCounts(items string) (map[string]int, map[string]int) {
	compartment1 := make(map[string]int)
	compartment2 := make(map[string]int)

	length := len(items) / 2
	for _, item := range strings.Split(items[:length], "") {
		compartment1[item] += 1
	}
	for _, item := range strings.Split(items[length:], "") {
		compartment2[item] += 1
	}
	return compartment1, compartment2
}
