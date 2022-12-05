package day4

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	Pair1          []int
	Pair2          []int
	FullOverlap    bool
	PartialOverlap bool
}

func PairFromString(s string) *Pair {
	p := &Pair{}
	rawPairs := strings.Split(s, ",")
	rawPair1 := strings.Split(rawPairs[0], "-")
	p.Pair1 = make([]int, 0)
	for _, s := range rawPair1 {
		i, _ := strconv.Atoi(s)
		p.Pair1 = append(p.Pair1, i)
	}
	rawPair2 := strings.Split(rawPairs[1], "-")
	p.Pair2 = make([]int, 0)
	for _, s := range rawPair2 {
		i, _ := strconv.Atoi(s)
		p.Pair2 = append(p.Pair2, i)
	}
	return p
}

func (p *Pair) SetOverlap() {
	p1FullOverlap := p.Pair1[0] >= p.Pair2[0] && p.Pair1[1] <= p.Pair2[1]
	p2FullOverlap := p.Pair2[0] >= p.Pair1[0] && p.Pair2[1] <= p.Pair1[1]

	if p1FullOverlap || p2FullOverlap {
		p.FullOverlap = true
	} else {
		p.FullOverlap = false
	}

	p1PartialOverlap := p.Pair1[0] >= p.Pair2[0] && p.Pair1[0] <= p.Pair2[1]
	p2PartialOverlap := p.Pair2[0] >= p.Pair1[0] && p.Pair2[0] <= p.Pair1[1]

	if p1PartialOverlap || p2PartialOverlap {
		p.PartialOverlap = true
	} else {
		p.PartialOverlap = false
	}

}

func day4PreBuild(path string) []string {
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

func partA(pairList []string) int {
	pairs := make([]*Pair, 0)
	for _, pairString := range pairList {
		p := PairFromString(pairString)
		p.SetOverlap()
		pairs = append(pairs, p)
	}

	overlapCount := 0
	for _, p := range pairs {
		if p.FullOverlap {
			overlapCount++
		}
	}

	return overlapCount
}

func partB(pairList []string) int {
	pairs := make([]*Pair, 0)
	for _, pairString := range pairList {
		p := PairFromString(pairString)
		p.SetOverlap()
		pairs = append(pairs, p)
	}

	overlapCount := 0
	for _, p := range pairs {
		if p.PartialOverlap {
			overlapCount++
		}
	}

	return overlapCount
}
