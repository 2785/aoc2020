package d10

import (
	"testing"

	"github.com/2785/aoc2020/pkg/input"
	"github.com/stretchr/testify/assert"
)

var testInput1 string = `16
10
15
5
1
11
7
19
6
12
4`

var testInput2 string = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 35, SolvePart1(input.MustParseInt([]byte(testInput1))))
}

func TestSolvePart2(t *testing.T) {
	// assert.Equal(t, 8, SolvePart2(input.MustParseInt([]byte(testInput1))))
	assert.Equal(t, 19208, SolvePart2(input.MustParseInt([]byte(testInput2))))
}

func TestPermOnes(t *testing.T) {
	assert.Equal(t, 1, confForOneChain(1))
	assert.Equal(t, 2, confForOneChain(2))
	assert.Equal(t, 4, confForOneChain(3))
	assert.Equal(t, 7, confForOneChain(4))
}
