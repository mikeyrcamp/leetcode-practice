package bubblesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	v := []int{7, 2, 6, 5, 4, 3, 1}
	SortInts(v)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, v)
}
