package d3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func TestDecode(t *testing.T) {
	out := ParseInput([]byte(testInput))
	assert := assert.New(t)
	assert.True(out[0](2))
	assert.True(out[0](3))
	assert.False(out[0](4))
	assert.False(out[0](12))
	assert.True(out[0](13))
	assert.True(out[0](14))
	assert.True(out[10](1))
}

func TestSolvePart1(t *testing.T) {
	theForest := ParseInput([]byte(testInput))
	sol := SolvePart1(theForest)

	assert.Equal(t, 7, sol)
}

func TestSolvePart2(t *testing.T) {
	theForest := ParseInput([]byte(testInput))
	sol := SolvePart2(theForest)

	assert.Equal(t, 336, sol)
}
