package d12

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/sync/errgroup"
)

// InstEntry ..
type InstEntry struct {
	Instruction Instruction
	Magnitude   int
}

// Instruction ..
type Instruction int

const (
	north Instruction = iota
	south
	east
	west
	left
	right
	forward
)

// Direction ..
type Direction int

// Turn ..
func (d Direction) Turn(angle int) Direction {
	fin := d + Direction(angle/90)
	if fin < 0 {
		fin += 4
	}
	if fin > 3 {
		fin -= 4
	}
	return fin
}

const (
	n Direction = 0
	e Direction = 1
	s Direction = 2
	w Direction = 3
)

// Boat ..
type Boat struct {
	X, Y      int
	Direction Direction
}

// Apply ..
func (b *Boat) Apply(ins Instruction, mag int) error {
	switch ins {
	case north:
		b.Y += mag
	case south:
		b.Y -= mag
	case east:
		b.X += mag
	case west:
		b.X -= mag
	case left:
		b.Direction = b.Direction.Turn(-mag)
	case right:
		b.Direction = b.Direction.Turn(mag)
	case forward:
		switch b.Direction {
		case n:
			b.Y += mag
		case s:
			b.Y -= mag
		case e:
			b.X += mag
		case w:
			b.X -= mag
		default:
			return errors.New("unknown direction to go forward")
		}
	default:
		return errors.New("unknown instruction")
	}
	return nil
}

// ApplyFancyTurn ..
func (b *Boat) ApplyFancyTurn(mag int) error {
	mag = mag / 90
	if mag < 0 {
		mag += 4
	}
	x := b.X
	y := b.Y
	switch mag {
	case 1:
		b.X = y
		b.Y = -x
	case 2:
		b.X = -x
		b.Y = -y
	case 3:
		b.X = -y
		b.Y = x
	default:
		return fmt.Errorf("can't process clockwise turn of %v degrees", mag)
	}
	return nil
}

// ParseInput ..
func ParseInput(f []byte) ([]InstEntry, error) {
	lines := strings.Split(string(f), "\n")
	out := make([]InstEntry, len(lines))

	eg := &errgroup.Group{}
	for i, l := range lines {
		eg.Go(func(ind int, txt string) func() error {
			return func() error {
				ins := txt[:1]
				mag := txt[1:]
				num, err := strconv.Atoi(mag)
				if err != nil {
					return err
				}
				var decodedInstruction Instruction
				switch ins {
				case "N":
					decodedInstruction = north
				case "S":
					decodedInstruction = south
				case "E":
					decodedInstruction = east
				case "W":
					decodedInstruction = west
				case "L":
					decodedInstruction = left
				case "R":
					decodedInstruction = right
				case "F":
					decodedInstruction = forward
				default:
					return fmt.Errorf("unknown instruction: %s", txt)
				}
				out[ind] = InstEntry{decodedInstruction, num}
				return nil
			}
		}(i, l))
	}
	err := eg.Wait()
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SolvePart1 ..
func SolvePart1(instructions []InstEntry) int {
	boat := &Boat{
		X:         0,
		Y:         0,
		Direction: e,
	}
	for _, v := range instructions {
		boat.Apply(v.Instruction, v.Magnitude)
	}
	sum := 0
	if boat.X < 0 {
		sum -= boat.X
	} else {
		sum += boat.X
	}

	if boat.Y < 0 {
		sum -= boat.Y
	} else {
		sum += boat.Y
	}

	return sum
}

// SolvePart2 ..
func SolvePart2(instructions []InstEntry) int {
	wayPoint := &Boat{
		X: 10,
		Y: 1,
	}
	boat := &Boat{
		X: 0,
		Y: 0,
	}

	for _, v := range instructions {
		switch v.Instruction {
		case north, west, east, south:
			wayPoint.Apply(v.Instruction, v.Magnitude)
		case left:
			wayPoint.ApplyFancyTurn(-v.Magnitude)
		case right:
			wayPoint.ApplyFancyTurn(v.Magnitude)
		case forward:
			boat.Apply(north, v.Magnitude*wayPoint.Y)
			boat.Apply(east, v.Magnitude*wayPoint.X)
		}
	}

	sum := 0
	if boat.X < 0 {
		sum -= boat.X
	} else {
		sum += boat.X
	}

	if boat.Y < 0 {
		sum -= boat.Y
	} else {
		sum += boat.Y
	}

	return sum
}
