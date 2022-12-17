package day12

import (
	"fmt"
	"strings"
	"testing"
)

func TestRange(t *testing.T) {
	i := 27
	for c := 'a'; c <= 'z'; c++ {
		fmt.Println(string(c))
		i--
	}
	fmt.Println('a' > 'c')
	fmt.Println(string('a' + 1))
}

func TestNewRuneGridFromLines(t *testing.T) {
	lines := strings.Split(sampleMap, "\n")
	mg := NewMapGrid(lines)
	for _, row := range mg.Grid {
		nVals := make([]string, 0)
		for _, n := range row {
			nVals = append(nVals, string(n.Val))
		}
		fmt.Println(nVals)
		fmt.Println()
	}
	n := mg.GetNode(Coords{X: 1, Y: 1})
	fmt.Println(n.ValStr())
}

func TestMod(t *testing.T) {
	fmt.Println(10%2 == 0)
}
