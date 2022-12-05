package main

import (
	"os"
	"io"
	"github.com/erietz/aoc2022/src/day04"
)

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	input := string(bytes)

	day04.Solve(input)
}
