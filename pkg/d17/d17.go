package d17

import "strings"

// ParseInput ..
func ParseInput(f []byte) map[int]map[int]map[int]bool {
	slice := make(map[int]map[int]bool)
	splitLine := strings.Split(string(f), "\n")
	for row, line := range splitLine {
		slice[row] = make(map[int]bool)
		split := strings.Split(line, "")
		for i, v := range split {
			if v == "#" {
				slice[row][i] = true
			}
		}
	}

	return map[int]map[int]map[int]bool{
		0: slice,
	}
}

// SolvePart1 ..
func SolvePart1(start map[int]map[int]map[int]bool) int {
	loc1 := make(map[int]map[int]map[int]bool, len(start))
	loc2 := make(map[int]map[int]map[int]bool, len(start))
	// grab a local copy in order to mutate
	for x, yz := range start {
		loc1[x] = make(map[int]map[int]bool)
		loc2[x] = make(map[int]map[int]bool)
		for y, z := range yz {
			loc1[x][y] = make(map[int]bool)
			loc2[x][y] = make(map[int]bool)
			for zval, state := range z {
				loc1[x][y][zval] = state
				loc2[x][y][zval] = state
			}
		}
	}

	type delta struct {
		dx, dy, dz int
	}

	deltas := func() []delta {
		out := make([]delta, 0)
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				for z := -1; z <= 1; z++ {
					if !(x == 0 && y == 0 && z == 0) {
						out = append(out, delta{x, y, z})
					}
				}
			}
		}
		return out
	}()

	for i := 0; i < 6; i++ {
		// start each round by filling up the empty spaces surrounding on states
		for x, xmap := range loc2 {
			for y, ymap := range xmap {
				for z, on := range ymap {
					if on {
						for _, del := range deltas {
							if _, ok := loc1[x+del.dx]; !ok {
								loc1[x+del.dx] = make(map[int]map[int]bool)
							}

							if _, ok := loc1[x+del.dx][y+del.dy]; !ok {
								loc1[x+del.dx][y+del.dy] = make(map[int]bool)
							}

							if _, ok := loc1[x+del.dx][y+del.dy][z+del.dz]; !ok {
								loc1[x+del.dx][y+del.dy][z+del.dz] = false
							}
						}
					}
				}
			}
		}

		// at this point loc1 has the off grids around the on grids filled up,
		// it contains all the points we are interested in evaluating in this
		// round, loc2 only contains the on states

		for x, xmap := range loc1 {
			for y, ymap := range xmap {
				for z, on := range ymap {
					// First we count the on states around each entry we are interested in
					count := 0
					for _, del := range deltas {
						if _, ok := loc1[x+del.dx]; ok {
							if _, ok := loc1[x+del.dx][y+del.dy]; ok {
								if locIsOn, ok := loc1[x+del.dx][y+del.dy][z+del.dz]; ok && locIsOn {
									// if that spot exists in loc2 it's on
									count++
								}
							}
						}
					}

					if on && !(count == 3 || count == 2) {
						// Delete it from loc2 as it is no longer relevant
						delete(loc2[x][y], z)
					}

					if !on && count == 3 {
						if _, ok := loc2[x]; !ok {
							loc2[x] = make(map[int]map[int]bool)
						}

						if _, ok := loc2[x][y]; !ok {
							loc2[x][y] = make(map[int]bool)
						}

						loc2[x][y][z] = true
					}
				}
			}
		}

		// at this point loc2 has been updated with the appropriate on states

		// now to set loc1 to a copy of loc2
		loc1 = make(map[int]map[int]map[int]bool, len(loc2))
		for x, yz := range loc2 {
			loc1[x] = make(map[int]map[int]bool)
			for y, z := range yz {
				loc1[x][y] = make(map[int]bool)
				for zval, state := range z {
					loc1[x][y][zval] = state
				}
			}
		}
	}

	sum := 0

	for _, yz := range loc2 {
		for _, z := range yz {
			sum += len(z)
		}
	}

	return sum
}

