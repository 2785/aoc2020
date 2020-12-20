package d20

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvertString(t *testing.T) {
	assert.Equal(t, "cba", invertString("abc"))
}

func TestTile(t *testing.T) {
	start := [][]bool{
		{true, false, false},
		{true, false, true},
		{false, false, false},
	}

	tile, err := NewTile(1, start)
	require.NoError(t, err)

	assert.Equal(t, "#..", tile.Edges[N])
	assert.Equal(t, ".#.", tile.Edges[E])

	tile1 := tile.rotateClockWise()
	assert.Equal(t, [][]bool{
		{false, true, true},
		{false, false, false},
		{false, true, false},
	}, tile1.Data)
	tile1 = tile1.rotateClockWise()
	assert.Equal(t, "#..", tile1.Edges[S])
	assert.Equal(t, []bool{false, false, true}, tile1.Data[2])

	shouldOverlap, ok := tile1.overlap("..#", false)
	assert.True(t, ok)
	assert.Equal(t, S, shouldOverlap)

	_, ok = tile1.overlap("#..", false)
	assert.False(t, ok)

	tile2 := tile.mirrorTopDown()
	assert.Equal(t, [][]bool{
		{false, false, false},
		{true, false, true},
		{true, false, false},
	}, tile2.Data)

	tile3 := tile.mirrorVertical()
	assert.Equal(t, [][]bool{
		{false, false, true},
		{true, false, true},
		{false, false, false},
	}, tile3.Data)
}

func TestParseInput(t *testing.T) {
	parsed, err := ParseInput([]byte(testInput))
	assert.NoError(t, err)
	assert.Equal(t, 9, len(parsed))
}

func TestSolvePart1(t *testing.T) {
	parsed, err := ParseInput([]byte(testInput))
	require.NoError(t, err)

	sol, err := SolvePart1(parsed)
	assert.NoError(t, err)
	assert.Equal(t, 20899048083289, sol)
}

func TestSolvePart2(t *testing.T) {
	parsed, err := ParseInput([]byte(testInput))
	require.NoError(t, err)

	sol, err := SolvePart2(parsed)
	assert.NoError(t, err)
	assert.Equal(t, 273, sol)
}

var testInput string = `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`
