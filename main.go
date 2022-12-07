package main

import (
	"github.com/erietz/aoc2022/src/day07"
	"io"
	"os"
)

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	input := string(bytes)

	day07.Solve(input)
}
