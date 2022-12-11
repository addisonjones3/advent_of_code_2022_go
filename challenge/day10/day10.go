package day10

import (
	"bufio"
	"os"
)

func day10PreBuild(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	commandLines := make([]string, 0)
	for s.Scan() {
		commandLines = append(commandLines, s.Text())
	}
	return commandLines
}

func partA(commandLines []string) int {
	xreg, _ := NewRegister("X Reg", X)
	circ := InitClockCircuit([]*Register{xreg})

	for _, cLine := range commandLines {
		cc, _ := CircuitCommandFromString(cLine)
		circ.HandleCommand(cc)
	}

	cycles := []int{20, 60, 100, 140, 180, 220}
	return circ.RegTypeSigStrength(X, cycles)

}

func partB(commandLines []string) string {
	xreg, _ := NewRegister("X Reg", X)
	circ := InitClockCircuit([]*Register{xreg})

	for _, cLine := range commandLines {
		cc, _ := CircuitCommandFromString(cLine)
		circ.HandleCommand(cc)
	}
	return circ.DrawCRT(X, 40)
}
