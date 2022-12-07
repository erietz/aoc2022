package main

import (
	"github.com/erietz/aoc2022/src/day06"
	"io"
	"os"
)

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	input := string(bytes)

	day06.Solve(input)
}
