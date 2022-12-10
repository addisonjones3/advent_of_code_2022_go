package day9

import (
	"bufio"
	"fmt"
	"os"
)

func day9PreBuild(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	bridgeInput := make([]string, 0)
	for s.Scan() {
		bridgeInput = append(bridgeInput, s.Text())
	}
	return bridgeInput
}

func partA(bridgeInput []string) int {
	bs := InitBridgeSeekSim(1)
	for _, cStr := range bridgeInput {
		c := NewCommandFromString(cStr)
		bs.MoveAndSeek(c)
	}

	tailVisitCount := 0
	for range bs.TailVisMap.Map {
		tailVisitCount++
	}
	fmt.Println(bs.DrawTailMap())
	return tailVisitCount
}

func partB(bridgeInput []string) int {
	bs := InitBridgeSeekSim(10)
	for _, cStr := range bridgeInput {
		c := NewCommandFromString(cStr)
		bs.MoveAndSeek(c)
	}

	tailVisitCount := 0
	for range bs.TailVisMap.Map {
		tailVisitCount++
	}
	// fmt.Println(bs.DrawTailMap())
	return tailVisitCount
}
