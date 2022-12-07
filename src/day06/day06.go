package day06

import (
	"fmt"
	"strings"
)


func Solve(input string) {
	position1, ok := part1(strings.TrimSpace(input))
	if !ok {
		panic("No marker found in string")
	}
	fmt.Println("Part 1 position", position1)

	position2, ok := part2(strings.TrimSpace(input))
	if !ok {
		panic("No marker found in string")
	}
	fmt.Println("Part 2 position", position2)
}

func part1(stream string) (int, bool) {
	for i, j, k, l := 0, 1, 2, 3; l < len(stream); i, j, k, l = i+1, j+1, k+1, l+1 {
		if stream[i] != stream[j] &&
			stream[i] != stream[k] &&
			stream[i] != stream[l] &&
			stream[j] != stream[k] &&
			stream[j] != stream[l] &&
			stream[k] != stream[l] {
			return l + 1, true
		}
	}
	return 0, false
}

func part2(stream string) (int, bool) {
	low := 0
	high := low + 14

	for high < len(stream) {
		if isUnique(stream[low:high]) {
			return high, true
		}
		low++
		high++
	}

	return -1, false
}

type nothing struct{}

func isUnique(str string) bool {
	runes := make(map[rune]nothing)
	for _, s := range str {
		runes[s] = nothing{}
	}
	if len(runes) == len(str) {
		return true
	}
	return false
}
