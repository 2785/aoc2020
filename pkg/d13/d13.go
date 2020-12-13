package d13

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/2785/aoc2020/pkg/numbers"
)

// ParseInput ..
func ParseInput(f []byte) ([]int, int, error) {
	split := strings.Split(string(f), "\n")
	if len(split) != 2 {
		return nil, 0, errors.New("unexpected number of lines")
	}
	r1, err := strconv.Atoi(split[0])
	if err != nil {
		return nil, 0, err
	}

	split = strings.Split(split[1], ",")
	out := make([]int, len(split))
	for i, v := range split {
		if v == "x" {
			out[i] = 0
		} else {
			num, err := strconv.Atoi(v)
			if err != nil {
				return nil, 0, err
			}
			out[i] = num
		}
	}

	return out, r1, nil
}

// SolvePart1 ..
func SolvePart1(schedule []int, departTime int) int {
	loc := make([]int, len(schedule))
	copy(loc, schedule)
	sort.Ints(loc)
	hiTime := loc[len(loc)-1]
	bus := loc[len(loc)-1]

	for _, v := range loc {
		if v == 0 {
			continue
		}
		waitTime := v - (departTime % v)
		if waitTime < hiTime {
			hiTime = waitTime
			bus = v
		}
	}

	return hiTime * bus
}

// SolvePart2 ..
func SolvePart2(schedule []int) (int, error) {
	done := false
	currMod := 0
	currInc := schedule[0]
	currSch := []int{schedule[0]}
	currSet := []int{schedule[0]}
	for !done {
		for i := len(currSch); i < len(schedule); i++ {
			currSch = append(currSch, schedule[i])
			if schedule[i] != 0 {
				currSet = append(currSet, schedule[i])
				break
			}
		}

		if len(currSch) == len(schedule) {
			done = true
		}

		currT := currMod

		newMod, err := solvePartSchedule(currSch, func() int {
			old := currT
			currT = currT + currInc
			return old
		})

		if err != nil {
			return 0, err
		}

		currMod = newMod
		currInc = numbers.LCM(1, 1, currSet...)
	}

	return currMod, nil
}

func solvePartSchedule(schedule []int, next func() int) (int, error) {
	good := func(num int) bool {
		for i, v := range schedule {
			if v != 0 && (num+i)%v != 0 {
				return false
			}
		}
		return true
	}

	for i := 1; i < 1e5; i++ {
		t := next()
		if good(t) {
			return t, nil
		}
	}

	return 0, errors.New("too many iterations")
}
