package day4

import (
	"fmt"

	"github.com/spf13/cobra"
)

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "4",
		Short: "Problems for Day3",
	}

	day.AddCommand(aCommand())
	day.AddCommand(bCommand())
	root.AddCommand(day)
}

func aCommand() *cobra.Command {
	path := "challenge/day4/input.txt"
	pairList := day4PreBuild(path)
	return &cobra.Command{
		Use:   "a",
		Short: "Day 4, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(pairList))
		},
	}
}

func bCommand() *cobra.Command {
	path := "challenge/day4/input.txt"
	pairList := day4PreBuild(path)
	return &cobra.Command{
		Use:   "b",
		Short: "Day 3, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(pairList))
		},
	}
}
