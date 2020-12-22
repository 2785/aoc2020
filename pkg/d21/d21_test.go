package d21

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testInput string = `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

func TestSolvePart1(t *testing.T) {
	parsed, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	sol, err := SolvePart1(parsed)
	assert.NoError(t, err)
	assert.Equal(t, 5, sol)
}

func TestSolvePart2(t *testing.T) {
	parsed, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	sol, err := SolvePart2(parsed)
	assert.NoError(t, err)
	assert.Equal(t, "mxmxvkd,sqjhc,fvjkl", sol)
}
