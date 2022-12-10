package day9

import (
	"fmt"
	"strings"
	"testing"
)

var TestCommands = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

var Part2Commands = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

var TestLines = strings.Split(TestCommands, "\n")
var Part2Lines = strings.Split(Part2Commands, "\n")

// func TestMoveHead(t *testing.T) {
// 	bs := InitBridgeSeekSim(1)
// 	fmt.Println(bs.KnotSlice)
// 	for _, cStr := range TestLines {
// 		c := NewCommandFromString(cStr)
// 		bs.MoveAndSeek(c)
// 	}

// 	fmt.Println(bs.TailVisMap)
// 	tailVisitCount := 0

// 	for range bs.TailVisMap.Map {
// 		tailVisitCount++
// 	}

// 	fmt.Println(bs.DrawTailMap())
// 	fmt.Println(tailVisitCount)
// }

// func TestPart2(t *testing.T) {
// 	bs := InitBridgeSeekSim(9)
// 	fmt.Println(bs.KnotSlice)
// 	for _, cStr := range Part2Lines {
// 		c := NewCommandFromString(cStr)
// 		bs.MoveAndSeek(c)
// 	}

// 	// fmt.Println(bs.TailVisMap)
// 	tailVisitCount := 0

// 	for range bs.TailVisMap.Map {
// 		tailVisitCount++
// 	}

// 	fmt.Println(tailVisitCount)

// }

// func TestDrawPart1(t *testing.T) {
// 	bs := InitBridgeSeekSim(1)
// 	bs.RunSim(TestLines)
// 	moveCount := 0
// 	fmt.Println(bs.DrawMapState(0))
// 	fmt.Println()
// 	fmt.Println("R 4")
// 	moveCount += 4
// 	fmt.Println(bs.DrawMapState(moveCount))
// 	fmt.Println()
// 	fmt.Println("U 4")
// 	moveCount += 4
// 	fmt.Println(bs.DrawMapState(moveCount))
// 	fmt.Println()
// 	fmt.Println("Last")
// 	fmt.Println(bs.DrawMapState(len(bs.BridgeStateMap) - 1))
// 	// fmt.Println(bs.TotalMoves)
// 	// fmt.Println(bs.BridgeStateMap[0][1].Name)
// }

func TestDrawPart2(t *testing.T) {
	bs := InitBridgeSeekSim(10)
	bs.RunSim(Part2Lines)
	moveCount := 0
	fmt.Println(bs.DrawMapState(0))
	fmt.Println()
	fmt.Println("R 4")
	moveCount += 4
	fmt.Println(bs.DrawMapState(moveCount))
	fmt.Println()
	fmt.Println("U 4")
	moveCount += 4
	fmt.Println(bs.DrawMapState(moveCount))
	fmt.Println()
	fmt.Println("Last")
	fmt.Println(bs.DrawMapState(len(bs.BridgeStateMap) - 1))
	// fmt.Println(bs.TotalMoves)
	// fmt.Println(bs.BridgeStateMap[0][1].Name)
}

// func TestMakeKnotCopies(t *testing.T) {
// 	k1 := &Knot{Position: &day8.Coords{X: 0, Y: 0}, Name: "H"}
// 	k2 := &Knot{Position: &day8.Coords{X: 0, Y: 0}, Name: "1"}
// 	knotPs := []*Knot{k1, k2}
// 	copy1 := MakeKnotsCopy(knotPs)
// 	fmt.Println("copy1:")
// 	fmt.Println(copy1[0].Name, copy1[0].Position.ToString())
// 	fmt.Println(copy1[0].Name, copy1[1].Position.ToString())
// 	k1.Move(NewCommandFromString("R 1"))
// 	k2.SeekHead(k1)
// 	fmt.Println("copy2:")
// 	copy2 := MakeKnotsCopy(knotPs)
// 	fmt.Println(copy2[0].Name, copy2[0].Position.ToString())
// 	fmt.Println(copy2[0].Name, copy2[1].Position.ToString())
// 	fmt.Println("Copy1 again:")
// 	fmt.Println(copy1[0].Name, copy1[0].Position.ToString())
// 	fmt.Println(copy1[0].Name, copy1[1].Position.ToString())
// 	// k1.Move(NewCommandFromString("R 1"))
// 	// k2.SeekHead(k1)
// 	// fmt.Println(k1.Position.ToString())
// 	// fmt.Println(k2.Position.ToString())
// }

// func TestDrawPart2(t *testing.T) {
// 	bs := InitBridgeSeekSim(10)
// 	for _, cStr := range Part2Lines {
// 		c := NewCommandFromString(cStr)
// 		bs.MoveAndSeek(c)
// 	}
// 	mStr := bs.DrawTailMap()
// 	fmt.Println(mStr)
// }

// func TestSeekHeadX(t *testing.T) {
// 	head := &Head{Position: &day8.Coords{X: 2, Y: 0}}
// 	tail := &Tail{Position: &day8.Coords{X: 0, Y: 0}}
// 	tail.SeekHead(head)
// 	fmt.Println(tail.Position.ToString())
// }

// func TestSeekHeadY(t *testing.T) {
// 	head := &Head{Position: &day8.Coords{X: 0, Y: 2}}
// 	tail := &Tail{Position: &day8.Coords{X: 0, Y: 0}}
// 	tail.SeekHead(head)
// 	fmt.Println(tail.Position.ToString())
// }

// func TestSeekHeadDiag(t *testing.T) {
// 	head := &Head{Position: &day8.Coords{X: 1, Y: 2}}
// 	tail := &Tail{Position: &day8.Coords{X: 0, Y: 0}}
// 	tail.SeekHead(head)
// 	fmt.Println(tail.Position.ToString())
// }
