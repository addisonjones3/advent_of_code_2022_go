package day3

import (
	"fmt"

	"github.com/spf13/cobra"
)

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "3",
		Short: "Problems for Day3",
	}

	day.AddCommand(aCommand())
	day.AddCommand(bCommand())
	root.AddCommand(day)
}

func aCommand() *cobra.Command {
	path := "challenge/day3/input.txt"
	itemList := day3PreBuild(path)
	return &cobra.Command{
		Use:   "a",
		Short: "Day 3, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(itemList))
		},
	}
}

func bCommand() *cobra.Command {
	path := "challenge/day3/input.txt"
	itemList := day3PreBuild(path)
	return &cobra.Command{
		Use:   "b",
		Short: "Day 3, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(itemList))
		},
	}
}
