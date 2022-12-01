package day1

import (
	"fmt"

	"github.com/spf13/cobra"
)

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "1",
		Short: "Problems for Day1",
	}

	day.AddCommand(aCommand())
	day.AddCommand(bCommand())
	// day.AddCommand(aSampleCommand())
	// day.AddCommand(bSampleCommand())
	root.AddCommand(day)
}

func aCommand() *cobra.Command {
	path := "challenge/day1/input.txt"
	packMap := preBuildDay1(path)

	return &cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(packMap))
		},
	}
}

// func aSampleCommand() *cobra.Command {
// 	path := "challenge/day1/sample_input.txt"
// 	packMap := preBuildDay1(path)

// 	return &cobra.Command{
// 		Use:   "asample",
// 		Short: "Day 1, Problem A",
// 		Run: func(_ *cobra.Command, _ []string) {
// 			fmt.Printf("Answer: %d\n", partA(packMap))
// 		},
// 	}
// }

func bCommand() *cobra.Command {
	path := "challenge/day1/input.txt"
	packMap := preBuildDay1(path)
	return &cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(packMap))
		},
	}
}

// func bSampleCommand() *cobra.Command {
// 	path := "challenge/day1/sample_input.txt"
// 	packMap := preBuildDay1(path)

// 	return &cobra.Command{
// 		Use:   "sampleB",
// 		Short: "Day 1, Problem B",
// 		Run: func(_ *cobra.Command, _ []string) {
// 			fmt.Printf("Answer: %d\n", partB(packMap))
// 		},
// 	}
// }
