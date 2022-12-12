package day11

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ParseIdFromString(s string) int {
	re := regexp.MustCompile(`([0-9])+`)
	idStr, _ := strconv.Atoi(re.FindString(s))
	return idStr
}

func ParseItemsFromSring(s string) ItemSlice {
	re := regexp.MustCompile(`([0-9])+`)
	valBytes := re.FindAll([]byte(s), -1)
	items := make(ItemSlice, 0)
	for _, nByte := range valBytes {
		worryVal, _ := strconv.Atoi(string(nByte))
		items = append(items, NewItem(worryVal))
	}
	return items
}

func ParseOperationFromString(s string) *Operation {
	strFields := strings.Fields(s)
	oprVal, notInt := strconv.Atoi(strFields[len(strFields)-1])
	selfOp := false
	if notInt != nil {
		selfOp = oprVal == 0
	}
	var op Operator
	switch strFields[len(strFields)-2] {
	case string(mult):
		op = mult
	case string(add):
		op = add
	case string(sub):
		op = sub
	}
	return &Operation{
		Operator: op,
		SelfOp:   selfOp,
		Value:    oprVal,
	}
}

func ParseTestOperationFromString(toLines string) *TestOperation {
	opFields := strings.Fields(toLines)
	toVal, _ := strconv.Atoi(opFields[len(opFields)-1])
	return &TestOperation{
		Val: toVal,
	}
}

func ParseTargetMonkeyIds(tLines []string) (int, int) {
	trueTargetId, _ := strconv.Atoi(string(tLines[0][len(tLines[0])-1]))
	falseTargetId, _ := strconv.Atoi(string(tLines[1][len(tLines[1])-1]))
	return trueTargetId, falseTargetId
}

func ParseMonkeyGroupFromString(s string) *MonkeyGroup {
	scanner := bufio.NewScanner(strings.NewReader(s))
	writer := strings.Builder{}
	mg := &MonkeyGroup{MonkeyMap: make(map[int]*Monkey)}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			m := NewMonkeyFromString(writer.String())
			mg.MonkeyMap[m.id] = m
			writer.Reset()
		} else {
			writer.WriteString(line)
			writer.WriteString("\n")
		}
	}
	m := NewMonkeyFromString(writer.String())
	mg.MonkeyMap[m.id] = m

	divisor := 1
	for _, m := range mg.MonkeyMap {
		fmt.Println(m.TestOperation.Val)
		divisor *= m.TestOperation.Val
		fmt.Println(divisor)
	}
	mg.ModDivisor = divisor
	return mg
}
