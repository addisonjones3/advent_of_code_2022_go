package day2

import (
	"fmt"
	"testing"
)

func TestNewShapeFromChar(t *testing.T) {
	rockObj := NewShapeFromChar("A")

	paperObj := NewShapeFromChar("B")

	RockVsPaper := rockObj.play(paperObj)
	PaperVsRock := paperObj.play(rockObj)
	fmt.Println("Rock vs. Paper: ", RockVsPaper)
	fmt.Println("Paper vs. Rock: ", PaperVsRock)
}

func TestSample(t *testing.T) {
	o1 := NewShapeFromChar("A")
	p1 := NewShapeFromChar("Y")
	o2 := NewShapeFromChar("B")
	p2 := NewShapeFromChar("X")
	o3 := NewShapeFromChar("C")
	p3 := NewShapeFromChar("Z")

	g1 := p1.play(o1)
	g2 := p2.play(o2)
	g3 := p3.play(o3)

	fmt.Println("Game 1: ", g1)
	fmt.Println("Game 2: ", g2)
	fmt.Println("Game 3: ", g3)
}

func TestPlayRounds(t *testing.T) {
	r1s := "A Y"
	r2s := "B X"
	r3s := "C Z"

	rounds := make([]PlayRound, 0)
	rounds = append(rounds, newPlayRound(r1s))
	rounds = append(rounds, newPlayRound(r2s))
	rounds = append(rounds, newPlayRound(r3s))

	score := 0
	for _, round := range rounds {
		score += round.Player.play(round.Opponent)
	}

	fmt.Println(score)

}

func TestOutcomeRounds(t *testing.T) {
	r1s := "A Y"
	r2s := "B X"
	r3s := "C Z"

	rounds := make([]OutcomeRound, 0)
	rounds = append(rounds, newOutcomeRound(r1s))
	rounds = append(rounds, newOutcomeRound(r2s))
	rounds = append(rounds, newOutcomeRound(r3s))

	score := 0
	for i, round := range rounds {
		roundScore := round.Opponent.playForOutcome(round.Outcome)
		fmt.Printf("Round %d score: %d\n", i, roundScore)
		score += round.Opponent.playForOutcome(round.Outcome)
	}

	fmt.Println(score)

}
