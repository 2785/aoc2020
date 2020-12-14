package d14

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApplyMask(t *testing.T) {
	var mask Mask = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	assert.Equal(t, "000000000000000000000000000001001001", mask.Apply(11))
	assert.Equal(t, "000000000000000000000000000001100101", mask.Apply(101))
	assert.Equal(t, "000000000000000000000000000001000000", mask.Apply(0))
}

var testInput1 string = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

func TestParsInput(t *testing.T) {
	in, _, err := ParseInput([]byte(testInput1))
	assert.NoError(t, err)
	assert.Equal(t, 101, in["XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"][7])
}

func TestSolvePart1(t *testing.T) {
	in, order, err := ParseInput([]byte(testInput1))
	require.NoError(t, err)

	p1, err := SolvePart1(in, order)
	assert.NoError(t, err)
	assert.Equal(t, 165, p1)

}
