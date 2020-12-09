package d9

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

// ParseInput ..
func ParseInput(f []byte) ([]int, error) {
	split := strings.Split(string(f), "\n")
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

func isValidSum(operands []int, sum int) bool {
	for i, left := range operands[:len(operands)-1] {
		for _, right := range operands[i+1:] {
			if left+right == sum {
				return true
			}
		}
	}
	return false
}

// SolvePart1 ..
func SolvePart1(nums []int, lookback int) (int, error) {
	for i, v := range nums[lookback:] {
		if !isValidSum(nums[i:i+lookback], v) {
			return v, nil
		}
	}

	return 0, errors.New("not found")
}

// SolvePart2 ..
func SolvePart2(nums []int, target int) (int, error) {
	start, end, sum := 0, 0, 0
	for start < len(nums)-1 && end <= len(nums) {
		if sum == target {
			copied := make([]int, end-start)
			copy(copied, nums[start:end])
			sort.Ints(copied)
			return copied[0] + copied[len(copied)-1], nil
		}
		if sum < target {
			sum += nums[end]
			end++
		} else {
			sum -= nums[start]
			start++
		}
	}

	return 0, errors.New("not found")
}
