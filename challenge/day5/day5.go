package day5

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day5PreBuild(path string) ([]string, []*Move) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	crateYardInput := make([]string, 0)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		crateYardInput = append(crateYardInput, s.Text())
	}

	moveInput := make([]*Move, 0)
	for s.Scan() {
		moveInput = append(moveInput, NewMoveFromString(s.Text()))
	}

	return crateYardInput, moveInput
}

func partA(crateYardInput []string, moveInput []*Move) string {
	sy := NewStackYardFromString(crateYardInput[len(crateYardInput)-1])
	sy.YardCrane = &CrateMover9000{}

	// Load stacks in reverse order
	for i := len(crateYardInput) - 2; i >= 0; i-- {
		LoadStackYardFromString(crateYardInput[i], sy)
	}

	for i, mv := range moveInput {
		fmt.Println("Move", i)
		sy.StackMove(mv)
	}
	sb := strings.Builder{}
	for i := 1; i <= len(sy.Stacks); i++ {
		s := sy.Stacks[i]
		fmt.Println(s.CrateVals())
		sb.WriteString(s.CrateVals()[len(s.CrateVals())-1])
		// topStackString := append(topStackString, s)
	}
	topStackString := sb.String()
	return topStackString
}

func partB(crateYardInput []string, moveInput []*Move) string {
	sy := NewStackYardFromString(crateYardInput[len(crateYardInput)-1])
	sy.YardCrane = &CrateMover9001{}

	// Load stacks in reverse order
	for i := len(crateYardInput) - 2; i >= 0; i-- {
		LoadStackYardFromString(crateYardInput[i], sy)
	}

	for i, mv := range moveInput {
		fmt.Println("Move", i)
		sy.StackMove(mv)
	}
	sb := strings.Builder{}
	for i := 1; i <= len(sy.Stacks); i++ {
		s := sy.Stacks[i]
		fmt.Println(s.CrateVals())
		sb.WriteString(s.CrateVals()[len(s.CrateVals())-1])
		// topStackString := append(topStackString, s)
	}
	topStackString := sb.String()
	return topStackString
}
