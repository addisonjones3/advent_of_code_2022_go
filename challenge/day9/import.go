package day9

import (
	"fmt"

	"github.com/spf13/cobra"
)

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "9",
		Short: "Problems for Day9",
	}

	day.AddCommand(aCommand())
	day.AddCommand(bCommand())
	root.AddCommand(day)
}

func aCommand() *cobra.Command {
	path := "challenge/day9/input.txt"
	bridgeInput := day9PreBuild(path)
	return &cobra.Command{
		Use:   "a",
		Short: "Day 9, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(bridgeInput))
		},
	}
}

func bCommand() *cobra.Command {
	path := "challenge/day9/input.txt"
	bridgeInput := day9PreBuild(path)
	return &cobra.Command{
		Use:   "b",
		Short: "Day 9, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(bridgeInput))
		},
	}
}
