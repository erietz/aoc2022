package main

import (
	"os"
	"io"
	"github.com/erietz/aoc2022/src/day01"
)

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	input := string(bytes)

	day01.Solve(input)
}
