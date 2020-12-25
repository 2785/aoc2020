package d25

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLoop(t *testing.T) {
	r1, err := getLoopCountFromPublicKey(5764801)
	assert.NoError(t, err)
	assert.Equal(t, 8, r1)

	r2, err := getLoopCountFromPublicKey(17807724)
	assert.NoError(t, err)
	assert.Equal(t, 11, r2)
}

func TestSolvePart1(t *testing.T) {
	r, err := SolvePart1(5764801, 17807724)
	assert.NoError(t, err)
	assert.Equal(t, 14897079, r)
}
