package day02

import (
	"fmt"
	"strings"
)

const (
	ROCK     = iota // 0
	PAPER           // 1
	SCIZZORS        // 2
)

// Each choice from the strategy guide
type RawChoice struct {
	elfChoice string // {A, B, C}
	myChoice  string // {X, Y, Z}
}

// Each choice from strategy guide interpreted as rock, paper, scizzors
type Choice struct {
	elfChoice int
	myChoice  int
}

// Score for each round
type Score struct {
	elfScore int
	myScore  int
}

type RawChoices []RawChoice
type Choices []Choice
type Scores []Score

func Solve(input string) {
	elfScore, myScore := part1(input)
	fmt.Println("Part 1")
	fmt.Printf("Elf Score: %v, My Score %v\n", elfScore, myScore)

	fmt.Println()

	elfScore, myScore = part2(input)
	fmt.Println("Part 2")
	fmt.Printf("Elf Score: %v, My Score %v\n", elfScore, myScore)
}

func part1(input string) (int, int) {
	rawChoices := parseStrategyGuide(input)
	choices := rawChoices.Parse(true)
	scores := choices.Tally()
	return scores.Sum()
}

func part2(input string) (int, int) {
	rawChoices := parseStrategyGuide(input)
	choices := rawChoices.Parse(false)
	scores := choices.Tally()
	return scores.Sum()
}

func (scores *Scores) Sum() (int, int) {
	elfScore := 0
	myScore := 0
	for _, score := range *scores {
		elfScore += score.elfScore
		myScore += score.myScore
	}
	return elfScore, myScore
}

func parseStrategyGuide(input string) RawChoices {
	rawChoices := []RawChoice{}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		choices := strings.Split(line, " ")
		rawChoice := RawChoice{
			elfChoice: choices[0],
			myChoice:  choices[1],
		}
		rawChoices = append(rawChoices, rawChoice)
	}
	return rawChoices
}

func (rc *RawChoice) ToChoicePart1() Choice {
	choice := Choice{}

	switch rc.elfChoice {
	case "A":
		choice.elfChoice = ROCK
	case "B":
		choice.elfChoice = PAPER
	case "C":
		choice.elfChoice = SCIZZORS
	}

	switch rc.myChoice {
	case "X":
		choice.myChoice = ROCK
	case "Y":
		choice.myChoice = PAPER
	case "Z":
		choice.myChoice = SCIZZORS
	}
	return choice
}

func (rc *RawChoice) ToChoicePart2() Choice {
	choice := Choice{}

	switch {
	case rc.elfChoice == "A" && rc.myChoice == "X": // lose
		choice.elfChoice = ROCK
		choice.myChoice = SCIZZORS
	case rc.elfChoice == "A" && rc.myChoice == "Y": // draw
		choice.elfChoice = ROCK
		choice.myChoice = ROCK
	case rc.elfChoice == "A" && rc.myChoice == "Z": // win
		choice.elfChoice = ROCK
		choice.myChoice = PAPER
	case rc.elfChoice == "B" && rc.myChoice == "X": // lose
		choice.elfChoice = PAPER
		choice.myChoice = ROCK
	case rc.elfChoice == "B" && rc.myChoice == "Y": // draw
		choice.elfChoice = PAPER
		choice.myChoice = PAPER
	case rc.elfChoice == "B" && rc.myChoice == "Z": // win
		choice.elfChoice = PAPER
		choice.myChoice = SCIZZORS
	case rc.elfChoice == "C" && rc.myChoice == "X": // lose
		choice.elfChoice = SCIZZORS
		choice.myChoice = PAPER
	case rc.elfChoice == "C" && rc.myChoice == "Y": // draw
		choice.elfChoice = SCIZZORS
		choice.myChoice = SCIZZORS
	case rc.elfChoice == "C" && rc.myChoice == "Z": // win
		choice.elfChoice = SCIZZORS
		choice.myChoice = ROCK
	}
	return choice
}

func (rawChoices *RawChoices) Parse(part1 bool) Choices {
	choices := []Choice{}
	for _, rc := range *rawChoices {
		var choice Choice
		if part1 {
			choice = rc.ToChoicePart1()
		} else {
			choice = rc.ToChoicePart2()
		}
		choices = append(choices, choice)
	}
	return choices
}

func (c *Choice) ToScore() Score {
	score := Score{}

	switch {
	case c.elfChoice == ROCK && c.myChoice == ROCK:
		score.elfScore = 1 + 3
		score.myScore = 1 + 3
	case c.elfChoice == ROCK && c.myChoice == PAPER:
		score.elfScore = 1 + 0
		score.myScore = 2 + 6
	case c.elfChoice == ROCK && c.myChoice == SCIZZORS:
		score.elfScore = 1 + 6
		score.myScore = 3 + 0
	case c.elfChoice == PAPER && c.myChoice == ROCK:
		score.elfScore = 2 + 6
		score.myScore = 1 + 0
	case c.elfChoice == PAPER && c.myChoice == PAPER:
		score.elfScore = 2 + 3
		score.myScore = 2 + 3
	case c.elfChoice == PAPER && c.myChoice == SCIZZORS:
		score.elfScore = 2 + 0
		score.myScore = 3 + 6
	case c.elfChoice == SCIZZORS && c.myChoice == ROCK:
		score.elfScore = 3 + 0
		score.myScore = 1 + 6
	case c.elfChoice == SCIZZORS && c.myChoice == PAPER:
		score.elfScore = 3 + 6
		score.myScore = 2 + 0
	case c.elfChoice == SCIZZORS && c.myChoice == SCIZZORS:
		score.elfScore = 3 + 3
		score.myScore = 3 + 3
	}
	return score
}

func (choices *Choices) Tally() Scores {
	scores := []Score{}
	for _, choice := range *choices {
		score := choice.ToScore()
		scores = append(scores, score)
	}
	return scores
}
