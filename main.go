package main

import (
	"github.com/erietz/aoc2022/src/day08"
	"io"
	"os"
)

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	input := string(bytes)

	day08.Solve(input)
}
