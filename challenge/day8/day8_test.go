package day8

import (
	"fmt"
	"strings"
	"testing"
)

var TestForest = `30373
25512
65332
33549
35390`

var TestLines = strings.Split(TestForest, "\n")

// func TestBuildForest(t *testing.T) {
// 	f := &Forest{TreeMap: make(map[string]*Tree)}
// 	tree := NewTree(0, 0, 3)
// 	f.TreeMap[tree.Coords.ToString()] = tree
// 	fmt.Println(f.TreeMap["0,0"].Coords)
// }

// func TestBuildForestFromLines(t *testing.T) {
// 	rowCount := len(TestLines)
// 	colCount := len(TestLines[0])
// 	fmt.Println("Row count", rowCount)
// 	fmt.Println("Col count", colCount)
// 	f := NewForestFromLines(TestLines)
// 	for _, tree := range f.TreeMap {
// 		fmt.Println(tree.Height)
// 	}
// }

// func TestCalcForestVis(t *testing.T) {
// 	f := NewForestFromLines(TestLines)
// 	// fmt.Println(f.TreeMap[ToCoordString(4, f.ColCount-1)])
// 	fVis := TreeVisibility(f)
// 	// tl5want := Tree{Coords: Coords{1, 1}, Height: 5, VisLeft: true, VisRight: false, VisUp: true, VisDown: false}
// 	// tm5want := Tree{Coords: Coords{1, 2}, Height: 5, VisLeft: false, VisRight: true, VisUp: true, VisDown: false}
// 	// lm5want := Tree{Coords: Coords{2, 1}, Height: 5, VisLeft: false, VisRight: true, VisUp: false, VisDown: false}
// 	// br4want := Tree{Coords: Coords{3, 3}, Height: 4, VisLeft: false, VisRight: false, VisUp: false, VisDown: false}
// 	rbc0want := Tree{Coords: Coords{4, 4}, Height: 0, VisLeft: false, VisRight: true, VisUp: false, VisDown: true}
// 	for i := 0; i < f.RowCount; i++ {
// 		for j := 0; j < f.ColCount; j++ {
// 			tree := f.TreeMap[ToCoordString(i, j)]
// 			if i == 1 && j == 1 {
// 				// fmt.Println(*tree == tl5want)
// 				// fmt.Println(tree)
// 				// fmt.Println(tl5want)
// 			}
// 			if i == 1 && j == 2 {
// 				// fmt.Println(*tree == tm5want)
// 				// fmt.Println(tree)
// 				// fmt.Println(tm5want)
// 			}
// 			if i == 2 && j == 1 {
// 				// fmt.Println(*tree == lm5want)
// 				// fmt.Println(tree)
// 				// fmt.Println(lm5want)
// 			}
// 			if i == 4 && j == 4 {
// 				fmt.Println(*tree == rbc0want)
// 				fmt.Println(tree)
// 				fmt.Println(rbc0want)
// 			}
// 			// fmt.Println(tree.VisLeft, tree.VisRight, tree.VisUp, tree.VisDown, tree.Coords.ToString())
// 		}
// 	}

// 	fmt.Println(fVis)
// }

func TestForestPrint(t *testing.T) {
	f := NewForestFromLines(TestLines)
	fmt.Println(f.Print())
}

func TestScenicScore(t *testing.T) {
	f := NewForestFromLines(TestLines)
	// fVis := TreeVisibility(f)
	// fmt.Println(fVis)
	// f.TreeMap["1,2"].SetScenicScore(f)
	f.TreeMap["3,2"].SetScenicScore(f)
}
