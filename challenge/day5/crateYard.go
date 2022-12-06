package day5

import (
	"strconv"
	"strings"
)

var emptyCrateString = "   "

type Crate struct {
	Id string
}

func NewCrate(id string) *Crate {
	id = strings.Replace(id, "]", "", -1)
	id = strings.Replace(id, "[", "", -1)
	return &Crate{Id: id}
}

type Move struct {
	Text             string
	MoveCount        int
	SourceStackIndex int
	TargetStackIndex int
}

func NewMoveFromString(s string) *Move {
	mv := &Move{Text: s}

	strSplit := strings.Split(s, " ")
	mvCount, _ := strconv.Atoi(strSplit[1])
	mv.MoveCount = mvCount
	sourceIndex, _ := strconv.Atoi(strSplit[3])
	mv.SourceStackIndex = sourceIndex
	targetIndex, _ := strconv.Atoi(strSplit[5])
	mv.TargetStackIndex = targetIndex
	return mv

}

type Stack struct {
	Crates []*Crate
}

func (s *Stack) CrateVals() []string {
	crateIds := make([]string, 0)
	for _, crate := range s.Crates {
		crateIds = append(crateIds, crate.Id)
	}
	return crateIds
}

type CrateMover9000 struct {
}

func (c *CrateMover9000) CraneMove(sStack *Stack, tStack *Stack, moveCount int) {
	i := 0
	for i < moveCount {
		tStack.Crates = append(tStack.Crates, sStack.Crates[len(sStack.Crates)-1])
		// fmt.Println("New tStack: ", tStack.CrateVals())

		if len(sStack.Crates) == 1 {
			// fmt.Println("sCrate reset")
			sStack.Crates = make([]*Crate, 0)
		} else {
			sStack.Crates = sStack.Crates[:len(sStack.Crates)-1]
			// fmt.Println("New sStack: ", sStack.CrateVals())
		}
		i++
	}
}

type CrateMover9001 struct {
}

func (c *CrateMover9001) CraneMove(sStack *Stack, tStack *Stack, moveCount int) {
	tStack.Crates = append(tStack.Crates, sStack.Crates[len(sStack.Crates)-moveCount:]...)
	sStack.Crates = sStack.Crates[:len(sStack.Crates)-moveCount]
}

type Crane interface {
	CraneMove(sStack *Stack, tStack *Stack, moveCount int)
}

type StackYard struct {
	Stacks    map[int]*Stack
	YardCrane Crane
}

func NewStackYardFromString(s string) *StackYard {
	stackIndices := strings.Fields(s)
	stackMap := make(map[int]*Stack)
	for _, iStr := range stackIndices {
		stackI, _ := strconv.Atoi(iStr)
		stackMap[stackI] = &Stack{}
	}

	return &StackYard{Stacks: stackMap}
}

func LoadStackYardFromString(s string, sy *StackYard) *StackYard {
	iInterval := 4
	crateMember := 1

	for i := 0; i < len(s); i = i + iInterval {
		crateString := s[i : i+iInterval-1]
		if crateString != emptyCrateString {
			sy.Stacks[crateMember].Crates = append(sy.Stacks[crateMember].Crates, NewCrate(crateString))
		}
		crateMember++
	}
	return sy
}

func (sy *StackYard) StackMove(mv *Move) {
	sy.YardCrane.CraneMove(sy.Stacks[mv.SourceStackIndex], sy.Stacks[mv.TargetStackIndex], mv.MoveCount)
}
