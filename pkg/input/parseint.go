package input

import (
	"strconv"
	"strings"
)

// MustParseInt ..
func MustParseInt(f []byte) []int {
	split := strings.Split(string(f), "\n")
	out := make([]int, len(split))
	for i, v := range split {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		out[i] = num
	}
	return out
}
