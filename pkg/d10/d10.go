package d10

import (
	"sort"
	"sync"
	"sync/atomic"
)

// SolvePart1 ..
func SolvePart1(nums []int) int {
	loc := make([]int, len(nums))
	copy(loc, nums)
	sort.Ints(loc)
	var diff1, diff3 int64
	wg := &sync.WaitGroup{}
	wg.Add(len(nums) - 1)

	for i, v := range loc[1:] {
		go func(ind, val int) {
			defer wg.Done()
			if val-loc[ind] == 3 {
				atomic.AddInt64(&diff3, 1)
			}
			if val-loc[ind] == 1 {
				atomic.AddInt64(&diff1, 1)
			}
		}(i, v)
	}

	switch loc[0] {
	case 1:
		diff1++
	case 3:
		diff3++
	}

	wg.Wait()
	return int(diff1 * (diff3 + 1))
}

// SolvePart2 ..
func SolvePart2(nums []int) int {
	loc := make([]int, len(nums))
	copy(loc, nums)
	sort.Ints(loc)
	diffs := make([]int, len(nums))
	diffs[0] = loc[0]
	wg := &sync.WaitGroup{}
	wg.Add(len(nums) - 1)

	for i, v := range loc[1:] {
		go func(ind, val int) {
			defer wg.Done()
			diffs[ind+1] = val - loc[ind]
		}(i, v)
	}

	wg.Wait()

	counter := 0
	prod := 1
	register := map[int]int{}
	for i, v := range diffs {
		if v == 3 {
			if _, ok := register[counter]; !ok {
				conf := confForOneChain(counter)
				register[counter] = conf
			}
			prod *= register[counter]
			counter = 0
		} else {
			counter++
			if i == len(diffs)-1 {
				if _, ok := register[counter]; !ok {
					conf := confForOneChain(counter)
					register[counter] = conf
				}
				prod *= register[counter]
			}
		}

	}

	return prod
}

func confForOneChain(l int) int {
	var out int64
	wg := sync.WaitGroup{}
	for count2 := 0; count2 <= l/2; count2++ {
		for count3 := 0; count3 <= (l-2*count2)/3; count3++ {
			wg.Add(1)

			go func(c2, c3 int) {
				defer wg.Done()
				c1 := l - 2*c2 - 3*c3
				counts := []int{c1, c2, c3}
				sort.Ints(counts)

				prod := int64(1)
				for p := counts[2] + 1; p <= c1+c2+c3; p++ {
					prod *= int64(p)
				}

				for p := 2; p <= counts[1]; p++ {
					prod /= int64(p)
				}

				for p := 2; p <= counts[0]; p++ {
					prod /= int64(p)
				}
				atomic.AddInt64(&out, prod)

			}(count2, count3)

		}
	}

	wg.Wait()

	return int(out)
}
