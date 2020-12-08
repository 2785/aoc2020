package d8

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeOne(t *testing.T) {
	testString := "acc +3"
	ins, err := DecodeOne(testString)
	assert.NoError(t, err)
	assert.Equal(t, acc, ins.Op)
	assert.Equal(t, 3, ins.Val)

	testString = "acc -5"
	ins, err = DecodeOne(testString)
	assert.NoError(t, err)
	assert.Equal(t, acc, ins.Op)
	assert.Equal(t, -5, ins.Val)
}

var testInput string = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestParseInput(t *testing.T) {
	set, err := ParseInput([]byte(testInput))
	assert.NoError(t, err)
	assert.Len(t, set, 9)
	assert.Equal(t, acc, set[5].Op)
}

func TestSolvePart1(t *testing.T) {
	set, err := ParseInput([]byte(testInput))
	require.NoError(t, err)

	assert.Equal(t, 5, SolvePart1(set))
}

func TestSolvePart2(t *testing.T) {
	set, err := ParseInput([]byte(testInput))
	require.NoError(t, err)

	sol, err := SolvePart2(set)

	assert.NoError(t, err)
	assert.Equal(t, 8, sol)
}
