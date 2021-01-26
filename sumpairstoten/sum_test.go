package sumpairstoten

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairSumIsTen(t *testing.T) {
	assert.Equal(t, []int{1, 9}, PairSumIsTen([]int{1, 2, 5, 6, 9, 5}))
	assert.Equal(t, []int{-2, 12}, PairSumIsTen([]int{1, -2, 5, 12, 6, 8}))
	assert.Nil(t, PairSumIsTen(nil))
	assert.Nil(t, PairSumIsTen([]int{1, 4, -5, 3, 4}))
}
