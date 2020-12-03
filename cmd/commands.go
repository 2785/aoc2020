package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/2785/aoc2020/pkg/d1"
	"github.com/2785/aoc2020/pkg/d2"
	"github.com/2785/aoc2020/pkg/d3"
	"github.com/spf13/cobra"
)

var inputFile string

var d1Cmd = &cobra.Command{
	Use:   "d1",
	Short: "Run solution for day 1",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("running day 1 puzzle with input file %s\n", inputFile)

		f, err := ioutil.ReadFile(filepath.Clean(inputFile))
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		parsed, err := d1.ParseFile(f)

		if err != nil {
			return fmt.Errorf("cannot parse file %s: %w", inputFile, err)
		}

		part1 := d1.SolvePart1(parsed)

		fmt.Printf("Part 1 solution: %v\n", part1)

		part2 := d1.SolvePart2(parsed)

		fmt.Printf("Part 2 solution: %v\n", part2)
		return nil
	},
}

var d2Cmd = &cobra.Command{
	Use:   "d2",
	Short: "Run solution for day 2",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("running day 2 puzzle with input file %s\n", inputFile)

		f, err := ioutil.ReadFile(filepath.Clean(inputFile))
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		parsed, err := d2.ParseInput(f)

		if err != nil {
			return fmt.Errorf("cannot parse file %s: %w", inputFile, err)
		}

		part1 := d2.SolvePart1(parsed)

		fmt.Printf("Part 1 solution: %v\n", part1)

		part2 := d2.SolvePart2(parsed)

		fmt.Printf("Part 2 solution: %v\n", part2)
		return nil
	},
}

var d3Cmd = &cobra.Command{
	Use:   "d3",
	Short: "Run solution for day 3",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("running day 3 puzzle with input file %s\n", inputFile)

		f, err := ioutil.ReadFile(filepath.Clean(inputFile))
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		parsed := d3.ParseInput(f)

		part1 := d3.SolvePart1(parsed)

		fmt.Printf("Part 1 solution: %v\n", part1)

		part2 := d3.SolvePart2(parsed)

		fmt.Printf("Part 2 solution: %v\n", part2)
		return nil
	},
}

func init() {
	d1Cmd.Flags().StringVarP(&inputFile, "input", "i", "inputs/d1", "input file path")
	rootCmd.AddCommand(d1Cmd)

	d2Cmd.Flags().StringVarP(&inputFile, "input", "i", "inputs/d2", "input file path")
	rootCmd.AddCommand(d2Cmd)

	d3Cmd.Flags().StringVarP(&inputFile, "input", "i", "inputs/d3", "input file path")
	rootCmd.AddCommand(d3Cmd)
}
