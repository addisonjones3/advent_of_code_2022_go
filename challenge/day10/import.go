package day10

import (
	"fmt"

	"github.com/spf13/cobra"
)

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "10",
		Short: "Problems for Day10",
	}

	day.AddCommand(aCommand())
	day.AddCommand(bCommand())
	root.AddCommand(day)
}

func aCommand() *cobra.Command {
	path := "challenge/day10/input.txt"
	commandLines := day10PreBuild(path)
	return &cobra.Command{
		Use:   "a",
		Short: "Day 10, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(commandLines))
		},
	}
}

func bCommand() *cobra.Command {
	path := "challenge/day10/input.txt"
	commandLines := day10PreBuild(path)
	return &cobra.Command{
		Use:   "b",
		Short: "Day 10, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: \n%s\n", partB(commandLines))
		},
	}
}
