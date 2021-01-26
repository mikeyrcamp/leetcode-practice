package secondlargest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecondLargest(t *testing.T) {
	assert.Equal(t, 4, SecondLargest([]int{4, 3, 5, 1, 2}))
	assert.Equal(t, -1, SecondLargest([]int{1}))
}
