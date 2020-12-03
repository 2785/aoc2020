package d3

import (
	"strings"
	"sync"
	"sync/atomic"
)

// IsTree ..
type IsTree func(int) bool

// ParseInput ..
func ParseInput(f []byte) []IsTree {
	lines := strings.Split(string(f), "\n")
	out := make([]IsTree, len(lines))

	for i, v := range lines {
		line := v
		out[i] = func(col int) bool {
			return line[col%len(line)] == '#'
		}
	}

	return out
}

// SolvePart1 .
func SolvePart1(m []IsTree) int {
	var counter int64
	wg := &sync.WaitGroup{}
	wg.Add(len(m))
	for rowNum := range m {
		go func(r int) {
			if m[r](3 * r) {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}(rowNum)
	}

	wg.Wait()

	return int(counter)
}

// SolvePart2 ..
func SolvePart2(m []IsTree) int {
	type direction struct {
		Right, Down int
	}

	makeInt64 := func() *int64 {
		var in int64
		return &in
	}

	directions := map[*direction]*int64{
		{1, 1}: makeInt64(),
		{3, 1}: makeInt64(),
		{5, 1}: makeInt64(),
		{7, 1}: makeInt64(),
		{1, 2}: makeInt64(),
	}

	wg := sync.WaitGroup{}
	wg.Add(len(m) * len(directions))

	for dir := range directions {
		for rowNum := range m {
			go func(r int, d *direction) {
				defer wg.Done()
				if r%d.Down != 0 {
					return
				}
				if m[r](d.Right * (r / d.Down)) {
					atomic.AddInt64(directions[d], 1)
				}
			}(rowNum, dir)
		}
	}

	wg.Wait()

	prod := 1

	for _, v := range directions {
		prod *= int(*v)
	}

	return prod
}
