package day08

import (
	"os"
	"testing"
)


func generateForest() Forest {
	bytes, err := os.ReadFile("../../inputs/day08/test1.txt")
	if err != nil {
		panic(err)
	}
	input := string(bytes)
	forest := parseInput(input)
	return forest
}

func TestFromLeft(t *testing.T) {
	forest := generateForest()
	visible := allocArray(forest)
	numTrees := fromLeft(forest, visible)

	if numTrees != 2 {
		t.Errorf("got %v, wanted %v", numTrees, 2)
	}
}

func TestFromRight(t *testing.T) {
	forest := generateForest()
	visible := allocArray(forest)
	numTrees := fromRight(forest, visible)

	if numTrees != 3 {
		t.Errorf("got %v, wanted %v", numTrees, 3)
	}
}

func TestFromTop(t *testing.T) {
	forest := generateForest()
	visible := allocArray(forest)
	numTrees := fromTop(forest, visible)

	if numTrees != 2 {
		t.Errorf("got %v, wanted %v", numTrees, 2)
	}
}

func TestFromBottom(t *testing.T) {
	forest := generateForest()
	visible := allocArray(forest)
	numTrees := fromBottom(forest, visible)

	if numTrees != 1 {
		t.Errorf("got %v, wanted %v", numTrees, 1)
	}
}

func TestEdges(t *testing.T) {
	forest := generateForest()
	numTrees := edges(forest)

	if numTrees != 16 {
		t.Errorf("got %v, wanted %v", numTrees, 16)
	}
}

func TestPart1(t *testing.T) {
	forest := generateForest()
	numTrees := part1(forest)

	if numTrees != 21 {
		t.Errorf("got %v, wanted %v", numTrees, 21)
	}
}

func TestViewingDistance(t *testing.T) {
	forest := generateForest()
	testCases := []struct{
		row int
		col int
		distance int
	}{
		{1, 2, 4},
		{3, 2, 8},
	}
	for _, tc := range testCases {
		distance := viewingDistance(forest, tc.row, tc.col)
		if distance != tc.distance {
			t.Errorf("got %v, wanted %v", distance, tc.distance)
		}
	}
}

func TestPart2(t *testing.T) {
	forest := generateForest()
	distance := part2(forest)

	if distance != 8 {
		t.Errorf("got %v, wanted %v", distance, 8)
	}
}

