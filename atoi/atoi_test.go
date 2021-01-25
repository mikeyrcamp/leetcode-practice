package atoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	v, err := Atoi("100")
	assert.NoError(t, err)
	assert.Equal(t, 100, v)

	v, err = Atoi("-100.00")
	assert.NoError(t, err)
	assert.Equal(t, -100, v)

	v, err = Atoi("--   ")
	assert.Error(t, err)
	assert.Equal(t, 0, v)

	v, err = Atoi("1233455")
	assert.NoError(t, err)
	assert.Equal(t, 1233455, v)

	v, err = Atoi("-1233455")
	assert.NoError(t, err)
	assert.Equal(t, -1233455, v)

	v, err = Atoi("123a45B5")
	assert.Error(t, err)
	assert.Equal(t, 0, v)
}
