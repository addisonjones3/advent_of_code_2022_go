package day5

import (
	"fmt"
	"strings"
	"testing"
)

var testCrateYard = " 1   2   3   4   5   6   7   8   9"
var testCrateStringFull = "[D] [H] [L] [N] [N] [M] [D] [D] [B]"
var testCrateStringPartial = "[D] [H]     [N] [N] [M] [D] [D] [B]"

func TestScratchMove(t *testing.T) {
	sourceList := []string{"A", "B", "C"}
	targetList := []string{"D", "E"}

	targetList = append(targetList, sourceList[len(sourceList)-1:]...)
	sourceList = sourceList[:len(sourceList)-1]
	fmt.Println(sourceList)
	fmt.Println(targetList)
}

func TestMove(t *testing.T) {
	crateA := NewCrate("A")
	crateB := NewCrate("B")
	crateC := NewCrate("C")
	crateD := NewCrate("D")
	stackA := &Stack{Crates: []*Crate{crateA, crateB}}
	stackB := &Stack{Crates: []*Crate{crateC, crateD}}

	stackA.MoveCrates(stackB, 2)
	fmt.Println(stackA.CrateVals())
	fmt.Println(stackB.CrateVals())
}

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
