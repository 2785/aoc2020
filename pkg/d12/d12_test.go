package d12

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testInput string = `F10
N3
F7
R90
F11`

func TestParseInput(t *testing.T) {
	_, err := ParseInput([]byte(testInput))
	assert.NoError(t, err)
}

func TestSolvePart1(t *testing.T) {
	parsed, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	assert.Equal(t, 25, SolvePart1(parsed))
}

func TestSolvePart2(t *testing.T) {
	parsed, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	assert.Equal(t, 286, SolvePart2(parsed))
}
