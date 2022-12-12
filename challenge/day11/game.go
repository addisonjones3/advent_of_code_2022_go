package day11

import (
	"fmt"
	"strings"
)

type GameType string

const (
	A GameType = "A"
	B GameType = "B"
)

type Game struct {
	*MonkeyGroup
	RoundCount int
	GameType
}

func NewGame(mg *MonkeyGroup, gt GameType) *Game {
	return &Game{MonkeyGroup: mg, RoundCount: 0}
}

func (g *Game) PlayRound() {
	// Each monkey
	for id := 0; id < len(g.MonkeyMap); id++ {
		m := g.MonkeyMap[id]
		for _, item := range m.Items {
			m.Inspect(item, g.ModDivisor, g.GameType)
			testResult := m.Test(item)
			if testResult {
				g.MonkeyGroup.ThrowItem(m.id, m.TrueTargetMonkeyId, item)
			} else {
				g.MonkeyGroup.ThrowItem(m.id, m.FalseTargetMonkeyId, item)

			}
		}
	}
}

func (g *Game) PrintState() {
	builder := strings.Builder{}
	inspectBuilder := strings.Builder{}
	for id := 0; id < len(g.MonkeyMap); id++ {
		m := g.MonkeyMap[id]
		builder.WriteString(fmt.Sprintf("Monkey %d: ", m.id))
		builder.WriteString(m.Items.ToString())
		builder.WriteString("\n")
		inspectBuilder.WriteString(fmt.Sprintf("Monkey %d inspected %d items\n", m.id, m.InspectCount))
	}
	fmt.Println(builder.String())
	fmt.Println(inspectBuilder.String())
	fmt.Println()
	fmt.Println("Divisor:", g.ModDivisor)
}
