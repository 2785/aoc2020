package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/2785/aoc2020/pkg/d1"
	"github.com/2785/aoc2020/pkg/d10"
	"github.com/2785/aoc2020/pkg/d11"
	"github.com/2785/aoc2020/pkg/d2"
	"github.com/2785/aoc2020/pkg/d3"
	"github.com/2785/aoc2020/pkg/d4"
	"github.com/2785/aoc2020/pkg/d5"
	"github.com/2785/aoc2020/pkg/d6"
	"github.com/2785/aoc2020/pkg/d7"
	"github.com/2785/aoc2020/pkg/d8"
	"github.com/2785/aoc2020/pkg/d9"
	"github.com/2785/aoc2020/pkg/input"
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

var d4Cmd = &cobra.Command{
	Use:   "d4",
	Short: "Run solution for day 4",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("running day 4 puzzle with input file %s\n", inputFile)

		f, err := ioutil.ReadFile(filepath.Clean(inputFile))
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		parsed, err := d4.ParseInput(f)

		if err != nil {
			return fmt.Errorf("cannot parse file %s: %w", inputFile, err)
		}

		part1 := d4.SolvePart1(parsed)

		fmt.Printf("Part 1 solution: %v\n", part1)

		part2 := d4.SolvePart2(parsed)

		fmt.Printf("Part 2 solution: %v\n", part2)
		return nil
	},
}

var d5Cmd = &cobra.Command{
	Use:   "d5",
	Short: "Run solution for day 5",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("running day 5 puzzle with input file %s\n", inputFile)

		f, err := ioutil.ReadFile(filepath.Clean(inputFile))
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		parsed, err := d5.ParseInput(f)

		if err != nil {
			return fmt.Errorf("cannot parse file %s: %w", inputFile, err)
		}

		part1 := d5.SolvePart1(parsed)

		fmt.Printf("Part 1 solution: %v\n", part1)

		part2 := d5.SolvePart2(parsed)

		fmt.Printf("Part 2 solution: %v\n", part2)
		return nil
	},
}

var d6Cmd = &cobra.Command{
	Use:   "d6",
	Short: "Run solution for day 6",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("running day 6 puzzle with input file %s\n", inputFile)

		f, err := ioutil.ReadFile(filepath.Clean(inputFile))
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		parsed := d6.ParseInput(f)

		part1 := d6.SolvePart1(parsed)

		fmt.Printf("Part 1 solution: %v\n", part1)

		part2 := d6.SolvePart2(parsed)

		fmt.Printf("Part 2 solution: %v\n", part2)
		return nil
	},
}

var d7Cmd = &cobra.Command{
	Use:   "d7",
	Short: "Run solution for day 7",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("running day 7 puzzle with input file %s\n", inputFile)

		f, err := ioutil.ReadFile(filepath.Clean(inputFile))
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		parsed, err := d7.ParseInput(f)

		if err != nil {
			return fmt.Errorf("cannot parse file %s: %w", inputFile, err)
		}

		part1 := d7.SolvePart1(parsed)

		fmt.Printf("Part 1 solution: %v\n", part1)

		part2 := d7.SolvePart2(parsed)

		fmt.Printf("Part 2 solution: %v\n", part2)
		return nil
	},
}

var d8Cmd = &cobra.Command{
	Use:   "d8",
	Short: "Run solution for day 8",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("running day 8 puzzle with input file %s\n", inputFile)

		f, err := ioutil.ReadFile(filepath.Clean(inputFile))
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		parsed, err := d8.ParseInput(f)

		if err != nil {
			return fmt.Errorf("cannot parse file %s: %w", inputFile, err)
		}

		part1 := d8.SolvePart1(parsed)

		fmt.Printf("Part 1 solution: %v\n", part1)

		part2, err := d8.SolvePart2(parsed)

		if err != nil {
			return fmt.Errorf("error solving part 2: %w", err)
		}

		fmt.Printf("Part 2 solution: %v\n", part2)
		return nil
	},
}

var d9Cmd = &cobra.Command{
	Use:   "d9",
	Short: "Run solution for day 9",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("running day 9 puzzle with input file %s\n", inputFile)

		f, err := ioutil.ReadFile(filepath.Clean(inputFile))
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		parsed, err := d9.ParseInput(f)

		if err != nil {
			return fmt.Errorf("cannot parse file %s: %w", inputFile, err)
		}

		part1, err := d9.SolvePart1(parsed, 25)

		if err != nil {
			return fmt.Errorf("error solving part 1: %w", err)
		}

		fmt.Printf("Part 1 solution: %v\n", part1)

		part2, err := d9.SolvePart2(parsed, part1)

		if err != nil {
			return fmt.Errorf("error solving part 2: %w", err)
		}

		fmt.Printf("Part 2 solution: %v\n", part2)
		return nil
	},
}

var d10Cmd = &cobra.Command{
	Use:   "d10",
	Short: "Run solution for day 10",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("running day 10 puzzle with input file %s\n", inputFile)

		f, err := ioutil.ReadFile(filepath.Clean(inputFile))
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		parsed := input.MustParseInt(f)

		part1 := d10.SolvePart1(parsed)

		fmt.Printf("Part 1 solution: %v\n", part1)

		part2 := d10.SolvePart2(parsed)

		fmt.Printf("Part 2 solution: %v\n", part2)
		return nil
	},
}

var d11Cmd = &cobra.Command{
	Use:   "d11",
	Short: "Run solution for day 11",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("running day 11 puzzle with input file %s\n", inputFile)

		f, err := ioutil.ReadFile(filepath.Clean(inputFile))
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}

		parsed, err := d11.ParseInput(f)

		if err != nil {
			return fmt.Errorf("cannot parse file %s: %w", inputFile, err)
		}

		part1, err := d11.SolvePart1(parsed)

		if err != nil {
			return fmt.Errorf("error solving part 1: %w", err)
		}

		fmt.Printf("Part 1 solution: %v\n", part1)

		part2, err := d11.SolvePart2(parsed)

		if err != nil {
			return fmt.Errorf("error solving part 2: %w", err)
		}

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

	d4Cmd.Flags().StringVarP(&inputFile, "input", "i", "inputs/d4", "input file path")
	rootCmd.AddCommand(d4Cmd)

	d5Cmd.Flags().StringVarP(&inputFile, "input", "i", "inputs/d5", "input file path")
	rootCmd.AddCommand(d5Cmd)

	d6Cmd.Flags().StringVarP(&inputFile, "input", "i", "inputs/d6", "input file path")
	rootCmd.AddCommand(d6Cmd)

	d7Cmd.Flags().StringVarP(&inputFile, "input", "i", "inputs/d7", "input file path")
	rootCmd.AddCommand(d7Cmd)

	d8Cmd.Flags().StringVarP(&inputFile, "input", "i", "inputs/d8", "input file path")
	rootCmd.AddCommand(d8Cmd)

	d9Cmd.Flags().StringVarP(&inputFile, "input", "i", "inputs/d9", "input file path")
	rootCmd.AddCommand(d9Cmd)

	d10Cmd.Flags().StringVarP(&inputFile, "input", "i", "inputs/d10", "input file path")
	rootCmd.AddCommand(d10Cmd)

	d11Cmd.Flags().StringVarP(&inputFile, "input", "i", "inputs/d11", "input file path")
	rootCmd.AddCommand(d11Cmd)
}
