package d6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeOne(t *testing.T) {
	assert.Equal(t, Group{
		Count: 1,
		Questions: map[rune]int{
			'a': 1,
			'b': 1,
			'c': 1,
		},
	}, DecodeOne("abc"))

	assert.Equal(t, Group{
		Count: 3,
		Questions: map[rune]int{
			'a': 1,
			'b': 1,
			'c': 1,
		},
	}, DecodeOne("a\nb\nc"))

	assert.Equal(t, Group{
		Count: 2,
		Questions: map[rune]int{
			'a': 2,
			'b': 1,
			'c': 1,
		},
	}, DecodeOne("ab\nac"))
}

var testInput string = `abc

a
b
c

ab
ac

a
a
a
a

b`

func TestParseInput(t *testing.T) {
	grps := ParseInput([]byte(testInput))
	assert.Equal(t, []Group{
		{
			Count: 1,
			Questions: map[rune]int{
				'a': 1,
				'b': 1,
				'c': 1,
			},
		},
		{
			Count: 3,
			Questions: map[rune]int{
				'a': 1,
				'b': 1,
				'c': 1,
			},
		},
		{
			Count: 2,
			Questions: map[rune]int{
				'a': 2,
				'b': 1,
				'c': 1,
			},
		},
		{
			Count: 4,
			Questions: map[rune]int{
				'a': 4,
			},
		},
		{
			Count: 1,
			Questions: map[rune]int{
				'b': 1,
			},
		},
	}, grps)
}

func TestSolvePart1(t *testing.T) {
	grps := ParseInput([]byte(testInput))
	assert.Equal(t, 11, SolvePart1(grps))
}

func TestSolvePart2(t *testing.T) {
	grps := ParseInput([]byte(testInput))
	assert.Equal(t, 6, SolvePart2(grps))
}
