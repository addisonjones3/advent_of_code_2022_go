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
	// fmt.Println("len commands", len(cCommands))

	for i := 0; i < len(cCommands); i++ {
		// cc := cCommands[i]
		circ.HandleCommand(cCommands[i])
		// xr := circ.GetReg(X)
		// if len(xr.CycleValMap) >= 218 {
		// 	fmt.Println("command", i)
		// 	fmt.Println(cc.CircuitCommandType, cc.Val)
		// 	fmt.Println(xr.CycleValMap[(len(xr.CycleValMap))])
		// }
	}

	xr := circ.GetReg(X)

	cycles := []int{20, 60, 100, 140, 180, 220}
	xr.PrintSignalStrength(1)
	xr.PrintSignalStrength(20)
	xr.PrintSignalStrength(60)
	xr.PrintSignalStrength(100)
	xr.PrintSignalStrength(140)
	xr.PrintSignalStrength(180)
	xr.PrintSignalStrength(220)
	fmt.Println(circ.RegTypeSigStrength(X, cycles))
}
