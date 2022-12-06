package day5

import (
	"strconv"
	"strings"
)

var emptyCrateString = "   "

type Crate struct {
	Id string
}

type Stack struct {
	Crates []*Crate
}

type StackYard struct {
	Stacks map[int]*Stack
}

func NewCrate(id string) *Crate {
	id = strings.Replace(id, "]", "", -1)
	id = strings.Replace(id, "[", "", -1)
	return &Crate{Id: id}
}

func (sStack *Stack) MoveCrates(tStack *Stack, c int) {
	moveCrates := sStack.Crates[len(sStack.Crates)-c:]
	sStack.Crates = sStack.Crates[:len(sStack.Crates)-c]

	tStack.Crates = append(tStack.Crates, moveCrates...)
}

func (s *Stack) CrateVals() []string {
	crateIds := make([]string, 0)
	for _, crate := range s.Crates {
		crateIds = append(crateIds, crate.Id)
	}
	return crateIds
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
			// c := sy.Stacks[crateMember].Crates
			sy.Stacks[crateMember].Crates = append(sy.Stacks[crateMember].Crates, NewCrate(crateString))
		}
		crateMember++
	}
	return sy
}
