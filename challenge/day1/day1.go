package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Pack struct {
	Meals []int
}

type PackMap map[int]Pack

func (p *Pack) totalCals() int {
	cals := 0
	for _, meal := range p.Meals {
		cals += meal
	}
	return cals
}

func sum(l []int) int {
	sum := 0
	for _, i := range l {
		sum += i
	}
	return sum
}

func preBuildDay1(path string) PackMap {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	packMap := make(PackMap)
	elfCounter := 0

	currPack := &Pack{}

	for s.Scan() {
		mealVal := s.Text()
		// fmt.Println(mealVal)

		if mealVal == "" {
			// Print cal counts to validate
			// currCalValue := currPack.totalCals()
			// fmt.Printf("Pack %d cal count is %d\n", elfCounter, currCalValue)
			packMap[elfCounter] = *currPack

			// Move counter, reinitialize blank pack
			elfCounter += 1
			currPack = &Pack{}
		} else {
			mealIntVal, _ := strconv.Atoi(mealVal)
			currPack.Meals = append(currPack.Meals, mealIntVal)
		}
	}

	packMap[elfCounter] = *currPack

	return packMap
}

func partA(pm PackMap) int {
	highestCalValue := 0
	for _, pack := range pm {
		packCals := pack.totalCals()
		if packCals > highestCalValue {
			highestCalValue = packCals
		}
	}
	return highestCalValue
}

func partB(pm PackMap) int {

	highestCalValue := 0
	secondCalValue := 0
	thirdCalValue := 0

	for _, pack := range pm {
		currCalValue := pack.totalCals()
		if currCalValue > highestCalValue {
			thirdCalValue = secondCalValue
			secondCalValue = highestCalValue
			highestCalValue = currCalValue
		} else if currCalValue > secondCalValue {
			thirdCalValue = secondCalValue
			secondCalValue = currCalValue
		} else if currCalValue > thirdCalValue {
			thirdCalValue = currCalValue
		}
	}
	fmt.Println("Part B")
	return sum([]int{highestCalValue, secondCalValue, thirdCalValue})
}
