package itoa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_itoa(t *testing.T) {
	assert.Equal(t, "100", itoa(100))
	assert.Equal(t, "-101010", itoa(-101010))
	assert.Equal(t, "123546789", itoa(123546789))
	assert.Equal(t, "399910", itoa(399910))
}
