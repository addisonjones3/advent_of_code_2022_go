package day5

import (
	"fmt"
	"testing"
)

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

	stackA.Move(stackB, 2)
	fmt.Println(stackA.CrateVals())
	fmt.Println(stackB.CrateVals())
}
