package d15

import (
	"strconv"
	"strings"
)

// ParseInput ..
func ParseInput(f []byte) ([]int, error) {
	split := strings.Split(string(f), ",")
	out := make([]int, len(split))
	for i, v := range split {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		out[i] = num
	}
	return out, nil
}

func playTheGame(start []int, stop int) int {

	type record struct{ Count, Ind, IndPrev int }

	register := make(map[int]record)
	memory := make([]int, stop)

	for i := 0; i < stop; i++ {
		curr := 0
		if i < len(start) {
			memory[i] = start[i]
			curr = start[i]
		} else if register[memory[i-1]].Count == 1 {
			memory[i] = 0
			curr = 0
		} else {
			memory[i] = register[memory[i-1]].Ind - register[memory[i-1]].IndPrev
			curr = memory[i]
		}

		if val, ok := register[curr]; ok {
			register[curr] = record{
				Count:   val.Count + 1,
				Ind:     i,
				IndPrev: val.Ind,
			}
		} else {
			register[curr] = record{
				Count:   1,
				Ind:     i,
				IndPrev: 0,
			}
		}
	}

	return memory[stop-1]
}

// SolvePart1 ..
func SolvePart1(start []int) int {
	return playTheGame(start, 2020)
}

// SolvePart2 ..
func SolvePart2(start []int) int {
	return playTheGame(start, 30000000)
}
