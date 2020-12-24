package d24

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testInput string = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

func TestParseInput(t *testing.T) {
	parsed, err := ParseInput([]byte(testInput))
	assert.NoError(t, err)
	assert.Equal(t, se, parsed[0][0])
}

func TestSolvePart1(t *testing.T) {
	parsed, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	assert.Equal(t, 10, SolvePart1(parsed))
}

func TestSolvePart2(t *testing.T) {
	parsed, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	assert.Equal(t, 2208, SolvePart2(parsed))
}
