package day5

import (
	"fmt"

	"github.com/spf13/cobra"
)

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "5",
		Short: "Problems for Day5",
	}

	day.AddCommand(aCommand())
	day.AddCommand(bCommand())
	root.AddCommand(day)
}

func aCommand() *cobra.Command {
	path := "challenge/day5/input.txt"
	crateYardInput, moveInput := day5PreBuild(path)
	return &cobra.Command{
		Use:   "a",
		Short: "Day 5, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %s\n", partA(crateYardInput, moveInput))
		},
	}
}

func bCommand() *cobra.Command {
	path := "challenge/day5/input.txt"
	crateYardInput, moveInput := day5PreBuild(path)
	return &cobra.Command{
		Use:   "b",
		Short: "Day 5, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %s\n", partB(crateYardInput, moveInput))
		},
	}
}
