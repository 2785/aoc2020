package d22

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mitchellh/hashstructure/v2"
)

// ParseInput ..
func ParseInput(f []byte) (IntQueue, IntQueue, error) {
	file := strings.ReplaceAll(string(f), "\r", "")
	split := strings.Split(file, "\n\n")
	if len(split) != 2 {
		return nil, nil, errors.New("unexpected number of segments")
	}

	player1s := strings.Split(split[0], "\n")[1:]
	player1num := make([]int, len(player1s))
	for i, v := range player1s {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, nil, err
		}
		player1num[i] = num
	}

	player2s := strings.Split(split[1], "\n")[1:]
	player2num := make([]int, len(player2s))
	for i, v := range player2s {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, nil, err
		}
		player2num[i] = num
	}

	return player1num, player2num, nil
}

func playGame(p1, p2 IntQueue) (IntQueue, IntQueue) {
	loc1, loc2 := make(IntQueue, len(p1)), make(IntQueue, len(p2))
	copy(loc1, p1)
	copy(loc2, p2)
	for len(loc1) > 0 && len(loc2) > 0 {
		hand1 := loc1.MustPop()
		hand2 := loc2.MustPop()
		if hand1 > hand2 {
			loc1.Push(hand1)
			loc1.Push(hand2)
		} else {
			loc2.Push(hand2)
			loc2.Push(hand1)
		}
	}

	return loc1, loc2
}

func playRecursiveGame(p1, p2 IntQueue) (int, IntQueue, IntQueue, error) {
	loc1, loc2 := make(IntQueue, len(p1)), make(IntQueue, len(p2))
	copy(loc1, p1)
	copy(loc2, p2)

	configStore := make(map[uint64]map[uint64]struct{})
	for len(loc1) > 0 && len(loc2) > 0 {
		hash1, err := hashstructure.Hash(loc1, hashstructure.FormatV2, nil)
		if err != nil {
			return 0, nil, nil, err
		}
		hash2, err := hashstructure.Hash(loc2, hashstructure.FormatV2, nil)
		if err != nil {
			return 0, nil, nil, err
		}

		if _, ok := configStore[hash1]; !ok {
			configStore[hash1] = make(map[uint64]struct{})
		}

		if _, ok := configStore[hash1][hash2]; ok {
			return 1, loc1, loc2, nil
		}

		configStore[hash1][hash2] = struct{}{}

		roundWinner := 0
		hand1 := loc1.MustPop()
		hand2 := loc2.MustPop()
		if len(loc1) >= hand1 && len(loc2) >= hand2 {
			// enters subgame
			roundWinner, _, _, err = playRecursiveGame(loc1[:hand1], loc2[:hand2])
			if err != nil {
				return 0, nil, nil, fmt.Errorf("error playing subgame: %w", err)
			}
		} else {
			switch hand1 > hand2 {
			case true:
				roundWinner = 1
			case false:
				roundWinner = 2
			}
		}

		switch roundWinner {
		case 1:
			loc1.Push(hand1)
			loc1.Push(hand2)
		case 2:
			loc2.Push(hand2)
			loc2.Push(hand1)
		default:
			return 0, nil, nil, fmt.Errorf("got unexpected round winner")
		}

	}

	if len(loc1) == 0 {
		return 2, loc1, loc2, nil
	}
	return 1, loc1, loc2, nil
}

// SolvePart1 ..
func SolvePart1(p1, p2 IntQueue) int {
	outcome1, outcome2 := playGame(p1, p2)
	winner := func() IntStack {
		if len(outcome1) == 0 {
			return IntStack(outcome2)
		}
		return IntStack(outcome1)
	}()

	score := 0

	l := len(winner)

	for i := 1; i <= l; i++ {
		score += i * winner.MustPop()
	}

	return score
}

// SolvePart2 ..
func SolvePart2(p1, p2 IntQueue) (int, error) {
	player, outcome1, outcome2, err := playRecursiveGame(p1, p2)

	if err != nil {
		return 0, err
	}

	winner := func() IntStack {
		switch player {
		case 1:
			return IntStack(outcome1)
		case 2:
			return IntStack(outcome2)
		default:
			return nil
		}
	}()

	score := 0

	l := len(winner)

	for i := 1; i <= l; i++ {
		score += i * winner.MustPop()
	}

	return score, nil
}

// IntQueue ..
type IntQueue []int

// Push ..
func (q *IntQueue) Push(n int) {
	*q = append(*q, n)
}

// Pop ..
func (q *IntQueue) Pop() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	}
	out := (*q)[0]
	*q = (*q)[1:]
	return out, true
}

// MustPop ..
func (q *IntQueue) MustPop() int {
	if len(*q) == 0 {
		panic("nothing to pop")
	}
	out := (*q)[0]
	*q = (*q)[1:]
	return out
}

// IntStack ..
type IntStack []int

// Push ..
func (s *IntStack) Push(n int) {
	*s = append(*s, n)
}

// Pop ..
func (s *IntStack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	out := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return out, true
}

// MustPop ..
func (s *IntStack) MustPop() int {
	if len(*s) == 0 {
		panic("nothing to pop")
	}
	out := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return out
}
