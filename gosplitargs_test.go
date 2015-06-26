// gosplitargs
package gosplitargs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitArgs(t *testing.T) {
	test(t, "I said 'I am sorry.', and he said \"it doesn't matter.\"")
	test(t, "I said \"I am sorry.\", and he said \"it doesn't matter.\"")
	test(t, `I said "I am sorry.", and he said "it doesn't matter."`)
	test(t, `I said 'I am sorry.', and he said "it doesn't matter."`)
}

func test(t *testing.T, i string) {
	o, err := SplitArgs(i, "")
	assert.Nil(t, err)
	assert.Equal(t, len(o), 7)
	assert.Equal(t, o[0], "I")
	assert.Equal(t, o[1], "said")
	assert.Equal(t, o[2], "I am sorry.,")
	assert.Equal(t, o[3], "and")
	assert.Equal(t, o[4], "he")
	assert.Equal(t, o[5], "said")
	assert.Equal(t, o[6], "it doesn't matter.")
}
