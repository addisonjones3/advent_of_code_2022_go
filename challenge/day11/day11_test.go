package day11

import (
	"fmt"
	"sort"
	"testing"
)

// var sampleMG = ParseMonkeyGroupFromString(exampleBlock)

// func TestNewItem(t *testing.T) {
// 	i := NewItem(10)
// 	fmt.Println(i.id)
// 	fmt.Println(i.WorryValue)
// }

// func TestPrintItems(t *testing.T) {
// 	islice := ItemSlice{}
// 	islice = append(islice, NewItem(10))
// 	islice = append(islice, NewItem(11))
// 	islice = append(islice, NewItem(12))
// 	islice.PrintVals()
// }

// func TestParseItemsFromSring(t *testing.T) {
// 	items := ParseItemsFromSring(itemLine)
// 	for _, i := range items {
// 		fmt.Println(i.id, i.WorryValue)
// 	}
// }

// func TestParseOperationFromString(t *testing.T) {
// 	testOprString := "Operation: new = old * old"
// 	op := ParseOperationFromString(testOprString)
// 	fmt.Println(op)
// }

// func TestNewMonkeyFromString(t *testing.T) {
// 	scanner := bufio.NewScanner(strings.NewReader(exampleBlock))
// 	writer := strings.Builder{}
// 	monkeyGroups := make([]string, 0)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if line == "" {
// 			monkeyGroups = append(monkeyGroups, writer.String())
// 			writer.Reset()
// 		} else {
// 			writer.WriteString(line)
// 			writer.WriteString("\n")
// 		}
// 	}
// 	m0 := NewMonkeyFromString(monkeyGroups[0])
// 	fmt.Println(m0.id)
// 	for _, i := range m0.Items {
// 		fmt.Println(i)
// 	}
// }

// func TestNewMonkeyGroup(t *testing.T) {
// 	mg := ParseMonkeyGroupFromString(exampleBlock)
// 	for _, monkey := range mg.MonkeyMap {
// 		fmt.Println(monkey.id)
// 	}
// }

// func TestMGThrow(t *testing.T) {
// 	for i := 0; i < 2; i++ {
// 		fmt.Println(sampleMG.MonkeyMap[i].Items)
// 	}
// 	sampleMG.ThrowItem(0, 3, sampleMG.MonkeyMap[0].Items[0])
// 	for i := 0; i < 2; i++ {
// 		fmt.Println(sampleMG.MonkeyMap[i].Items)
// 	}
// }

// func TestInspect(t *testing.T) {
// 	m := sampleMG.MonkeyMap[0]
// 	i0 := m.Items[0]
// 	fmt.Println(i0.WorryValue)
// 	fmt.Println(m.InspectCount)
// 	m.Inspect(i0)
// 	fmt.Println(i0.WorryValue)
// 	fmt.Println(m.InspectCount)
// }

// func TestPlayRound(t *testing.T) {
// 	mg := ParseMonkeyGroupFromString(exampleBlock)
// 	game := &Game{MonkeyGroup: mg}
// 	game.PrintState()
// 	game.PlayRound()
// 	game.PrintState()
// }

func TestPlayNRounds(t *testing.T) {
	// mg := ParseMonkeyGroupFromString(exampleBlock, true)
	mg := ParseMonkeyGroupFromString(exampleBlock)
	game := NewGame(mg, B)
	game.PrintState()
	for i := 0; i < 10000; i++ {
		game.PlayRound()
	}
	game.PrintState()
	inspectCounts := make(sort.IntSlice, 0)
	for _, m := range game.MonkeyGroup.MonkeyMap {
		inspectCounts = append(inspectCounts, m.InspectCount)
	}
	sort.Sort(sort.Reverse(inspectCounts))
	fmt.Println(inspectCounts[0] * inspectCounts[1])
}

// func TestDiv(t *testing.T) {
// 	i1 := 400
// 	i1Mult := 3
// 	i1D := 10
// 	i2D := 32

// 	i1M := i1 * i1Mult
// 	mod := i1D * i2D

// 	fmt.Println(i1M)
// 	i1MultModded := i1M % mod
// 	fmt.Println(i1MultModded)

// 	if i1M%i1D == 0 {
// 		fmt.Println("Unmodded true")
// 	}
// 	if i1MultModded%i1D == 0 {
// 		fmt.Println("Modded true")
// 	}
// 	// fmt.Println(30 % (3 * 1000))
// }
