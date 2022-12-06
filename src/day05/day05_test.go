package day05

import (
	"os"
	"testing"
)

func TestExample1(t *testing.T) {
	input, err := os.ReadFile("../../inputs/day05/test1.txt")
	if err != nil {
		panic(err)
	}

	blocksInput, rulesInput := parseInput(string(input))
	stacks := parseBlocks(blocksInput)
	rules := parseRules(rulesInput)

	message := part1(stacks, rules)
	if message != "CMZ" {
		t.Errorf("got %s, want CMZ", message)
	}
}
