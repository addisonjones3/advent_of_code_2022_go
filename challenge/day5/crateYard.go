package day5

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
	return &Crate{Id: id}
}

func (sStack *Stack) Move(tStack *Stack, c int) {
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
