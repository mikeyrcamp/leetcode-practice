package rooksaresafe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRooksAreSafe(t *testing.T) {
	board := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	assert.True(t, RooksAreSafe(board))

	board = [][]int{
		{1, 0, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	assert.False(t, RooksAreSafe(board))

	board = [][]int{
		{1, 0, 0},
		{0, 1, 1},
		{0, 0, 1},
	}
	assert.False(t, RooksAreSafe(board))
}
