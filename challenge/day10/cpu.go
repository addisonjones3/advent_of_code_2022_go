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
	CycleValMap map[int]int
	RegisterType
}

func (r *Register) SignalStrength(cycle int) int {
	if cycle == 1 {
		return r.CycleValMap[cycle] * cycle
	} else {
		return r.CycleValMap[cycle-1] * cycle
	}
}

func (r *Register) PrintSignalStrength(cycle int) {
	cycleVal := r.SignalStrength(cycle)
	fmt.Println(cycle, r.CycleValMap[cycle], cycleVal)
}

func NewRegister(name string, regType RegisterType) (*Register, error) {
	switch regType {
	case X:
		return &Register{
			Name: name, CurrentVal: 1, CycleValMap: map[int]int{1: 1}, RegisterType: regType,
		}, nil
	}
	return nil, errors.New("no valid registertype passed")
}

type RegisterMap map[RegisterType]*Register

type ClockCircuit struct {
	RegisterMap
	CurrCycle int
}

func (cc *ClockCircuit) GetReg(rtype RegisterType) *Register {
	return cc.RegisterMap[rtype]
}

func InitClockCircuit(registers []*Register) *ClockCircuit {
	cc := &ClockCircuit{
		RegisterMap: make(RegisterMap),
		CurrCycle:   1,
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
		reg.CycleValMap[cc.CurrCycle] = reg.CurrentVal
		reg.CurrentVal += c.Val
		reg.CycleValMap[cc.CurrCycle+1] += reg.CurrentVal

		cc.CurrCycle += 2
	case noop:
		for rType := range cc.RegisterMap {
			reg := cc.RegisterMap[rType]
			reg.CycleValMap[cc.CurrCycle] = reg.CurrentVal
		}
		cc.CurrCycle += 1
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
	// reg := cc.RegisterMap[regType]
	// pixelRows := make([]string, len(reg.CycleValMap)/rowWidth)

	return "done"
}
