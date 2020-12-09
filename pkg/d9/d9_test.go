package d9

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsValidSum(t *testing.T) {
	seq := []int{35, 20, 15, 25, 47}
	assert.True(t, isValidSum(seq, 40))
	seq = []int{20, 15, 25, 47, 40}
	assert.True(t, isValidSum(seq, 62))
	seq = []int{95, 102, 117, 150, 182}
	assert.False(t, isValidSum(seq, 127))

}

var testInput string = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestSolvePart1(t *testing.T) {
	nums, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	sol, err := SolvePart1(nums, 5)
	assert.NoError(t, err)
	assert.Equal(t, 127, sol)
}

func TestSolvePart2(t *testing.T) {
	nums, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	sol, err := SolvePart2(nums, 127)
	assert.NoError(t, err)
	assert.Equal(t, 62, sol)
}
