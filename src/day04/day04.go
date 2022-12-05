package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type section struct {
	elf1SecMin int
	elf1SecMax int
	elf2SecMin int
	elf2SecMax int
}

func Solve(input string) {
	sections := parseSections(input)
	sectionsContained := part1(sections)
	fmt.Println("Sections contained in another section:", sectionsContained)
}

func parseSections(input string) []section {
	sections := make([]section, 0)
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		secs := strings.Split(line, ",")
		elf1 := strings.Split(secs[0], "-")
		elf2 := strings.Split(secs[1], "-")

		sections = append(
			sections,
			section{
				elf1SecMin: parseInt(elf1[0]),
				elf1SecMax: parseInt(elf1[1]),
				elf2SecMin: parseInt(elf2[0]),
				elf2SecMax: parseInt(elf2[1]),
			},
		)
	}
	return sections
}

func parseInt(n string) int {
	num, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	return num
}

func isContained(sec section) bool {
	if sec.elf1SecMin <= sec.elf2SecMin && sec.elf1SecMax >= sec.elf2SecMax {
		return true
	}
	if sec.elf2SecMin <= sec.elf1SecMin && sec.elf2SecMax >= sec.elf1SecMax {
		return true
	}
	return false
}

func part1(sections []section) int {
	sectionsContained := 0
	for _, sec := range sections {
		if isContained(sec) {
			sectionsContained += 1
		}
	}
	return sectionsContained
}
