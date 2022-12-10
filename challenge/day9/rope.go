package day9

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/addisonjones3/aoc22/challenge/day8"
)

type TailVisitMap struct {
	Map map[day8.Coords]int
}

type Direction string

const (
	right Direction = "R"
	left  Direction = "L"
	up    Direction = "U"
	down  Direction = "D"
)

const (
	tailNotVis = "."
	tailVis    = "#"
	start      = "s"
)

type Command struct {
	Direction Direction
	Magnitude int
}

func NewCommand(d Direction, m int) Command {
	return Command{Direction: d, Magnitude: m}
}
func NewCommandFromString(s string) Command {
	fields := strings.Fields(s)
	direction := fields[0]
	mag, _ := strconv.Atoi(fields[1])
	return NewCommand(Direction(direction), mag)
}

type Knot struct {
	Position *day8.Coords
	Name     string
}

type KnotCopy struct {
	Position day8.Coords
	Name     string
}

func (k *Knot) Move(c Command) {
	switch {
	case c.Direction == right:
		k.Position.X++
	case c.Direction == left:
		k.Position.X--
	case c.Direction == up:
		k.Position.Y++
	case c.Direction == down:
		k.Position.Y--
	}
}

func (t *Knot) SeekHead(h *Knot) {
	xDist := h.Position.X - t.Position.X
	xFloat := float64(xDist)
	yDist := h.Position.Y - t.Position.Y
	yFloat := float64(yDist)
	// both are within 1 or 0, don't move
	if (-1 <= xDist && xDist <= 1) && (-1 <= yDist && yDist <= 1) {
		return
	}
	// X has moved far enough away to chase
	if math.Abs(xFloat) == 2 {
		if xDist > 0 {
			t.Position.X = h.Position.X - 1
		}
		if xDist < 0 {
			t.Position.X = h.Position.X + 1
		}
		if math.Abs(xFloat) == 2 && math.Abs(yFloat) == 1 {
			t.Position.Y = h.Position.Y
		}
	}
	// Y has moved far enough away to chase
	if math.Abs(yFloat) == 2 {
		if yDist > 0 {
			t.Position.Y = h.Position.Y - 1
		}
		if yDist < 0 {
			t.Position.Y = h.Position.Y + 1
		}
		if math.Abs(yFloat) == 2 && math.Abs(float64(xDist)) == 1 {
			t.Position.X = h.Position.X
		}
	}
}

type BridgeSeekSim struct {
	Head           *Knot
	Knots          []*Knot
	Tail           *Knot
	TailVisMap     *TailVisitMap
	BridgeStateMap map[int][]KnotCopy
	TotalMoves     int
	MaxX           int
	MinX           int
	MaxY           int
	MinY           int
}

func MakeKnotsCopy(ks []*Knot) []KnotCopy {
	knotCopies := make([]KnotCopy, 0)
	for _, k := range ks {
		kCopy := *k
		knotCopies = append(knotCopies, KnotCopy{Position: *kCopy.Position, Name: kCopy.Name})
	}
	return knotCopies
}

func InitBridgeSeekSim(tailCount int) *BridgeSeekSim {
	head := &Knot{Position: &day8.Coords{X: 0, Y: 0}, Name: "H"}
	knotSlice := []*Knot{head}
	for i := 1; i < tailCount; i++ {
		knotSlice = append(knotSlice, &Knot{Position: &day8.Coords{X: 0, Y: 0}, Name: fmt.Sprint(i)})
	}
	tail := knotSlice[len(knotSlice)-1]
	tvMap := &TailVisitMap{Map: make(map[day8.Coords]int)}
	tvMap.Map[*tail.Position]++
	bridgeStateMap := map[int][]KnotCopy{0: MakeKnotsCopy(knotSlice)}
	return &BridgeSeekSim{Head: head, Knots: knotSlice, Tail: tail, TailVisMap: tvMap, MaxX: 0, MinX: 0, MaxY: 0, MinY: 0, BridgeStateMap: bridgeStateMap, TotalMoves: 1}
}

func (bs *BridgeSeekSim) MoveAndSeek(c Command) {
	for i := 0; i < c.Magnitude; i++ {
		bs.Head.Move(c)
		for j := 0; j < len(bs.Knots)-1; j++ {
			bs.Knots[j+1].SeekHead(bs.Knots[j])
		}
		bs.TailVisMap.Map[*bs.Tail.Position]++
		bs.BridgeStateMap[bs.TotalMoves] = MakeKnotsCopy(bs.Knots)
		bs.TotalMoves++
	}
	switch {
	case bs.Head.Position.X > bs.MaxX:
		bs.MaxX = bs.Head.Position.X
	case bs.Head.Position.X < bs.MinX:
		bs.MinX = bs.Head.Position.X
	case bs.Head.Position.Y > bs.MaxY:
		bs.MaxY = bs.Head.Position.Y
	case bs.Head.Position.Y < bs.MinY:
		bs.MinY = bs.Head.Position.Y
	}
}

func (bs *BridgeSeekSim) DrawTailMap() string {
	mapStringRows := make([]string, 0)
	for row := bs.MinY; row <= bs.MaxX; row++ {
		rowBuilder := &strings.Builder{}
		for col := bs.MinX; col <= bs.MaxX; col++ {
			if _, ok := bs.TailVisMap.Map[day8.Coords{X: col, Y: row}]; ok {
				if col == 0 && row == 0 {
					rowBuilder.WriteString(start)
				} else {
					rowBuilder.WriteString(tailVis)
				}
			} else {
				rowBuilder.WriteString(tailNotVis)
			}
		}
		rowString := rowBuilder.String()
		mapStringRows = append([]string{rowString}, mapStringRows...)
	}
	return strings.Join(mapStringRows, "\n")
}

func (bs *BridgeSeekSim) DrawMapState(move int) string {
	mapState := bs.BridgeStateMap[move]
	mapStringCells := make([][]string, 0)
	for i := bs.MaxY; i >= bs.MinY; i-- {
		currRow := make([]string, 0)
		for j := bs.MinX; j <= bs.MaxX; j++ {
			currRow = append(currRow, tailNotVis)
		}
		mapStringCells = append(mapStringCells, currRow)
	}

	rowStrings := make([]string, 0)

	for i := len(bs.Knots) - 1; i >= 0; i-- {
		k := mapState[i]
		mapStringCells[bs.MaxY-k.Position.Y][k.Position.X-bs.MinX] = k.Name
	}
	for _, row := range mapStringCells {
		rowStrings = append(rowStrings, strings.Join(row, ""))
	}
	return strings.Join(rowStrings, "\n")
}

func (bs *BridgeSeekSim) RunSim(commandStrs []string) {
	for _, cStr := range commandStrs {
		c := NewCommandFromString(cStr)
		bs.MoveAndSeek(c)
	}
}
