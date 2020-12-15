package d15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayTheGame(t *testing.T) {
	assert.Equal(t, 436, playTheGame([]int{0, 3, 6}, 2020))
	assert.Equal(t, 1, playTheGame([]int{1, 3, 2}, 2020))
	assert.Equal(t, 1836, playTheGame([]int{3, 1, 2}, 2020))
}
