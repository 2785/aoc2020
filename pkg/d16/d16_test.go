package d16

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testInput string = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

func TestParseInput(t *testing.T) {
	rule, your, nearby, err := ParseInput([]byte(testInput))
	assert.NoError(t, err)
	assert.True(t, rule["class"](2))
	assert.False(t, rule["seat"](43))
	assert.Equal(t, []int{7, 1, 14}, your)
	assert.Equal(t, [][]int{{7, 3, 47}, {40, 4, 50}, {55, 2, 20}, {38, 6, 12}}, nearby)
}

func TestSolvePart1(t *testing.T) {
	rule, _, nearby, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	assert.Equal(t, 71, SolvePart1(nearby, rule))
}

func TestSolvePart2(t *testing.T) {
	rule, yours, nearby, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	sol, err := SolvePart2(yours, nearby, rule)
	assert.NoError(t, err)
	assert.Equal(t, 71, sol)
}