// SolvePart2 ..
func SolvePart2(start map[int]map[int]map[int]bool) int {
	loc1 := make(map[int]map[int]map[int]map[int]bool)
	loc2 := make(map[int]map[int]map[int]map[int]bool)
	loc1[0] = make(map[int]map[int]map[int]bool)
	loc2[0] = make(map[int]map[int]map[int]bool)
	// grab a local copy in order to mutate
	for x, yz := range start {
		loc1[0][x] = make(map[int]map[int]bool)
		loc2[0][x] = make(map[int]map[int]bool)
		for y, z := range yz {
			loc1[0][x][y] = make(map[int]bool)
			loc2[0][x][y] = make(map[int]bool)
			for zval, state := range z {
				loc1[0][x][y][zval] = state
				loc2[0][x][y][zval] = state
			}
		}
	}

	type delta struct {
		dx, dy, dz, dw int
	}

	deltas := func() []delta {
		out := make([]delta, 0)
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				for z := -1; z <= 1; z++ {
					for w := -1; w <= 1; w++ {
						if !(x == 0 && y == 0 && z == 0 && w == 0) {
							out = append(out, delta{x, y, z, w})
						}
					}
				}
			}
		}
		return out
	}()

	for i := 0; i < 6; i++ {
		// start each round by filling up the empty spaces surrounding on states
		for x, xmap := range loc2 {
			for y, ymap := range xmap {
				for z, zmap := range ymap {
					for w, on := range zmap {
						if on {
							for _, del := range deltas {
								if _, ok := loc1[x+del.dx]; !ok {
									loc1[x+del.dx] = make(map[int]map[int]map[int]bool)
								}

								if _, ok := loc1[x+del.dx][y+del.dy]; !ok {
									loc1[x+del.dx][y+del.dy] = make(map[int]map[int]bool)
								}

								if _, ok := loc1[x+del.dx][y+del.dy][z+del.dz]; !ok {
									loc1[x+del.dx][y+del.dy][z+del.dz] = make(map[int]bool)
								}

								if _, ok := loc1[x+del.dx][y+del.dy][z+del.dz][w+del.dw]; !ok {
									loc1[x+del.dx][y+del.dy][z+del.dz][w+del.dw] = false
								}
							}
						}
					}
				}
			}
		}

		// at this point loc1 has the off grids around the on grids filled up,
		// it contains all the points we are interested in evaluating in this
		// round, loc2 only contains the on states

		for x, xmap := range loc1 {
			for y, ymap := range xmap {
				for z, zmap := range ymap {
					for w, on := range zmap {
						// First we count the on states around each entry we are interested in
						count := 0
						for _, del := range deltas {
							if _, ok := loc1[x+del.dx]; ok {
								if _, ok := loc1[x+del.dx][y+del.dy]; ok {
									if _, ok := loc1[x+del.dx][y+del.dy][z+del.dz]; ok {
										if locIsOn, ok := loc1[x+del.dx][y+del.dy][z+del.dz][w+del.dw]; ok && locIsOn {
											count++
										}
									}
								}
							}
						}

						if on && !(count == 3 || count == 2) {
							// Delete it from loc2 as it is no longer relevant
							delete(loc2[x][y][z], w)
						}

						if !on && count == 3 {
							if _, ok := loc2[x]; !ok {
								loc2[x] = make(map[int]map[int]map[int]bool)
							}

							if _, ok := loc2[x][y]; !ok {
								loc2[x][y] = make(map[int]map[int]bool)
							}

							if _, ok := loc2[x][y][z]; !ok {
								loc2[x][y][z] = make(map[int]bool)
							}

							loc2[x][y][z][w] = true
						}
					}

				}
			}
		}

		// at this point loc2 has been updated with the appropriate on states

		// now to set loc1 to a copy of loc2
		loc1 = make(map[int]map[int]map[int]map[int]bool, len(start))
		for x, yzw := range loc2 {
			loc1[x] = make(map[int]map[int]map[int]bool)
			for y, zw := range yzw {
				loc1[x][y] = make(map[int]map[int]bool)
				for z, w := range zw {
					loc1[x][y][z] = make(map[int]bool)
					for wval, state := range w {
						loc1[x][y][z][wval] = state
					}
				}
			}
		}
	}

	sum := 0

	for _, yzw := range loc2 {
		for _, zw := range yzw {
			for _, w := range zw {
				sum += len(w)
			}
		}
	}

	return sum
}
