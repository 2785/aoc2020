package d17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput string = `.#.
..#
###`

func TestParseInput(t *testing.T) {
	parsed := ParseInput([]byte(testInput))
	assert.Equal(t, map[int]map[int]map[int]bool{
		0: {
			0: {
				1: true,
			},
			1: {
				2: true,
			},
			2: {
				0: true,
				1: true,
				2: true,
			},
		},
	}, parsed)
}

func TestSolvePart1(t *testing.T) {
	parsed := ParseInput([]byte(testInput))
	assert.Equal(t, 112, SolvePart1(parsed))
}

func TestSolvePart2(t *testing.T) {
	parsed := ParseInput([]byte(testInput))
	assert.Equal(t, 848, SolvePart2(parsed))
}
