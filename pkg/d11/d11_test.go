package d11

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseInput(t *testing.T) {
	testInput := `L.L..
LL..L
.....`

	out, err := ParseInput([]byte(testInput))
	assert.NoError(t, err)
	assert.Equal(t, floor, out[1][2])
	assert.Equal(t, empty, out[0][2])
}

func TestIfShouldFlip(t *testing.T) {
	testInput := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.##
..L.L...##
LLLLLLLL#L
L.LLLLLL.L
L.LLLLL.LL`

	assert := assert.New(t)

	parsed, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	t1, err := parsed.ShouldFlip(6, 9)
	assert.NoError(err)
	assert.True(t1)
	t2, err := parsed.ShouldFlip(7, 8)
	assert.NoError(err)
	assert.False(t2)
	t3, err := parsed.ShouldFlip(9, 9)
	assert.NoError(err)
	assert.True(t3)
}

var testInput1 string = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

func TestSolvePart1(t *testing.T) {
	parsed, err := ParseInput([]byte(testInput1))
	require.NoError(t, err)
	sol, err := SolvePart1(parsed)
	assert.NoError(t, err)
	assert.Equal(t, 37, sol)
}

func TestSolvePart2(t *testing.T) {
	parsed, err := ParseInput([]byte(testInput1))
	require.NoError(t, err)
	sol, err := SolvePart2(parsed)
	assert.NoError(t, err)
	assert.Equal(t, 26, sol)
}
