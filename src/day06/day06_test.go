package day06

import (
	"fmt"
	"testing"
)

func TestExamplesPart1(t *testing.T) {
	testCases := []struct{
		stream string
		position int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s | %d", tc.stream, tc.position), func(t *testing.T) {
			position, ok := part1(tc.stream)
			if !ok {
				t.Errorf("got %v, wanted %v", ok, true)
			}
			if position != tc.position {
				t.Errorf("got %v, wanted %v", position, tc.position)
			}
		})
	}
}

func TestExamplesPart2(t *testing.T) {
	testCases := []struct{
		stream string
		position int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s | %d", tc.stream, tc.position), func(t *testing.T) {
			position, ok := findUniqueSequence(tc.stream)
			if !ok {
				t.Errorf("got %v, wanted %v", ok, true)
			}
			if position != tc.position {
				t.Errorf("got %v, wanted %v", position, tc.position)
			}
		})
	}
}
