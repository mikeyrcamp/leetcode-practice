package shortestcompletingword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shortestCompletingWord(t *testing.T) {
	assert.Equal(t, "steps", shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
	assert.Equal(t, "pest", shortestCompletingWord("1s3 456", []string{"looks", "pest", "stew", "show"}))
	assert.Equal(t, "according", shortestCompletingWord("GrC8950", []string{"measure", "other", "every", "base", "according", "level", "meeting", "none", "marriage", "rest"}))
}
