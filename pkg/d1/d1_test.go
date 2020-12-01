package d1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testIn string = `1721
979
366
299
675
1456`

func TestPart1(t *testing.T) {
	parsed, err := ParseFile([]byte(testIn))
	assert.NoError(t, err)
	t.Log(parsed)

	out := SolvePart1(parsed)

	assert.Equal(t, out, 514579)
}

func TestPart2(t *testing.T) {
	parsed, err := ParseFile([]byte(testIn))
	assert.NoError(t, err)
	t.Log(parsed)

	out := SolvePart2(parsed)

	assert.Equal(t, out, 241861950)
}
