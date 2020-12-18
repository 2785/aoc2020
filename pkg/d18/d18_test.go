package d18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput string = `2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`

func TestEval(t *testing.T) {
	res, err := ModifiedShuntingYard(ParseInput([]byte(testInput))[0])
	assert.NoError(t, err)
	assert.Equal(t, 26, res)
}

func TestPart2(t *testing.T) {
	res, err := ModifiedShuntingYard2(ParseInput([]byte(testInput))[0])
	assert.NoError(t, err)
	assert.Equal(t, 46, res)
	res, err = ModifiedShuntingYard2(ParseInput([]byte(testInput))[1])
	assert.NoError(t, err)
	assert.Equal(t, 1445, res)
	res, err = ModifiedShuntingYard2(ParseInput([]byte(testInput))[2])
	assert.NoError(t, err)
	assert.Equal(t, 669060, res)
	res, err = ModifiedShuntingYard2(ParseInput([]byte(testInput))[3])
	assert.NoError(t, err)
	assert.Equal(t, 23340, res)
}
