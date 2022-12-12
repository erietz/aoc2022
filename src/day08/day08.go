package day08

import (
	"fmt"
	"strconv"
	"strings"
)

type Forest struct {
	Rows int
	Cols int
	Trees [][]int
}

func Solve(input string) {
	forest := parseInput(input)
	fmt.Println("part 1", part1(forest))
	fmt.Println("part 2", part2(forest))
}

func parseInput(input string) Forest {
	trees := make([][]int, 0)
	rows := 0

	for _, line := range strings.Fields(strings.TrimSpace(input)) {
		row := make([]int, 0)
		for _, char := range strings.Split(line, "") {
			height, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}
			row = append(row, height)
		}
		trees = append(trees, row)
		rows++
	}

	return Forest{
		Trees: trees,
		Cols: len(trees[0]),
		Rows: rows,
	}
}

func allocArray(forest Forest) [][]int {
	maxHeightForest := make([][]int, forest.Rows)
	for i := 0; i < forest.Rows; i++ {
		row := make([]int, forest.Cols)
		maxHeightForest[i] = row
	}
	return maxHeightForest
}

func fromLeft(forest Forest, visible [][]int) int {
	total := 0
	tmp := allocArray(forest)

	for i := 1; i < forest.Rows - 1; i++ {
		max := 0
		for j := 0; j < forest.Cols; j++ {
			if forest.Trees[i][j] > max {
				max = forest.Trees[i][j]
			}
			tmp[i][j] = max
		}
	}

	for i := 1; i < forest.Rows - 1; i++ {
		for j := 1; j < forest.Cols - 2; j++ {
			if forest.Trees[i][j] >= tmp[i][j] && forest.Trees[i][j] > tmp[i][j-1]{
				total += 1
				visible[i][j] = 1
			}
		}
	}

	return total
}

func fromRight(forest Forest, visible [][]int) int {
	total := 0
	tmp := allocArray(forest)

	for i := 1; i < forest.Rows - 1; i++ {
		max := 0
		for j := forest.Cols - 1; j >= 0 ; j-- {
			if forest.Trees[i][j] > max {
				max = forest.Trees[i][j]
			}
			tmp[i][j] = max
		}
	}

	for i := 1; i < forest.Rows - 1; i++ {
		for j := forest.Cols - 2; j >= 1; j-- {
			if forest.Trees[i][j] >= tmp[i][j] && forest.Trees[i][j] > tmp[i][j+1]{
				total += 1
				visible[i][j] = 1
			}
		}
	}

	return total
}

func fromTop(forest Forest, visible [][]int) int {
	total := 0
	tmp := allocArray(forest)

	for col := 1; col < forest.Cols - 1; col++ {
		max := 0
		for row := 0; row < forest.Rows; row++ {
			if forest.Trees[row][col] > max {
				max = forest.Trees[row][col]
			}
			tmp[row][col] = max
		}
	}

	for col := 1; col < forest.Cols - 1; col++ {
		for row := 1; row < forest.Cols - 2; row++ {
			if forest.Trees[row][col] >= tmp[row][col] && forest.Trees[row][col] > tmp[row-1][col]{
				total += 1
				visible[row][col] = 1
			}
		}
	}

	return total
}

func fromBottom(forest Forest, visible [][]int) int {
	total := 0
	tmp := allocArray(forest)

	for col := 1; col < forest.Cols - 1; col++ {
		max := 0
		for row := forest.Rows - 1; row >= 0; row-- {
			if forest.Trees[row][col] > max {
				max = forest.Trees[row][col]
			}
			tmp[row][col] = max
		}
	}

	for col := 1; col < forest.Cols - 1; col++ {
		for row := forest.Rows - 2; row >= 1; row-- {
			if forest.Trees[row][col] >= tmp[row][col] && forest.Trees[row][col] > tmp[row+1][col]{
				total += 1
				visible[row][col] = 1
			}
		}
	}

	return total
}

func edges(forest Forest) int {
	return 2*forest.Cols + 2*forest.Rows - 4
}

func part1(forest Forest) int {
	visible := allocArray(forest)

	fromLeft(forest, visible)
	fromRight(forest, visible)
	fromBottom(forest, visible)
	fromTop(forest, visible)

	inner := 0
	for i := 0; i < forest.Rows; i++ {
		for j := 0; j < forest.Cols; j++ {
			if visible[i][j] == 1 {
				inner += 1
			}
		}
	}

	return inner + edges(forest)
}

func viewingDistance(forest Forest, row, col int) int {
	left, right, up, down := 0, 0, 0, 0
	height := forest.Trees[row][col]

	// left
	for i := col - 1; i >= 0; i-- {
		left++
		if forest.Trees[row][i] >= height {
			break
		}
	}

	// right
	for i := col + 1; i < forest.Cols; i++ {
		right++
		if forest.Trees[row][i] >= height {
			break
		}
	}

	// up
	for i := row - 1; i >= 0; i-- {
		up++
		if forest.Trees[i][col] >= height {
			break
		}
	}

	// down
	for i := row + 1; i < forest.Rows; i++ {
		down++
		if forest.Trees[i][col] >= height {
			break
		}
	}

	return left * right * up * down
}

func part2(forest Forest) int {
	max := 0
	for i := 1; i < forest.Rows - 1; i++ {
		for j := 1; j < forest.Cols - 1; j++ {
			distance := viewingDistance(forest, i, j)
			if distance > max {
				max = distance
			}
		}
	}
	return max
}
