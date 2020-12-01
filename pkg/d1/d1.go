package d1

import (
	"strconv"
	"strings"

	"github.com/gammazero/workerpool"
)

// ParseFile ..
func ParseFile(f []byte) ([]int, error) {
	parts := strings.Split(string(f), "\n")
	out := make([]int, len(parts))

	for i, v := range parts {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		out[i] = num
	}
	return out, nil
}

// SolvePart1 ..
func SolvePart1(in []int) int {
	stopChan := make(chan struct{})
	var num1, num2 int
	wp := workerpool.New(20)

	for i, v1 := range in {
		ind := i
		val := v1
		wp.Submit(func() {
			for _, v2 := range in[ind+1:] {
				if val+v2 == 2020 {
					num1 = val
					num2 = v2
					stopChan <- struct{}{}
				}
			}
		})
	}

	<-stopChan
	wp.Stop()

	return num1 * num2
}

// SolvePart2 ...
func SolvePart2(in []int) int {
	stopChan := make(chan struct{})
	var num1, num2, num3 int
	wp := workerpool.New(20)

	for i, v1 := range in {
		ind1 := i
		val1 := v1
		for j, v2 := range in[ind1+1:] {
			ind2 := j
			val2 := v2
			wp.Submit(func() {
				for _, v3 := range in[ind1+ind2+2:] {
					if val1+val2+v3 == 2020 {
						num1 = val1
						num2 = val2
						num3 = v3
						stopChan <- struct{}{}
					}
				}
			})
		}
	}

	<-stopChan
	wp.Stop()

	return num1 * num2 * num3
}
