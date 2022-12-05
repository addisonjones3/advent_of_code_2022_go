package day3

import (
	"bufio"
	"os"
	"strings"
)

type Compartment struct {
	itemMap map[string]int
}

type Pack struct {
	Compartment1        *Compartment
	Compartment2        *Compartment
	CombinedCompartment *Compartment
	SharedItem          string
}

type PackGroup struct {
	Packs []*Pack
	Badge string
}

func NewPackFromString(s string) *Pack {
	l := len(s) / 2
	ci := strings.Split(s, "")
	i1 := strings.Split(s[:l], "")
	i2 := strings.Split(s[l:], "")

	comp1 := NewCompartment(i1)
	comp2 := NewCompartment(i2)
	combinedComp := NewCompartment(ci)

	sharedItem := ""
	for k := range comp1.itemMap {
		if _, ok := comp2.itemMap[k]; ok {
			sharedItem = k
			break
		}
	}
	return &Pack{CombinedCompartment: combinedComp, Compartment1: comp1, Compartment2: comp2, SharedItem: sharedItem}
}

func PackGroupsFromList(groupSize int, packs []*Pack) []*PackGroup {
	packGroups := make([]*PackGroup, 0)

	i := 1
	groupCounter := 0
	currPackGroup := &PackGroup{}
	for i < len(packs)+1 {
		currPackGroup.Packs = append(currPackGroup.Packs, packs[i-1])
		if i%groupSize == 0 {
			groupCounter++
			packGroups = append(packGroups, currPackGroup)
			currPackGroup.SetBadge()
			// badgeVal := ItemValues[currPackGroup.Badge]
			// fmt.Printf("Group %d badge %s value %v\n", groupCounter, currPackGroup.Badge, badgeVal)
			currPackGroup = &PackGroup{}
		}
		i++
	}

	return packGroups
}

func (pg *PackGroup) SetBadge() {
	packMaps := make([]map[string]int, 0)
	for _, pack := range pg.Packs {
		packMaps = append(packMaps, pack.CombinedCompartment.itemMap)
	}

	reqMatchCount := len(packMaps)

	// root loop
	pm0 := packMaps[0]
	for item := range pm0 {
		matchCount := 0
		// each additional pack
		for _, pack := range packMaps[0:] {
			// If item matches, increment count
			if _, ok := pack[item]; ok {
				matchCount++
			} else {
				// Reset matchCount if any item doesn't match
				matchCount = 0
			}
			if matchCount == reqMatchCount {
				pg.Badge = item
				return
			}
		}
	}
}

func NewCompartment(items []string) *Compartment {
	compMap := make(map[string]int)

	for _, c := range items {
		compMap[c] += 1
	}

	return &Compartment{itemMap: compMap}
}

func day3PreBuild(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	itemList := make([]string, 0)

	for s.Scan() {
		itemList = append(itemList, s.Text())
	}

	return itemList
}

func partA(itemList []string) int {
	prioritySum := 0

	for _, i := range itemList {
		newPack := NewPackFromString(i)
		prioritySum += ItemValues[newPack.SharedItem]
	}
	return prioritySum
}

func partB(itemList []string) int {
	packs := make([]*Pack, 0)

	for _, i := range itemList {
		newPack := NewPackFromString(i)
		packs = append(packs, newPack)
	}

	badgeSum := 0
	fullGroup := PackGroupsFromList(3, packs)

	for _, group := range fullGroup {
		badgeSum += ItemValues[group.Badge]
	}

	return badgeSum
}
