package day11

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Item struct {
	id         string
	WorryValue int
}

type ItemSlice []*Item

func (is ItemSlice) ToString() string {
	itemVals := make([]string, 0)
	for _, i := range is {
		itemVals = append(itemVals, fmt.Sprintf("%d", i.WorryValue))
	}
	return strings.Join(itemVals, ", ")
}

func (is ItemSlice) PrintVals() {
	// itemVals := make([]int, 0)
	// for _, i := range is {
	// 	itemVals = append(itemVals, i.WorryValue)
	// }
	// fmt.Println(itemVals)
	fmt.Println(is.ToString())
}

func NewItem(worryVal int) *Item {
	id := uuid.NewString()
	return &Item{
		id:         id,
		WorryValue: worryVal,
	}
}
