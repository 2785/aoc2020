package d8

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

// Operation ..
type Operation int

const (
	nop Operation = iota
	acc
	jmp
)

// Instruction ..
type Instruction struct {
	Op  Operation
	Val int
}

// DecodeOne ..
func DecodeOne(s string) (*Instruction, error) {
	split := strings.Split(s, " ")
	if len(split) != 2 {
		return nil, fmt.Errorf("unrecognizable string: %s", s)
	}

	op, err := func() (Operation, error) {
		switch split[0] {
		case "nop":
			return nop, nil
		case "acc":
			return acc, nil
		case "jmp":
			return jmp, nil
		}
		return nop, fmt.Errorf("unrecognizable operation: %s", split[0])
	}()

	if err != nil {
		return nil, err
	}

	num, err := strconv.Atoi(split[1])
	if err != nil {
		return nil, fmt.Errorf("unrecognizable number: %s", split[1])
	}

	return &Instruction{op, num}, nil
}

// ParseInput ..
func ParseInput(f []byte) ([]*Instruction, error) {
	split := strings.Split(string(f), "\n")
	out := make([]*Instruction, len(split))
	errgrp := &errgroup.Group{}

	for ind, line := range split {
		errgrp.Go(func(i int, c string) func() error {
			return func() error {
				ins, err := DecodeOne(c)
				if err != nil {
					return err
				}
				out[i] = ins
				return nil
			}
		}(ind, line))
	}

	err := errgrp.Wait()
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SolvePart1 .
func SolvePart1(inst []*Instruction) int {
	reg := 0
	track := make(map[int]struct{})
	curser := 0

	for {
		if _, ok := track[curser]; ok {
			break
		}
		if curser >= len(inst) {
			break
		}
		track[curser] = struct{}{}
		switch inst[curser].Op {
		case acc:
			reg += inst[curser].Val
			curser++
		case nop:
			curser++
		case jmp:
			curser += inst[curser].Val
		}
	}

	return reg
}

// SolvePart2 .
func SolvePart2(inst []*Instruction) (int, error) {
	indicesToFlip := []int{}

	for i, v := range inst {
		if v.Op == nop || v.Op == jmp {
			indicesToFlip = append(indicesToFlip, i)
		}
	}

	successfulOutputs := []int{}
	mu := &sync.Mutex{}

	wg := &sync.WaitGroup{}
	wg.Add(len(indicesToFlip))

	for _, ind := range indicesToFlip {
		go func(indToMutate int) {
			defer wg.Done()

			copied := make([]*Instruction, len(inst))
			copy(copied, inst)

			switch copied[indToMutate].Op {
			case nop:
				copied[indToMutate] = &Instruction{jmp, copied[indToMutate].Val}
			case jmp:
				copied[indToMutate] = &Instruction{nop, copied[indToMutate].Val}
			}

			reg := 0
			track := make(map[int]struct{})
			curser := 0

			for {
				if _, ok := track[curser]; ok {
					return
				}
				if curser >= len(copied) {
					mu.Lock()
					successfulOutputs = append(successfulOutputs, reg)
					mu.Unlock()
					return
				}
				track[curser] = struct{}{}
				switch copied[curser].Op {
				case acc:
					reg += copied[curser].Val
					curser++
				case nop:
					curser++
				case jmp:
					curser += copied[curser].Val
				}
			}
		}(ind)
	}

	wg.Wait()

	if len(successfulOutputs) != 1 {
		return 0, fmt.Errorf("got %v ones that work", len(successfulOutputs))
	}

	return successfulOutputs[0], nil
}
