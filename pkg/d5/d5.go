package d5

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/sync/errgroup"
)

// Seat ..
type Seat struct {
	Row, Column int
}

// ID ..
func (s *Seat) ID() int {
	return s.Row*8 + s.Column
}

// DecodeOne ..
func DecodeOne(s string) (*Seat, error) {
	rowString := s[:7]
	colString := s[7:]

	if len(rowString) != 7 || len(colString) != 3 {
		return nil, fmt.Errorf("wrong number of sections")
	}

	rowBString := strings.ReplaceAll(strings.ReplaceAll(rowString, "F", "0"), "B", "1")
	colBString := strings.ReplaceAll(strings.ReplaceAll(colString, "L", "0"), "R", "1")

	row, err := strconv.ParseInt(rowBString, 2, 64)
	if err != nil {
		return nil, err
	}
	col, err := strconv.ParseInt(colBString, 2, 64)
	if err != nil {
		return nil, err
	}

	return &Seat{
		int(row), int(col),
	}, nil
}

// ParseInput ..
func ParseInput(f []byte) ([]*Seat, error) {
	split := strings.Split(string(f), "\n")
	seats := make([]*Seat, len(split))
	errGrp := &errgroup.Group{}
	for i, v := range split {
		errGrp.Go(func(ind int, val string) func() error {
			return func() error {
				seat, err := DecodeOne(val)
				if err != nil {
					return err
				}
				seats[ind] = seat
				return nil
			}
		}(i, v))
	}

	err := errGrp.Wait()
	if err != nil {
		return nil, err
	}
	return seats, nil
}

// SolvePart1 ..
func SolvePart1(seats []*Seat) int {
	ids := func() []int {
		out := make([]int, len(seats))
		for i, v := range seats {
			out[i] = v.ID()
		}
		return out
	}()
	sort.Ints(ids)
	return ids[len(ids)-1]
}

// SolvePart2 ..
func SolvePart2(seats []*Seat) int {
	ids := func() []int {
		out := make([]int, len(seats))
		for i, v := range seats {
			out[i] = v.ID()
		}
		return out
	}()
	sort.Ints(ids)
	possibleIDs := func() []int {
		out := []int{}
		for i := 0; i < 128; i++ {
			for j := 0; j < 8; j++ {
				out = append(out, i*8+j)
			}
		}
		return out
	}()

	exists := func(id int) bool {
		for _, v := range ids {
			if id == v {
				return true
			}
		}
		return false
	}

	diff := func() []int {
		out := []int{}
		for i, v := range possibleIDs {
			// if not in existing IDs then append
			if !exists(v) {
				out = append(out, i)
			}
		}
		return out
	}()

	for i, v := range diff {
		if v+1 != diff[i+1] {
			return possibleIDs[diff[i+1]]
		}
	}
	return 0
}
