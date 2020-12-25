package d25

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ParseInput ..
func ParseInput(f []byte) (int, int, error) {
	split := strings.Split(strings.ReplaceAll(string(f), "\r", ""), "\n")
	if len(split) != 2 {
		return 0, 0, fmt.Errorf("number of lines doesn't look right, expecting 2 got %v", len(split))
	}
	num1, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, err
	}
	num2, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, err
	}

	return num1, num2, nil
}

func getLoopCountFromPublicKey(k int) (int, error) {
	p := 1
	for i := 1; i < 1e7; i++ {
		p *= 7
		p = p % 20201227
		if p == k {
			return i, nil
		}
	}
	return 0, errors.New("run out of iterations")
}

// SolvePart1 ..
func SolvePart1(num1, num2 int) (int, error) {
	l2, err := getLoopCountFromPublicKey(num2)
	if err != nil {
		return 0, err
	}

	prod := 1
	for i := 0; i < l2; i++ {
		prod *= num1
		prod = prod % 20201227
	}
	return prod, nil
}
