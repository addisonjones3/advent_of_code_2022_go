package day4

import (
	"fmt"
	"testing"
)

func TestPairFromString(t *testing.T) {
	testString := "2-8,3-7"
	p := PairFromString(testString)
	fmt.Println(p.Pair1)
	fmt.Println(p.Pair2)
}

func TestFullOverlap(t *testing.T) {
	testString := "2-8,3-7"
	p := PairFromString(testString)
	p.SetOverlap()
	fmt.Println(p.FullOverlap)
}

func TestNoOverlap(t *testing.T) {
	testString := "2-6,5-7"
	p := PairFromString(testString)
	p.SetOverlap()
	fmt.Println(p.FullOverlap)
	fmt.Println(p.PartialOverlap)
}
