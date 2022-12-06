package day5

import (
	"fmt"
	"strings"
	"testing"
)

var testCrateYard = " 1   2   3   4   5   6   7   8   9"
var testCrateStringFull = "[D] [H] [L] [N] [N] [M] [D] [D] [B]"
var testCrateStringPartial = "[D] [H]     [N] [N] [M] [D] [D] [B]"
var testMoveString = "move 2 from 1 to 2"

func TestYardParse(t *testing.T) {
	// crateList = strings.Replace(crateList, " ", "", -1)
	numList := strings.Fields(testCrateYard)
	fmt.Println(numList)
}

func TestCrateParseFull(t *testing.T) {
	crateList := strings.Split(testCrateStringFull, " ")
	fmt.Println(crateList)
}

func TestLoadStackYardFromString(t *testing.T) {
	sy := NewStackYardFromString(testCrateYard)
	sy = LoadStackYardFromString(testCrateStringPartial, sy)
	sy = LoadStackYardFromString(testCrateStringPartial, sy)
	for _, s := range sy.Stacks {
		fmt.Println(s.CrateVals())
	}
}

// func TestNewStackYardFromString(t *testing.T) {
// 	sy := NewStackYardFromString(testCrateYard)
// 	for _, s := range sy.Stacks {
// 		fmt.Println(s.CrateVals())
// 	}
// 	sy = LoadStackYardFromString()
// }

// func TestParseMoveString(t *testing.T) {
// 	mv := NewMoveFromString(testMoveString)
// 	fmt.Println(mv.MoveCount, mv.SourceStackIndex, mv.TargetStackIndex)
// }

// func TestMoveStrToMove(t *testing.T) {
// 	mv := NewMoveFromString(testMoveString)
// 	crateA := NewCrate("A")
// 	crateB := NewCrate("B")
// 	crateC := NewCrate("C")
// 	crateD := NewCrate("D")
// 	stackA := &Stack{Crates: []*Crate{crateA, crateB}}
// 	stackB := &Stack{Crates: []*Crate{crateC, crateD}}

// 	stacks := map[int]*Stack{1: stackA, 2: stackB}
// 	sy := &StackYard{Stacks: stacks}

// 	sy.StackMove(mv)
// 	fmt.Println(sy.Stacks[1].CrateVals())
// 	fmt.Println(sy.Stacks[2].CrateVals())
// }

func TestDay5PreBuild(t *testing.T) {
	crateInput, moveList := day5PreBuild("input.txt")
	fmt.Println(crateInput)
	for _, mv := range moveList {
		fmt.Println(mv.Text)
	}
}

func TestCM9001Move(t *testing.T) {
	mv := NewMoveFromString(testMoveString)
	crateA := NewCrate("A")
	crateB := NewCrate("B")
	crateC := NewCrate("C")
	crateD := NewCrate("D")
	stackA := &Stack{Crates: []*Crate{crateA, crateB}}
	stackB := &Stack{Crates: []*Crate{crateC, crateD}}

	stacks := map[int]*Stack{1: stackA, 2: stackB}
	sy := &StackYard{Stacks: stacks}
	sy.YardCrane = &CrateMover9001{}

	// sy.YardCrane.CraneMove(sy.Stacks[1], sy.Stacks[2], mv.)
	sy.StackMove(mv)
	fmt.Println(sy.Stacks[1].CrateVals())
	fmt.Println(sy.Stacks[2].CrateVals())
}
