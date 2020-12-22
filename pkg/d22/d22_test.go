package d22

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testInput string = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

func TestSolvePart1(t *testing.T) {
	p1, p2, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	assert.Equal(t, 306, SolvePart1(p1, p2))
}

func TestSolvePart2(t *testing.T) {
	p1, p2, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	sol, err := SolvePart2(p1, p2)
	assert.NoError(t, err)
	assert.Equal(t, 291, sol)
}
