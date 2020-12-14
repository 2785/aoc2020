package d14

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/2785/aoc2020/pkg/regutil"
)

// Mask ..
type Mask string

// Apply ..
func (m Mask) Apply(num int) string {
	bin := []rune(fmt.Sprintf("%036b", num))
	for i, v := range m {
		switch v {
		case '1':
			bin[i] = '1'
		case '0':
			bin[i] = '0'
		}
	}
	return string(bin)
}

// InitializationProgram ..
type InitializationProgram map[Mask]map[int]int

// ParseInput ..
func ParseInput(f []byte) (InitializationProgram, []Mask, error) {
	lines := strings.Split(string(f), "\n")
	out := make(InitializationProgram)
	ordered := []Mask{}
	maskRe := regexp.MustCompile(`mask = (?P<mask>[X01]{36})`)
	instRe := regexp.MustCompile(`mem\[(?P<loc>\d+)\] = (?P<num>\d+)`)

	var currMask Mask = ""

	for i, v := range lines {
		if matches, err := regutil.MustCaptureNamedGroup(maskRe, []string{"mask"}, v); err == nil {
			currMask = Mask(matches["mask"])
			out[currMask] = make(map[int]int)
			ordered = append(ordered, currMask)
			continue
		}

		if i == 0 {
			return nil, nil, errors.New("first entry is not a mask")
		}

		if matches, err := regutil.MustCaptureNamedGroup(instRe, []string{"loc", "num"}, v); err == nil {
			loc, err := strconv.Atoi(matches["loc"])
			if err != nil {
				return nil, nil, fmt.Errorf("cannot parse %s into int: %w", matches["loc"], err)
			}
			num, err := strconv.Atoi(matches["num"])
			if err != nil {
				return nil, nil, fmt.Errorf("cannot parse %s into int: %w", matches["num"], err)
			}
			out[currMask][loc] = num
			continue
		}

		return nil, nil, fmt.Errorf("line '%s' did not match anything", v)
	}

	return out, ordered, nil
}

// SolvePart1 ..
func SolvePart1(p InitializationProgram, order []Mask) (int, error) {
	memory := make(map[int]int)
	for _, v := range order {
		instructions := p[v]
		for loc, val := range instructions {
			masked, err := strconv.ParseInt(v.Apply(val), 2, 64)
			if err != nil {
				return 0, err
			}
			memory[loc] = int(masked)
		}
	}

	sum := 0
	for _, v := range memory {
		sum += v
	}

	return sum, nil
}
