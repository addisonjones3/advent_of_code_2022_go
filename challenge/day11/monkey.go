package day11

import (
	"math"
	"strings"
)

type Monkey struct {
	id    int
	Items ItemSlice
	*Operation
	*TestOperation
	TrueTargetMonkeyId  int
	FalseTargetMonkeyId int
	InspectCount        int
	DivideOnInspect     bool
}

func NewMonkeyFromString(s string) *Monkey {
	lines := strings.Split(s, "\n")
	id := ParseIdFromString(lines[0])
	items := ParseItemsFromSring(lines[1])
	op := ParseOperationFromString(lines[2])
	to := ParseTestOperationFromString(lines[3])
	trueTgtId, falseTgtId := ParseTargetMonkeyIds(lines[4:])
	return &Monkey{
		id:                  id,
		Items:               items,
		Operation:           op,
		TestOperation:       to,
		TrueTargetMonkeyId:  trueTgtId,
		FalseTargetMonkeyId: falseTgtId,
		DivideOnInspect:     true,
	}
}

func NewMonkeyFromStringNoDivide(s string) *Monkey {
	lines := strings.Split(s, "\n")
	id := ParseIdFromString(lines[0])
	items := ParseItemsFromSring(lines[1])
	op := ParseOperationFromString(lines[2])
	to := ParseTestOperationFromString(lines[3])
	trueTgtId, falseTgtId := ParseTargetMonkeyIds(lines[4:])
	return &Monkey{
		id:                  id,
		Items:               items,
		Operation:           op,
		TestOperation:       to,
		TrueTargetMonkeyId:  trueTgtId,
		FalseTargetMonkeyId: falseTgtId,
		DivideOnInspect:     false,
	}
}
func (m *Monkey) Throw(item *Item, tMonkey *Monkey) {
	tMonkey.Items = append(tMonkey.Items, item)
	m.Items = m.Items[1:]
}

func (m *Monkey) Test(i *Item) bool {
	if i.WorryValue%m.TestOperation.Val == 0 {
		return true
	} else {
		return false
	}
}

func (m *Monkey) Inspect(i *Item, modDivisor int, gt GameType) {
	// fmt.Println(m.id, "inspecting", i.WorryValue)
	// fmt.Println(m.Operator, m.Value)
	var inspectVal int
	var inspectedVal int
	if m.SelfOp {
		inspectVal = i.WorryValue
	} else {
		inspectVal = m.Operation.Value
	}
	switch m.Operator {
	case mult:
		inspectedVal = i.WorryValue * inspectVal
	case add:
		inspectedVal = i.WorryValue + inspectVal
	}
	// fmt.Println("new worryval", opFloatVal)
	if gt == A {
		i.WorryValue = int(math.Floor(float64(inspectedVal) / 3))
	} else {
		i.WorryValue = inspectedVal % modDivisor
	}
	m.InspectCount++
}

type MonkeyGroup struct {
	MonkeyMap  map[int]*Monkey
	ModDivisor int
}

func (mg *MonkeyGroup) ThrowItem(sourceMId int, targetMId int, item *Item) {
	sMonkey := mg.MonkeyMap[sourceMId]
	tMonkey := mg.MonkeyMap[targetMId]
	sMonkey.Throw(item, tMonkey)
}
