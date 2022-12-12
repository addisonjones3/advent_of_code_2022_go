package day11

import (
	"sort"
)

// func day11PreBuild(path string) string {
// 	f, err := os.Open(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	s := bufio.NewScanner(f)
// 	s.Split(bufio.ScanLines)
// 	blockBuilder := strings.Builder{}
// 	for s.Scan() {
// 		blockBuilder.WriteString(fmt.Sprintf(s.Text(), "\n"))
// 	}
// 	return blockBuilder.String()
// }

func partA() int {
	monkeyGroup := ParseMonkeyGroupFromString(dayInput)
	game := NewGame(monkeyGroup, A)
	for i := 0; i < 20; i++ {
		game.PlayRound()
		game.PrintState()
	}
	inspectCounts := make(sort.IntSlice, 0)
	for _, m := range game.MonkeyGroup.MonkeyMap {
		inspectCounts = append(inspectCounts, m.InspectCount)
	}
	sort.Sort(sort.Reverse(inspectCounts))
	return inspectCounts[0] * inspectCounts[1]
}

func partB() int {
	monkeyGroup := ParseMonkeyGroupFromString(dayInput)
	game := NewGame(monkeyGroup, B)
	for i := 0; i < 10000; i++ {
		game.PlayRound()
		// game.PrintState()
	}
	inspectCounts := make(sort.IntSlice, 0)
	for _, m := range game.MonkeyGroup.MonkeyMap {
		inspectCounts = append(inspectCounts, m.InspectCount)
	}
	sort.Sort(sort.Reverse(inspectCounts))
	return inspectCounts[0] * inspectCounts[1]
}
