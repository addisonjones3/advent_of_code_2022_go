/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/addisonjones3/aoc22/challenge/day1"
	"github.com/addisonjones3/aoc22/challenge/day10"
	"github.com/addisonjones3/aoc22/challenge/day11"
	"github.com/addisonjones3/aoc22/challenge/day2"
	"github.com/addisonjones3/aoc22/challenge/day3"
	"github.com/addisonjones3/aoc22/challenge/day4"
	"github.com/addisonjones3/aoc22/challenge/day5"
	"github.com/addisonjones3/aoc22/challenge/day6"
	"github.com/addisonjones3/aoc22/challenge/day7"
	"github.com/addisonjones3/aoc22/challenge/day8"
	"github.com/addisonjones3/aoc22/challenge/day9"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc22",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

func addDays(cmd *cobra.Command) {
	day1.AddCommandsTo(rootCmd)
	day2.AddCommandsTo(rootCmd)
	day3.AddCommandsTo(rootCmd)
	day4.AddCommandsTo(rootCmd)
	day5.AddCommandsTo(rootCmd)
	day6.AddCommandsTo(rootCmd)
	day7.AddCommandsTo(rootCmd)
	day8.AddCommandsTo(rootCmd)
	day9.AddCommandsTo(rootCmd)
	day10.AddCommandsTo(rootCmd)
	day11.AddCommandsTo(rootCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	addDays(rootCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aoc22.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
