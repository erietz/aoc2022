package day01

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 75501
	input, err := os.ReadFile("../../inputs/day01/part1.txt")
	if err != nil { panic(err) }

	actual := part1(getCalsPerElf(string(input)))

	if actual != expected {
		t.Errorf("expeced %v got %v\n", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 215594
	input, err := os.ReadFile("../../inputs/day01/part1.txt")
	if err != nil { panic(err) }

	actual := part2(getCalsPerElf(string(input)))

	if actual != expected {
		t.Errorf("expeced %v got %v\n", expected, actual)
	}
}
