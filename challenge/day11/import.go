package day11

import (
	"fmt"

	"github.com/spf13/cobra"
)

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "11",
		Short: "Problems for Day11",
	}

	day.AddCommand(aCommand())
	day.AddCommand(bCommand())
	root.AddCommand(day)
}

func aCommand() *cobra.Command {
	// path := "challenge/day11/input.txt"
	// monkeyBlock := day11PreBuild(path)
	return &cobra.Command{
		Use:   "a",
		Short: "Day 11, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA())
		},
	}
}

func bCommand() *cobra.Command {
	// path := "challenge/day11/input.txt"
	// commandLines := day11PreBuild(path)
	return &cobra.Command{
		Use:   "b",
		Short: "Day 11, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB())
		},
	}
}
