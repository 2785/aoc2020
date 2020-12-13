package d13

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testInput string = `939
7,13,x,x,59,x,31,19`

func TestParseInput(t *testing.T) {
	schedule, num, err := ParseInput([]byte(testInput))
	assert.NoError(t, err)
	assert.Equal(t, []int{7, 13, 0, 0, 59, 0, 31, 19}, schedule)
	assert.Equal(t, 939, num)
}

func TestSolvePart1(t *testing.T) {
	schedule, num, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	assert.Equal(t, 295, SolvePart1(schedule, num))
}

func TestSolvePart2(t *testing.T) {
	s1 := []int{7, 13, 0, 0, 59, 0, 31, 19}
	sol1, err := SolvePart2(s1)
	assert.NoError(t, err)
	assert.Equal(t, 1068781, sol1)

	s2 := []int{17, 0, 13, 19}
	sol2, err := SolvePart2(s2)
	assert.NoError(t, err)
	assert.Equal(t, 3417, sol2)

	s3 := []int{1789, 37, 47, 1889}
	sol3, err := SolvePart2(s3)
	assert.NoError(t, err)
	assert.Equal(t, 1202161486, sol3)

	s4 := []int{1789, 0, 0, 1889}
	sol4, err := SolvePart2(s4)
	assert.NoError(t, err)
	assert.Equal(t, 2467031, sol4)
}

func TestSolvePartSchedule(t *testing.T) {
	s1 := []int{7, 13, 0, 0, 59, 0, 31, 19}
	curr := 59 - 4
	sol1, err := solvePartSchedule(s1, func() int {
		old := curr
		curr = curr + 59
		return old
	})
	assert.NoError(t, err)
	assert.Equal(t, 1068781, sol1)
}
