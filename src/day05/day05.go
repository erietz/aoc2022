package day05

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/erietz/aoc2022/src"
)

type rule struct {
	numToMove int
	from int
	to int
}

type crateStacks map[int]*aoc.Stack[string]

func (cs crateStacks) String() string {
	s := ""
	for i := 0; i < len(cs); i++ {
		s += fmt.Sprintf("Crate %d %s\n", i + 1, cs[i].String())
	}
	return s
}

func Solve(input string) {
	blocksInput, rulesInput := parseInput(input)
	stacks := parseBlocks(blocksInput)
	rules := parseRules(rulesInput)
	message := part1(stacks, rules)
	fmt.Println(message)
}

func applyRules(blocks crateStacks, rules []rule) {
	for _, rule := range rules {
		from := blocks[rule.from - 1]
		to := blocks[rule.to - 1]
		for i := 0; i < rule.numToMove; i++ {
			val, ok := from.Pop()
			if !ok {
				panic("tried to pop from empty stack")
			}
			to.Push(val)
		}
	}
}


func parseInput(input string) ([]string, []string) {
	blocks := make([]string, 0)
	rules := make([]string, 0)

	for _, line := range strings.Split(input, "\n") {
		blockExp, err := regexp.MatchString(`\[[A-Z]\]`, line)
		if err != nil {
			panic(err)
		}
		if blockExp {
			blocks = append(blocks, line)
		}

		ruleExp, err := regexp.MatchString(`^move`, line)
		if err != nil {
			panic(err)
		}
		if ruleExp {
			rules = append(rules, line)
		}
	}
	return blocks, rules
}

func parseBlocks(blockLines []string) crateStacks {
	numStacks := (len(blockLines[0]) + 1)/4

	stacks := make(map[int]*aoc.Stack[string])
	for i := 0; i < numStacks; i++ {
		stacks[i] = &aoc.Stack[string]{}
	}

	for l := len(blockLines)-1; l>=0; l-- {
		for i, idx := 0, 1; i < numStacks; i, idx = i+1, idx+4 {
			crate := string(blockLines[l][idx])
			if crate != " " {
				stacks[i].Push(crate)
			}
		}
	}
	return stacks
}

func parseRules(rulesInput []string) []rule {
	rules := make([]rule, 0)
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for _, line := range rulesInput {
		res := re.FindStringSubmatch(line)
		rules = append(
			rules,
			rule{
				numToMove: parseInt(res[1]),
				from: parseInt(res[2]),
				to: parseInt(res[3]),
			},
		)
	}
	return rules
}

func parseInt(num string) int {
	n, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	return n
}

func part1(stacks crateStacks, rules []rule) string {
	applyRules(stacks, rules)
	fmt.Println(stacks)

	s := ""
	for i := 0; i < len(stacks); i++ {
		top, ok := stacks[i].Peek()
		if !ok {
			panic("stack shouldn't be empty")
		}
		s += top
	}
	return s
}
