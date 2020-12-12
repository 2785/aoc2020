package d11

import (
	"errors"
	"fmt"
	"strings"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

// Spot ..
type Spot int

const (
	empty Spot = iota
	occupied
	floor
)

// Spots ..
type Spots [][]Spot

// ShouldFlip ..
func (s Spots) ShouldFlip(r, c int) (bool, error) {
	if r >= len(s) || c >= len(s[r]) || r < 0 || c < 0 {
		return false, errors.New("out of range")
	}

	if s[r][c] == floor {
		return false, nil
	}
	occupiedCount := 0

	checkOccupied := func(y, x int) {
		if s[y][x] == occupied {
			occupiedCount++
		}
	}

	// TL
	if r != 0 && c != 0 {
		checkOccupied(r-1, c-1)
	}
	// T
	if r != 0 {
		checkOccupied(r-1, c)
	}
	// TR
	if r != 0 && c != len(s[r])-1 {
		checkOccupied(r-1, c+1)
	}
	// L
	if c != 0 {
		checkOccupied(r, c-1)
	}
	// R
	if c != len(s[r])-1 {
		checkOccupied(r, c+1)
	}
	// BL
	if r != len(s)-1 && c != 0 {
		checkOccupied(r+1, c-1)
	}
	// B
	if r != len(s)-1 {
		checkOccupied(r+1, c)
	}
	// BR
	if r != len(s)-1 && c != len(s[r])-1 {
		checkOccupied(r+1, c+1)
	}

	switch s[r][c] {
	case occupied:
		if occupiedCount >= 4 {
			return true, nil
		}
	case empty:
		if occupiedCount == 0 {
			return true, nil
		}
	}
	return false, nil
}

// ValidCoordinate ..
func (s Spots) ValidCoordinate(r, c int) bool {
	return r >= 0 && r < len(s) && c >= 0 && c < len(s[r])
}

// ShouldFlipButInRetardMode ..
func (s Spots) ShouldFlipButInRetardMode(r, c int) (bool, error) {
	if r >= len(s) || c >= len(s[r]) || r < 0 || c < 0 {
		return false, errors.New("out of range")
	}

	if s[r][c] == floor {
		return false, nil
	}
	occupiedCount := 0

	checkOccupied := func(y, x, dy, dx int) {
		for {
			y += dy
			x += dx
			if !s.ValidCoordinate(y, x) {
				return
			}
			switch s[y][x] {
			case occupied:
				occupiedCount++
				return
			case empty:
				return
			}
		}
	}

	// TL
	checkOccupied(r, c, -1, -1)
	// T
	checkOccupied(r, c, -1, 0)
	// TR
	checkOccupied(r, c, -1, 1)
	// L
	checkOccupied(r, c, 0, -1)
	// R
	checkOccupied(r, c, 0, 1)
	// BL
	checkOccupied(r, c, 1, -1)
	// B
	checkOccupied(r, c, 1, 0)
	// BR
	checkOccupied(r, c, 1, 1)

	switch s[r][c] {
	case occupied:
		if occupiedCount >= 5 {
			return true, nil
		}
	case empty:
		if occupiedCount == 0 {
			return true, nil
		}
	}
	return false, nil
}

// Count ..
func (s Spots) Count(t Spot) int {
	count := 0
	for _, v := range s {
		for _, s := range v {
			if s == t {
				count++
			}
		}
	}
	return count
}

// ParseInput .
func ParseInput(f []byte) (Spots, error) {
	lines := strings.Split(string(f), "\n")
	spots := make(Spots, len(lines))

	eg := &errgroup.Group{}
	for i, line := range lines {
		eg.Go(func(ind int, val string) func() error {
			return func() error {
				split := strings.Split(val, "")
				out := make([]Spot, len(split))
				for i, v := range split {
					switch v {
					case "L":
						out[i] = empty
					case ".":
						out[i] = floor
					case "#":
						out[i] = occupied
					default:
						return fmt.Errorf("got unrecognized entry %s", v)
					}
				}
				spots[ind] = out
				return nil
			}
		}(i, line))
	}
	err := eg.Wait()
	if err != nil {
		return nil, err
	}
	return spots, nil
}

// SolvePart1 ..
func SolvePart1(spots Spots) (int, error) {
	loc := deepCopy(spots)
	counter := 0
	for counter <= 1e5 {
		copied := deepCopy(loc)
		eg := &errgroup.Group{}
		var flipped int64
		for i, v := range loc {
			for j, s := range v {
				eg.Go(func(y, x int, thing Spot) func() error {
					return func() error {
						shouldFlip, err := loc.ShouldFlip(y, x)
						if err != nil {
							return err
						}
						if shouldFlip {
							atomic.AddInt64(&flipped, 1)
							switch thing {
							case empty:
								copied[y][x] = occupied
							case occupied:
								copied[y][x] = empty
							default:
								return fmt.Errorf("unknown option")
							}
						}
						return nil
					}
				}(i, j, s))
			}
		}
		err := eg.Wait()
		if err != nil {
			return 0, err
		}
		if flipped == 0 {
			return loc.Count(occupied), nil
		}
		loc = copied
		counter++
	}
	return 0, errors.New("reached iteration limit")
}

// SolvePart2 ..
func SolvePart2(spots Spots) (int, error) {
	loc := deepCopy(spots)
	counter := 0
	for counter <= 1e5 {
		copied := deepCopy(loc)
		eg := &errgroup.Group{}
		var flipped int64
		for i, v := range loc {
			for j, s := range v {
				eg.Go(func(y, x int, thing Spot) func() error {
					return func() error {
						shouldFlip, err := loc.ShouldFlipButInRetardMode(y, x)
						if err != nil {
							return err
						}
						if shouldFlip {
							atomic.AddInt64(&flipped, 1)
							switch thing {
							case empty:
								copied[y][x] = occupied
							case occupied:
								copied[y][x] = empty
							default:
								return fmt.Errorf("unknown option")
							}
						}
						return nil
					}
				}(i, j, s))
			}
		}
		err := eg.Wait()
		if err != nil {
			return 0, err
		}
		if flipped == 0 {
			return loc.Count(occupied), nil
		}
		loc = copied
		counter++
	}
	return 0, errors.New("reached iteration limit")
}

func deepCopy(src Spots) Spots {
	out := make(Spots, len(src))
	for i, v := range src {
		out[i] = make([]Spot, len(v))
		copy(out[i], v)
	}
	return out
}
