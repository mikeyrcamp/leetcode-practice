package countnegintsquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountNegativeInts(t *testing.T) {
	board := [][]int{
		{-4, -3, -2, 1},
		{-2, -2, 1, 2},
		{-1, 1, 2, 3},
		{1, 2, 4, 5},
	}
	assert.Equal(t, 6, CountNegativeIntsNaive(board))
	assert.Equal(t, 6, CountNegativeInts(board))

	board = [][]int{
		{-2, 0},
		{-1, 0},
	}
	assert.Equal(t, 2, CountNegativeIntsNaive(board))
	assert.Equal(t, 2, CountNegativeInts(board))
}
