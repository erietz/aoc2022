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
	s := stream
	length := len(s)
	for i, j, k, l := 0, 1, 2, 3; l < length; i, j, k, l = i+1, j+1, k+1, l+1 {
		if s[i] != s[j] &&
			s[i] != s[k] &&
			s[i] != s[l] &&
			s[j] != s[k] &&
			s[j] != s[l] &&
			s[k] != s[l] {
			return l + 1, true
		}
	}
	return -1, false
}
