package d23

import (
	"container/ring"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlayRound(t *testing.T) {
	m := make([]*ring.Ring, 10)
	initialCups := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	game := ring.New(len(initialCups))
	for i := 0; i < game.Len(); i++ {
		game.Value = initialCups[i]
		m[initialCups[i]] = game
		game = game.Next()
	}

	game = playOneRound(game, 9, m)
	state := []int{}
	l := game.Len()
	for i := 0; i < l; i++ {
		state = append(state, game.Value.(int))
		game = game.Next()
	}

	assert.Equal(t, []int{2, 8, 9, 1, 5, 4, 6, 7, 3}, state)

}

var testInput string = "389125467"

func TestSolvePart1(t *testing.T) {
	parsed, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	assert.Equal(t, "67384529", SolvePart1(parsed))
}

func TestSolvePart2(t *testing.T) {
	parsed, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	assert.Equal(t, 149245887792, SolvePart2(parsed))
}
