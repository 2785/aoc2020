package d24

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/sync/errgroup"
)

// Direction ..
type Direction int

const (
	e Direction = iota
	se
	sw
	w
	nw
	ne
)

func newDirection(s string) (Direction, error) {
	switch s {
	case "e":
		return e, nil
	case "se":
		return se, nil
	case "sw":
		return sw, nil
	case "w":
		return w, nil
	case "nw":
		return nw, nil
	case "ne":
		return ne, nil
	default:
		return 0, fmt.Errorf("unknown direction %s", s)
	}
}

func (d Direction) getDeltas() (dx, dy int) {
	switch d {
	case e:
		dx, dy = 1, 0
	case se:
		dx, dy = 1, -1
	case sw:
		dx, dy = 0, -1
	case w:
		dx, dy = -1, 0
	case nw:
		dx, dy = -1, 1
	case ne:
		dx, dy = 0, 1
	default:
		panic("this should never happen")
	}
	return
}

// ParseInput ..
func ParseInput(f []byte) ([][]Direction, error) {
	dirReg := regexp.MustCompile("e|se|sw|w|nw|ne")
	split := strings.Split(strings.ReplaceAll(string(f), "\r", ""), "\n")
	out := make([][]Direction, len(split))
	eg := &errgroup.Group{}

	for i, line := range split {
		eg.Go(func(ind int, curr string) func() error {
			return func() error {
				matches := dirReg.FindAllString(curr, -1)
				if matches == nil {
					return fmt.Errorf("no matches found in %s", curr)
				}
				dirs := make([]Direction, len(matches))
				for i, v := range matches {
					d, err := newDirection(v)
					if err != nil {
						return err
					}
					dirs[i] = d
				}
				out[ind] = dirs
				return nil
			}
		}(i, line))
	}

	err := eg.Wait()
	if err != nil {
		return nil, err
	}

	return out, nil
}

// SolvePart1 ..
func SolvePart1(directions [][]Direction) int {
	grid := make(map[int]map[int]bool)
	for _, ds := range directions {
		x, y := 0, 0
		for _, d := range ds {
			dx, dy := d.getDeltas()
			x += dx
			y += dy
		}
		if _, ok := grid[x]; !ok {
			grid[x] = make(map[int]bool)
		}
		if _, ok := grid[x][y]; !ok {
			grid[x][y] = true
		} else {
			delete(grid[x], y)
		}
	}

	tot := 0

	for _, v := range grid {
		tot += len(v)
	}

	return tot
}

// SolvePart2 ..
func SolvePart2(directions [][]Direction) int {
	grid := make(map[int]map[int]bool)
	for _, ds := range directions {
		x, y := 0, 0
		for _, d := range ds {
			dx, dy := d.getDeltas()
			x += dx
			y += dy
		}
		if _, ok := grid[x]; !ok {
			grid[x] = make(map[int]bool)
		}
		if _, ok := grid[x][y]; !ok {
			grid[x][y] = true
		} else {
			delete(grid[x], y)
		}
	}

	// grid will by mutated
	// prepare grid points to evaluate
	var target map[int]map[int]bool

	for i := 0; i < 100; i++ {
		target = make(map[int]map[int]bool)
		for x, ymap := range grid {
			for y, black := range ymap {
				// should not be necessary, but just in case
				if black {
					if _, ok := target[x]; !ok {
						target[x] = make(map[int]bool)
					}
					target[x][y] = black
					xanchor, yanchor := x, y

					for _, d := range []Direction{e, se, sw, w, nw, ne} {
						dx, dy := d.getDeltas()
						xneighbour, yneigbhour := xanchor+dx, yanchor+dy
						// if the point is not black in the grid, set it to white in the target
						if func() bool {
							if _, ok := grid[xneighbour]; !ok {
								return true
							}
							if _, ok := grid[xneighbour][yneigbhour]; !ok {
								return true
							}
							if grid[xneighbour][yneigbhour] {
								return false
							}
							panic("should not get here :)")
						}() {
							if _, ok := target[xneighbour]; !ok {
								target[xneighbour] = make(map[int]bool)
							}
							target[xneighbour][yneigbhour] = false
						}
						// if point is black on the grid, ignore it, it will be dealt with later
					}

				}
			}
		}
		// done constructing the target
		// now loop through the target to determine flip conditions
		for x, ymap := range target {
			for y, black := range ymap {
				// count how many black pieces are around the x y tile
				count := 0
				xanchor, yanchor := x, y
				for _, d := range []Direction{e, se, sw, w, nw, ne} {
					dx, dy := d.getDeltas()
					xneighbour, yneigbhour := xanchor+dx, yanchor+dy
					if _, ok := target[xneighbour]; ok {
						if isblack, ok := target[xneighbour][yneigbhour]; ok && isblack {
							count++
						}
					}
				}

				if black {
					if count == 0 || count > 2 {
						delete(grid[x], y)
					}
				} else {
					if count == 2 {
						if _, ok := grid[x]; !ok {
							grid[x] = make(map[int]bool)
						}
						grid[x][y] = true
					}
				}
			}
		}
		// now grid has been flipped as appropriate
	}

	tot := 0
	for _, y := range grid {
		tot += len(y)
	}

	return tot
}
