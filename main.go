package main

import (
	"os"
	"io"
	"github.com/erietz/aoc2022/src/day02"
)

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	input := string(bytes)

	day02.Solve(input)
}
