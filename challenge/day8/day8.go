package day8

import (
	"bufio"
	"os"
)

func day8PreBuild(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	treeInput := make([]string, 0)
	for s.Scan() {
		treeInput = append(treeInput, s.Text())
	}
	return treeInput
}

func partA(treeInput []string) int {
	f := NewForestFromLines(treeInput)
	return TreeVisibility(f)
}

func partB(treeInput []string) int {
	f := NewForestFromLines(treeInput)
	maxScenicScore := 0
	for x := 1; x < f.RowCount-1; x++ {
		for y := 1; y < f.ColCount-1; y++ {
			currTree := f.GetTree(x, y)
			currTree.SetScenicScore(f)
			if currTree.ScenicScore > maxScenicScore {
				maxScenicScore = currTree.ScenicScore
			}
		}
	}
	return maxScenicScore
}
