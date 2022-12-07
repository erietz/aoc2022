package day06

import (
	"fmt"
	"strings"
)

func Solve(input string) {
	position1, _ := findUniqueSequence(strings.TrimSpace(input), 4)
	position2, _ := findUniqueSequence(strings.TrimSpace(input), 14)

	// Part 1 position 1848
	fmt.Println("Part 1 position", position1)
	// Part 2 position 2308
	fmt.Println("Part 2 position", position2)
}

func findUniqueSequence(stream string, markerSize int) (int, bool) {
	low := 0
	high := low + markerSize

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
