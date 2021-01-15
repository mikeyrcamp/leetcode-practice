package stringdivisible

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_divisible(t *testing.T) {
	assert.Equal(t, 2, divisible("abababababab", "ababab"))
	assert.Equal(t, 3, divisible("bcdbcdbcdbcdbcdbcdbcdbcdbcdbcdbcdbcdbcdbcdbcd", "bcdbcdbcd"))
	assert.Equal(t, 3, divisible("abcabc", "abc"))
	assert.Equal(t, -1, divisible("abcabcd", "a"))
	assert.Equal(t, 1, divisible("aaaaaaa", "a"))
	assert.Equal(t, -1, divisible("", ""))
}
