package day8

import (
	"fmt"

	"github.com/spf13/cobra"
)

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "8",
		Short: "Problems for Day8",
	}

	day.AddCommand(aCommand())
	day.AddCommand(bCommand())
	root.AddCommand(day)
}

func aCommand() *cobra.Command {
	path := "challenge/day8/input.txt"
	treeInput := day8PreBuild(path)
	return &cobra.Command{
		Use:   "a",
		Short: "Day 8, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(treeInput))
		},
	}
}

func bCommand() *cobra.Command {
	path := "challenge/day8/input.txt"
	treeInput := day8PreBuild(path)
	return &cobra.Command{
		Use:   "b",
		Short: "Day 8, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(treeInput))
		},
	}
}
