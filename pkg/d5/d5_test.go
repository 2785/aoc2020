package d5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeOne(t *testing.T) {
	for s, ID := range map[string]int{
		"FBFBBFFRLR": 357,
		"BFFFBBFRRR": 567,
		"FFFBBBFRRR": 119,
		"BBFFBBFRLL": 820,
	} {
		res, err := DecodeOne(s)
		assert.NoError(t, err)
		assert.Equal(t, ID, res.ID())
	}
}

func TestParseInput(t *testing.T) {
	testInput := `FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`

	_, err := ParseInput([]byte(testInput))
	assert.NoError(t, err)
}

func TestSolvePart1(t *testing.T) {
	testInput := `FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`

	seats, err := ParseInput([]byte(testInput))
	assert.NoError(t, err)

	assert.Equal(t, 820, SolvePart1(seats))
}
