package d7

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testInput string = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

func TestDecodeOne(t *testing.T) {
	assert := assert.New(t)
	name, rules, err := DecodeOne("light red bags contain 1 bright white bag, 2 muted yellow bags.")
	assert.NoError(err)
	assert.Equal("light red", name)
	assert.Equal(map[string]int{"bright white": 1, "muted yellow": 2}, rules)
}

func TestParseInput(t *testing.T) {
	assert := assert.New(t)
	book, err := ParseInput([]byte(testInput))
	assert.NoError(err)
	assert.Equal(map[string]int{"bright white": 1, "muted yellow": 2}, book["light red"])
	assert.Equal(map[string]int{"shiny gold": 2, "faded blue": 9}, book["muted yellow"])
}

func TestSolvePart1(t *testing.T) {
	assert := assert.New(t)

	book, err := ParseInput([]byte(testInput))
	require.NoError(t, err)

	count := SolvePart1(book)

	assert.Equal(4, count)

}

var testInput2 string = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`

func TestSolvePart2(t *testing.T) {
	book1, err := ParseInput([]byte(testInput))
	require.NoError(t, err)
	assert.Equal(t, 32, SolvePart2(book1))
	book2, err := ParseInput([]byte(testInput2))
	require.NoError(t, err)
	assert.Equal(t, 126, SolvePart2(book2))
}
