package d16

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/2785/aoc2020/pkg/regutil"
	"github.com/adam-hanna/arrayOperations"
)

// ParseInput ..
func ParseInput(f []byte) (rules map[string]func(int) bool, yourTicket []int, nearbyTickets [][]int, e error) {
	split1 := strings.Split(string(f), "\n\nyour ticket:\n")
	if len(split1) != 2 {
		e = fmt.Errorf("unexpected segments")
		return
	}

	ruleStr := strings.TrimSpace(split1[0])

	split2 := strings.Split(split1[1], "\n\nnearby tickets:\n")

	if len(split2) != 2 {
		e = fmt.Errorf("unexpected segments")
		return
	}

	yourTicketStr := strings.TrimSpace(split2[0])
	nearbyTicketStr := strings.TrimSpace(split2[1])

	// Process ruleset
	ruleRe := regexp.MustCompile(`(?P<field>[a-z ]+): (?P<x1>\d+)-(?P<x2>\d+) or (?P<x3>\d+)-(?P<x4>\d+)`)

	rules, e = func() (map[string]func(int) bool, error) {
		split := strings.Split(ruleStr, "\n")
		out := make(map[string]func(int) bool)
		for _, v := range split {
			match, err := regutil.MustCaptureNamedGroup(ruleRe, []string{"field", "x1", "x2", "x3", "x4"}, v)
			if err != nil {
				return nil, err
			}
			fieldName := match["field"]
			x1, err := strconv.Atoi(match["x1"])
			if err != nil {
				return nil, err
			}
			x2, err := strconv.Atoi(match["x2"])
			if err != nil {
				return nil, err
			}

			x3, err := strconv.Atoi(match["x3"])
			if err != nil {
				return nil, err
			}

			x4, err := strconv.Atoi(match["x4"])
			if err != nil {
				return nil, err
			}

			out[fieldName] = func(i int) bool {
				return (i >= x1 && i <= x2) || (i >= x3 && i <= x4)
			}
		}
		return out, nil
	}()

	if e != nil {
		return
	}

	// Process your ticket
	yourTicket, e = func() ([]int, error) {
		split := strings.Split(yourTicketStr, ",")
		out := make([]int, len(split))
		for i, v := range split {
			num, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			out[i] = num
		}
		return out, nil
	}()

	if e != nil {
		return
	}

	// Now process nearby tickets
	nearbyTickets, e = func() ([][]int, error) {
		splitLine := strings.Split(nearbyTicketStr, "\n")
		out := make([][]int, len(splitLine))
		for row, line := range splitLine {
			splitNumber := strings.Split(line, ",")
			out[row] = make([]int, len(splitNumber))
			for i, numStr := range splitNumber {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					return nil, err
				}
				out[row][i] = num
			}
		}
		return out, nil
	}()

	if e != nil {
		return
	}

	return
}

// SolvePart1 ..
func SolvePart1(nearby [][]int, rules map[string]func(int) bool) int {

	var sum int64

	wg := &sync.WaitGroup{}

	for _, ticket := range nearby {
		for _, num := range ticket {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				if !goodForAtLeast1(rules, i) {
					atomic.AddInt64(&sum, int64(i))
				}
			}(num)
		}
	}

	wg.Wait()

	return int(sum)
}

// SolvePart2 ..
func SolvePart2(yours []int, nearby [][]int, rules map[string]func(int) bool) (int, error) {
	round := [][][]string{}
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(len(nearby))
	for _, ticket := range nearby {
		go func(t []int) {
			defer wg.Done()
			validFields := make([][]string, len(t))
			for i, num := range t {
				possible := []string{}
				for name, good := range rules {
					if good(num) {
						possible = append(possible, name)
					}
				}
				if len(possible) == 0 {
					// to short circuit completely invalid tickets
					return
				}
				validFields[i] = possible
			}
			mu.Lock()
			round = append(round, validFields)
			mu.Unlock()
		}(ticket)
	}

	wg.Wait()

	entries := len(rules)
	for _, v := range round {
		if len(v) != entries {
			return 0, errors.New("length mismatch")
		}
	}

	intersect := make(map[int][]string, entries)

	for i := 0; i < entries; i++ {
		pos := make([][]string, len(round))
		for ind, ticket := range round {
			pos[ind] = ticket[i]
		}

		intersection := arrayOperations.IntersectString(pos...)

		intersect[i] = intersection
	}

	sorted := make(map[string]int)

	for len(intersect) > 0 {
		newWorkingGroup := make(map[int][]string)
		for ind, possibilities := range intersect {
			if len(possibilities) == 1 {
				sorted[possibilities[0]] = ind
			} else {
				remainingPossibilities := make([]string, 0)
				for _, v := range possibilities {
					if _, ok := sorted[v]; !ok {
						remainingPossibilities = append(remainingPossibilities, v)
					}
				}
				newWorkingGroup[ind] = remainingPossibilities
			}
		}
		intersect = newWorkingGroup
	}

	prod := 1

	fieldsOfInterest := []string{
		"departure location",
		"departure station",
		"departure platform",
		"departure track",
		"departure date",
		"departure time",
	}

	for _, v := range fieldsOfInterest {
		prod *= yours[sorted[v]]
	}

	return prod, nil
}

func goodForAtLeast1(rules map[string]func(int) bool, i int) bool {
	for _, good := range rules {
		if good(i) {
			return true
		}
	}
	return false
}
