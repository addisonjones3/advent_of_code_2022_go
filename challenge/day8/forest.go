package day8

import (
	"fmt"
	"strconv"
	"strings"
)

type Coords struct {
	X int
	Y int
}

type Tree struct {
	Coords      Coords
	Height      int
	ScenicScore int
	VisLeft     bool
	VisRight    bool
	VisUp       bool
	VisDown     bool
}

func (t *Tree) IsVisible() bool {
	return t.VisLeft || t.VisRight || t.VisUp || t.VisDown
}

func (t *Tree) SetScenicScore(f *Forest) {
	// Left
	leftScore := 0
	// fmt.Println("Left")
	for i := t.Coords.Y - 1; i >= 0; i-- {
		if nextTree, ok := f.TreeMap[ToCoordString(t.Coords.X, i)]; ok {
			// fmt.Println("x", i, "y", t.Coords.Y, "h", nextTree.Height)
			if nextTree.Height >= t.Height {
				leftScore++
				break
			} else if nextTree.Height < t.Height {
				leftScore++
			} else {
				break
			}
		} else {
			break
		}
	}
	// Right
	rightScore := 0
	// fmt.Println("Right")
	for i := t.Coords.Y + 1; i < f.ColCount; i++ {
		if nextTree, ok := f.TreeMap[ToCoordString(t.Coords.X, i)]; ok {
			// fmt.Println("x", i, "y", t.Coords.Y, "h", nextTree.Height)
			if nextTree.Height >= t.Height {
				rightScore++
				break
			} else if nextTree.Height < t.Height {
				rightScore++
			} else {
				break
			}
		} else {
			break
		}
	}
	// Up
	upScore := 0
	// fmt.Println("Up")
	for i := t.Coords.X - 1; i >= 0; i-- {
		if nextTree, ok := f.TreeMap[ToCoordString(i, t.Coords.Y)]; ok {
			// fmt.Println("x", t.Coords.X, "y", i, "h", nextTree.Height)
			if nextTree.Height >= t.Height {
				upScore++
				break
			} else if nextTree.Height < t.Height {
				upScore++
			}
		} else {
			break
		}
	}
	// Down
	downScore := 0
	// fmt.Println("Down")
	for i := t.Coords.X + 1; i < f.RowCount; i++ {
		if nextTree, ok := f.TreeMap[ToCoordString(i, t.Coords.Y)]; ok {
			// fmt.Println("x", t.Coords.X, "y", i, "h", nextTree.Height)
			if nextTree.Height >= t.Height {
				downScore++
				break
			} else if nextTree.Height < t.Height {
				downScore++
			} else {
				break
			}
		} else {
			break
		}
	}

	// fmt.Println("Left", leftScore)
	// fmt.Println("Right", rightScore)
	// fmt.Println("Up", upScore)
	// fmt.Println("Down", downScore)

	totalScore := leftScore * rightScore * upScore * downScore

	// fmt.Println("Total", totalScore)
	t.ScenicScore = totalScore
}

type Forest struct {
	TreeMap  map[string]*Tree
	ColCount int
	RowCount int
}

func (f *Forest) Print() string {
	rowBuilder := strings.Builder{}
	for x := 0; x < f.ColCount; x++ {
		for y := 0; y < f.RowCount; y++ {
			t := f.TreeMap[ToCoordString(x, y)]
			rowBuilder.WriteString(fmt.Sprint(t.Height))
		}
		rowBuilder.WriteString("\n")
	}
	return rowBuilder.String()
}

func (f *Forest) GetTree(x int, y int) *Tree {
	return f.TreeMap[ToCoordString(x, y)]
}

func NewTree(x int, y int, h int) *Tree {
	return &Tree{Coords: Coords{X: x, Y: y}, Height: h, VisLeft: false, VisRight: false}
}

func NewForestFromLines(lines []string) *Forest {
	forest := &Forest{TreeMap: map[string]*Tree{}}
	forest.RowCount = len(lines)
	forest.ColCount = len(lines[0])

	for rowIdx, row := range lines {
		for colIdx, height := range row {
			heightI, _ := strconv.Atoi(string(height))
			tree := NewTree(rowIdx, colIdx, heightI)
			forest.TreeMap[tree.Coords.ToString()] = tree
			if colIdx == 0 {
				tree.VisLeft = true
			}
			if colIdx == forest.ColCount-1 {
				tree.VisRight = true
			}
			if rowIdx == 0 {
				tree.VisUp = true
			}
			if rowIdx == forest.RowCount-1 {
				tree.VisDown = true
			}
		}
	}

	return forest
}

func (c *Coords) ToString() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

func ToCoordString(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func SetTreeRowVisibility(f *Forest, rowIdx int) {
	// Left calc
	leftTallest := f.TreeMap[ToCoordString(rowIdx, 0)]
	for colIdx := 0; colIdx < f.ColCount; colIdx++ {
		leftTree := f.TreeMap[ToCoordString(rowIdx, colIdx)]
		if leftTree.Height > leftTallest.Height {
			leftTree.VisLeft = true
			leftTallest = leftTree
		}
	}
	// Right calc
	rightTallest := f.TreeMap[ToCoordString(rowIdx, f.ColCount-1)]
	for colIdx := f.ColCount - 1; colIdx >= 0; colIdx-- {
		rightTree := f.TreeMap[ToCoordString(rowIdx, colIdx)]
		if rightTree.Height > rightTallest.Height {
			rightTree.VisRight = true
			rightTallest = rightTree
		}
	}
}

func SetTreeColVisibility(f *Forest, colIdx int) {
	// Up calc
	upTallest := f.TreeMap[ToCoordString(0, colIdx)]
	for rowIdx := 0; rowIdx < f.RowCount; rowIdx++ {
		upTree := f.TreeMap[ToCoordString(rowIdx, colIdx)]
		if upTree.Height > upTallest.Height {
			upTree.VisUp = true
			upTallest = upTree
		}
	}
	// Down calc
	downTallest := f.TreeMap[ToCoordString(f.RowCount-1, colIdx)]
	for rowIdx := f.ColCount - 1; rowIdx >= 0; rowIdx-- {
		downTree := f.TreeMap[ToCoordString(rowIdx, colIdx)]
		if downTree.Height > downTallest.Height {
			downTree.VisDown = true
			downTallest = downTree
		}
	}
}

func TreeVisibility(f *Forest) int {

	for rowIdx := 0; rowIdx < f.RowCount; rowIdx++ {
		SetTreeRowVisibility(f, rowIdx)
	}
	for colIdx := 0; colIdx < f.ColCount; colIdx++ {
		SetTreeColVisibility(f, colIdx)
	}

	visCount := 0
	for _, tree := range f.TreeMap {
		if tree.IsVisible() {
			visCount++
		}
	}

	return visCount
}
