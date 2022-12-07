package day06

import (
	"fmt"
	"strings"
)

func Solve(input string) {
	position, ok := part1(strings.TrimSpace(input))
	if !ok {
		panic("No marker found in string")
	}
	fmt.Println("Part 1 position", position)
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
