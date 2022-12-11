package day10

import (
	"fmt"
	"strings"
	"testing"
)

var ShortTestLines = strings.Split(TestCommandsShort, "\n")
var LongTestLines = strings.Split(TestCommandsLong, "\n")

var GenCommands = func(s []string) []*CircuitCommand {
	cCommands := make([]*CircuitCommand, 0)
	for _, cc := range s {
		cComm, _ := CircuitCommandFromString(cc)
		cCommands = append(cCommands, cComm)
	}
	return cCommands
}

// func TestNewRegister(t *testing.T) {
// 	reg, err := NewRegister("first", X)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("reg: %v\n", reg)
// }

// func TestInitCC(t *testing.T) {
// 	xreg, _ := NewRegister("X reg", X)
// 	cc := InitClockCircuit([]*Register{xreg})
// 	fmt.Println(cc.RegisterMap[X][0])
// }

// func TestCCommandFromString(t *testing.T) {
// 	cCommands := make([]*CircuitCommand, 0)
// 	for _, cc := range ShortTestLines {
// 		cComm, _ := CircuitCommandFromString(cc)
// 		cCommands = append(cCommands, cComm)
// 	}

// 	for _, cc := range cCommands {
// 		fmt.Println(cc)
// 	}
// }

// func TestHandleCommand(t *testing.T) {
// 	xreg, _ := NewRegister("X Reg", X)
// 	circ := InitClockCircuit([]*Register{xreg})

// 	cCommands := GenCommands(ShortTestLines)
// 	for i := 0; i < len(cCommands); i++ {
// 		// cc := cCommands[i]
// 		circ.HandleCommand(cCommands[i])
// 	}
// 	fmt.Println(circ.RegisterMap[X])
// }

func TestCycleStrength(t *testing.T) {
	xreg, _ := NewRegister("X Reg", X)
	circ := InitClockCircuit([]*Register{xreg})

	cCommands := GenCommands(LongTestLines)

	for i := 0; i < len(cCommands); i++ {
		circ.HandleCommand(cCommands[i])
	}
	xr := circ.RegisterMap[X]
	for i := 1; i < 22; i++ {
		fmt.Println(i, xr.CycleValMap[i])
	}

	// cycles := []int{20, 60, 100, 140, 180, 220}
	// fmt.Println(circ.RegTypeSigStrength(X, cycles))
}

func TestDrawCRT(t *testing.T) {
	xreg, _ := NewRegister("X Reg", X)
	circ := InitClockCircuit([]*Register{xreg})

	cCommands := GenCommands(LongTestLines)

	for i := 0; i < len(cCommands); i++ {
		circ.HandleCommand(cCommands[i])
	}
	fmt.Println(circ.DrawCRT(X, 40))
}
