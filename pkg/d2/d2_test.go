package d2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockIn = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

func TestParseInput(t *testing.T) {
	parsed, err := ParseInput([]byte(mockIn))
	assert.NoError(t, err)
	assert.Equal(t, []Record{
		{Letter: "a", Content: "abcde", Min: 1, Max: 3},
		{Letter: "b", Content: "cdefg", Min: 1, Max: 3},
		{Letter: "c", Content: "ccccccccc", Min: 2, Max: 9},
	}, parsed)
}

func TestSolvePart1(t *testing.T) {
	input := []Record{
		{Letter: "a", Content: "abcde", Min: 1, Max: 3},
		{Letter: "b", Content: "cdefg", Min: 1, Max: 3},
		{Letter: "c", Content: "ccccccccc", Min: 2, Max: 9},
	}

	res := SolvePart1(input)
	assert.Equal(t, 2, res)
}

func TestSolvePart2(t *testing.T) {
	input := []Record{
		{Letter: "a", Content: "abcde", Min: 1, Max: 3},
		{Letter: "b", Content: "cdefg", Min: 1, Max: 3},
		{Letter: "c", Content: "ccccccccc", Min: 2, Max: 9},
	}

	res := SolvePart2(input)
	assert.Equal(t, 1, res)
}
