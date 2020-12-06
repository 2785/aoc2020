package d6

import (
	"strings"
	"sync"
	"sync/atomic"
)

// Group .
type Group struct {
	Count     int
	Questions map[rune]int
}

// DecodeOne .
func DecodeOne(one string) Group {
	oneLine := strings.ReplaceAll(one, "\n", "")
	grp := Group{
		Count:     strings.Count(one, "\n") + 1,
		Questions: make(map[rune]int),
	}
	for _, v := range oneLine {
		if _, ok := grp.Questions[v]; !ok {
			grp.Questions[v] = 0
		}
		grp.Questions[v]++
	}

	return grp
}

// ParseInput .
func ParseInput(f []byte) []Group {
	split := strings.Split(string(f), "\n\n")
	grps := make([]Group, len(split))
	wg := &sync.WaitGroup{}
	wg.Add(len(grps))

	for i, v := range split {
		go func(ind int, val string) {
			grps[ind] = DecodeOne(val)
			wg.Done()
		}(i, v)
	}

	wg.Wait()

	return grps
}

// SolvePart1 .
func SolvePart1(grps []Group) int {
	var sum int
	for _, grp := range grps {
		sum += len(grp.Questions)
	}
	return sum
}

// SolvePart2 .
func SolvePart2(grps []Group) int {
	var sum int64

	wg := &sync.WaitGroup{}
	wg.Add(len(grps))

	for _, grp := range grps {
		go func(g Group) {
			defer wg.Done()
			for _, v := range g.Questions {
				if v == g.Count {
					atomic.AddInt64(&sum, 1)
				}
			}
		}(grp)
	}

	wg.Wait()

	return int(sum)
}
