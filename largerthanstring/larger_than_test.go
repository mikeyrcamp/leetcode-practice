package largerthanstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecondLargest(t *testing.T) {
	assert.True(t, LargerThan("1", ""))
	assert.False(t, LargerThan("", "123"))
	assert.True(t, LargerThan("1235", "1234"))
	assert.False(t, LargerThan("1235", "1235"))
	assert.False(t, LargerThan("1235", "1236"))
}
