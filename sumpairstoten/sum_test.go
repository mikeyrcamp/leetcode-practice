package sumpairstoten

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairSumIsTen(t *testing.T) {
	assert.Equal(t, []int{1, 9}, PairSumIsTen([]int{1, 2, 5, 6, 9, 5}))
	assert.Equal(t, []int{-2, 12}, PairSumIsTen([]int{1, -2, 5, 12, 6, 8}))
}
