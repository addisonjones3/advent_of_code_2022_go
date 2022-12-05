package day3

import (
	"fmt"
	"testing"
)

func TestNewPackFromString(t *testing.T) {
	testItems := "vJrwpWtwJgWrhcsFMMfFFhFp"

	p := NewPackFromString(testItems)

	fmt.Println(p.Compartment1)
	fmt.Println(p.Compartment2)
	fmt.Println(p.SharedItem)
	fmt.Println(ItemValues[p.SharedItem])

}

func TestPackGroupsFromList(t *testing.T) {
	testPack := &Pack{}
	im := map[string]int{"a": 1}
	testPack.Compartment1 = &Compartment{itemMap: im}
	p := &Pack{Compartment1: &Compartment{im}}
	TestPacks := make([]*Pack, 1)
	for i := 1; i < 10; i++ {
		TestPacks = append(TestPacks, p)
	}
	PackGroupsFromList(1, TestPacks)
}
