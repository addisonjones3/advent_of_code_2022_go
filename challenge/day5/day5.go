package day5

import (
	"bufio"
	"os"
)

func day5PreBuild(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	pairList := make([]string, 0)

	for s.Scan() {
		pairList = append(pairList, s.Text())
	}

	return pairList
}

func partA() int {
	return 1
}

func partB() int {
	return 1
}
