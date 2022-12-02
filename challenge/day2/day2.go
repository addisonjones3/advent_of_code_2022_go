package day2

import (
	"bufio"
	"os"
	"strings"
)

const (
	Rock     = "Rock"
	Paper    = "Paper"
	Scissors = "Scissors"
)

const (
	Play    = "Play"
	Outcome = "Outcome"
)

const (
	Win  = 6
	Loss = 0
	Draw = 3
)

type Shape struct {
	Name  string
	Value int
	Beats string
	Loses string
}

// type PlayRound []*Shape
type PlayRound struct {
	Player   *Shape
	Opponent *Shape
}

type OutcomeRound struct {
	Opponent *Shape
	Outcome  int
}

func newPlayRound(s string) PlayRound {
	playSlice := strings.Split(s, " ")
	shape0 := NewShapeFromChar(playSlice[0])
	shape1 := NewShapeFromChar(playSlice[1])
	return PlayRound{Opponent: shape0, Player: shape1}
}

func newOutcomeRound(s string) OutcomeRound {
	outcomeSlice := strings.Split(s, " ")
	shape0 := NewShapeFromChar(outcomeSlice[0])
	outcome := OutcomeMap[outcomeSlice[1]]
	return OutcomeRound{shape0, outcome}
}

var OutcomeMap = map[string]int{
	"X": Loss,
	"Y": Draw,
	"Z": Win,
}

var GuideCharMap = map[string]*Shape{
	"A": NewShape(Rock),
	"X": NewShape(Rock),
	"B": NewShape(Paper),
	"Y": NewShape(Paper),
	"C": NewShape(Scissors),
	"Z": NewShape(Scissors),
}

func (p *Shape) play(o *Shape) int {
	if p.Name == o.Name {
		return p.Value + Draw
	}
	if p.Beats == o.Name {
		return p.Value + Win
	}
	if p.Name == o.Beats {
		return p.Value + Loss
	}

	return 0
}

func (s *Shape) playForOutcome(o int) int {
	if o == Draw {
		pShape := NewShape(s.Name)
		return pShape.play(s)
	}
	if o == Win {
		pShape := NewShape(s.Loses)
		return pShape.play(s)
	}
	if o == Loss {
		pShape := NewShape(s.Beats)
		return pShape.play(s)
	}
	return 0
}

func NewShape(name string) *Shape {
	switch {
	case name == Rock:
		return &Shape{Name: Rock, Value: 1, Beats: Scissors, Loses: Paper}
	case name == Paper:
		return &Shape{Name: Paper, Value: 2, Beats: Rock, Loses: Scissors}
	case name == Scissors:
		return &Shape{Name: Scissors, Value: 3, Beats: Paper, Loses: Rock}
	}
	return nil
}

func NewShapeFromChar(c string) *Shape {
	if shape, ok := GuideCharMap[c]; ok {
		return shape
	}
	return nil
}

func day2PreBuildOutcome(path string) []OutcomeRound {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	outcomeRoundList := make([]OutcomeRound, 0)
	for s.Scan() {
		round := newOutcomeRound(s.Text())
		outcomeRoundList = append(outcomeRoundList, round)
	}

	return outcomeRoundList
}
func day2PreBuild(path string) []PlayRound {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	playRoundList := make([]PlayRound, 0)
	for s.Scan() {
		round := newPlayRound(s.Text())
		playRoundList = append(playRoundList, round)
	}

	return playRoundList
}

func partA(rl []PlayRound) int {
	score := 0
	for _, round := range rl {
		score += round.Player.play(round.Opponent)
	}

	return score
}

func partB(orl []OutcomeRound) int {
	score := 0
	for _, round := range orl {
		score += round.Opponent.playForOutcome(round.Outcome)
	}

	return score
}
