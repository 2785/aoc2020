package d23

import (
	"container/ring"
	"fmt"
	"strconv"
	"strings"
)

// ParseInput ..
func ParseInput(f []byte) ([]int, error) {
	split := strings.Split(string(f), "")
	ints := make([]int, len(split))
	for i, v := range split {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ints[i] = num
	}
	return ints, nil
}

func playOneRound(curr *ring.Ring, max int, ringmap []*ring.Ring) *ring.Ring {
	currValue := curr.Value.(int)
	take1 := curr.Next().Value.(int)
	take2 := curr.Move(2).Value.(int)
	take3 := curr.Move(3).Value.(int)

	destNum := currValue
	for {
		destNum--
		if destNum == 0 {
			destNum = max
		}
		if destNum != take1 && destNum != take2 && destNum != take3 {
			break
		}
	}

	dest := ringmap[destNum]
	dest.Link(curr.Link(curr.Move(4)))

	return curr.Next()
}

// SolvePart1 ..
func SolvePart1(input []int) string {

	m := make([]*ring.Ring, len(input)+1)
	game := ring.New(len(input))
	for _, v := range input {
		game.Value = v
		m[v] = game
		game = game.Next()
	}

	max := len(input)

	for i := 0; i < 100; i++ {

		game = playOneRound(game, max, m)

	}

	for game.Value.(int) != 1 {
		game = game.Next()
	}

	out := ""

	l := len(input)
	for i := 0; i < l-1; i++ {
		game = game.Next()
		out += fmt.Sprintf("%v", game.Value.(int))
	}

	return out
}

// SolvePart2 ..
func SolvePart2(input []int) int {
	loc := make([]int, 1e6)
	for i, v := range input {
		loc[i] = v
	}
	for i := len(input); i < 1e6; i++ {
		loc[i] = i + 1
	}
	m := make([]*ring.Ring, 1e6+1)

	game := ring.New(len(loc))
	for _, v := range loc {
		game.Value = v
		m[v] = game
		game = game.Next()
	}

	for i := 0; i < 1e7; i++ {
		game = playOneRound(game, 1e6, m)
	}

	for game.Value.(int) != 1 {
		game = game.Next()
	}

	out := 1

	out *= game.Next().Value.(int)
	out *= game.Move(2).Value.(int)

	return out
}
