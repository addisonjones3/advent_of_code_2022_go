package day10

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type RegisterType string

const (
	X RegisterType = "X"
)

type CircuitCommandType string

const (
	addx CircuitCommandType = "addx"
	noop CircuitCommandType = "noop"
)

type CircuitCommand struct {
	CircuitCommandType
	Val int
}

func CircuitCommandFromString(s string) (*CircuitCommand, error) {
	ccFields := strings.Fields(s)
	switch ccFields[0] {
	case string(addx):
		ccVal, _ := strconv.Atoi(ccFields[1])
		return &CircuitCommand{CircuitCommandType: addx, Val: ccVal}, nil
	case string(noop):
		return &CircuitCommand{CircuitCommandType: noop}, nil
	}
	return nil, errors.New("invalid cc string")
}

type Register struct {
	Name        string
	CurrentVal  int
	CycleValMap map[int]*Cycle
	RegisterType
}

func (r *Register) SignalStrength(cycle int) int {
	return r.CycleValMap[cycle].StartVal * cycle
}

func (r *Register) PrintSignalStrength(cycle int) {
	cycleVal := r.SignalStrength(cycle)
	fmt.Println(cycle, r.CycleValMap[cycle], cycleVal)
}

func NewRegister(name string, regType RegisterType) (*Register, error) {
	switch regType {
	case X:
		// cycle1 := &Cycle{StartVal: 1, EndVal: 1}
		return &Register{
			Name:         name,
			CurrentVal:   1,
			CycleValMap:  make(map[int]*Cycle),
			RegisterType: regType,
		}, nil
	}
	return nil, errors.New("no valid registertype passed")
}

type RegisterMap map[RegisterType]*Register

type Cycle struct {
	StartVal int
	EndVal   int
}

type ClockCircuit struct {
	RegisterMap
	CycleNumber int
}

func (cc *ClockCircuit) GetReg(rtype RegisterType) *Register {
	return cc.RegisterMap[rtype]
}

func InitClockCircuit(registers []*Register) *ClockCircuit {
	cc := &ClockCircuit{
		RegisterMap: make(RegisterMap),
		CycleNumber: 0,
	}
	for _, reg := range registers {
		cc.AddRegister(reg.Name, reg.RegisterType)
	}
	return cc
}

func (cc *ClockCircuit) AddRegister(name string, regType RegisterType) {
	newReg, err := NewRegister(name, regType)
	if err != nil {
		panic("Not a valid regtype")
	}
	cc.RegisterMap[newReg.RegisterType] = newReg
}

func (cc *ClockCircuit) HandleCommand(c *CircuitCommand) {
	switch c.CircuitCommandType {
	case addx:
		reg := cc.RegisterMap[X]
		if _, ok := reg.CycleValMap[cc.CycleNumber]; !ok {
			reg.CycleValMap[cc.CycleNumber] = &Cycle{StartVal: 1, EndVal: 1}
		}
		currCycle := reg.CycleValMap[cc.CycleNumber]

		cc.CycleNumber++
		reg.CycleValMap[cc.CycleNumber] = &Cycle{
			StartVal: currCycle.EndVal,
			EndVal:   currCycle.EndVal,
		}
		cc.CycleNumber++
		reg.CycleValMap[cc.CycleNumber] = &Cycle{
			StartVal: currCycle.EndVal,
			EndVal:   currCycle.EndVal + c.Val,
		}
	case noop:
		reg := cc.RegisterMap[X]
		currCycle := reg.CycleValMap[cc.CycleNumber]
		cc.CycleNumber++
		reg.CycleValMap[cc.CycleNumber] = &Cycle{
			StartVal: currCycle.EndVal,
			EndVal:   currCycle.EndVal,
		}
	}
}

func (cc *ClockCircuit) RegTypeSigStrength(regType RegisterType, cycles []int) int {
	totalSigStrength := 0

	reg := cc.RegisterMap[regType]
	for _, cycle := range cycles {
		cycleStrength := reg.SignalStrength(cycle)
		fmt.Printf("Cycle %d strength %d\n", cycle, cycleStrength)
		totalSigStrength += cycleStrength
	}

	fmt.Println(len(reg.CycleValMap) / 40)

	return totalSigStrength
}

func (cc ClockCircuit) DrawCRT(regType RegisterType, rowWidth int) string {
	onPxVal := "#"
	offPxVal := "."
	reg := cc.RegisterMap[regType]
	// pixelRows := make([]string, (len(reg.CycleValMap)/rowWidth)+1)
	rowBuilder := strings.Builder{}
	for i := 0; i < len(reg.CycleValMap)-1; i++ {
		spriteMiddle := reg.CycleValMap[i+1].StartVal
		rowLoc := i - (((i) / rowWidth) * rowWidth)
		// fmt.Println(rowLoc)
		// fmt.Println(i, "Sprite middle at", spriteMiddle)
		spriteLocs := map[int]string{
			spriteMiddle - 1: onPxVal,
			spriteMiddle:     onPxVal,
			spriteMiddle + 1: onPxVal,
		}
		if _, ok := spriteLocs[rowLoc]; ok {
			rowBuilder.WriteString(string(onPxVal))
		} else {
			rowBuilder.WriteString(string(offPxVal))
		}
		// fmt.Println(rowBuilder.String())
		if rowLoc == rowWidth-1 {
			rowBuilder.WriteString("\n")
			// pixelRows = append(pixelRows, rowBuilder.String())
			// fmt.Println(pixelRows)
			// fmt.Println("Reset")
			// rowBuilder.Reset()
			// rowBuilder = strings.Builder{}
		}
	}
	// return strings.Join(pixelRows, "\n")
	return rowBuilder.String()
}
