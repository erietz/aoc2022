package day02

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("../../inputs/day02/part1.txt")
	if err != nil {
		panic(err)
	}

	elfScore, myScore := part1(string(input))

	if elfScore != 9939 {
		t.Errorf("got %v, wanted %v", elfScore, 9939)
	}

	if myScore != 14827 {
		t.Errorf("got %v, wanted %v", myScore, 14827)
	}

}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("../../inputs/day02/part1.txt")
	if err != nil {
		panic(err)
	}

	elfScore, myScore := part2(string(input))

	if elfScore != 11178 {
		t.Errorf("got %v, wanted %v", elfScore, 11178)
	}

	if myScore != 13889 {
		t.Errorf("got %v, wanted %v", myScore, 13889)
	}

}
