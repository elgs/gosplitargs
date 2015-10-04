// gosplitargs
package gosplitargs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitArgs(t *testing.T) {
	testSpace(t, "I said 'I am sorry.', and he said \"it doesn't matter.\"")
	testSpace(t, "I said \"I am sorry.\", and he said \"it doesn't matter.\"")
	testSpace(t, `I said "I am sorry.", and he said "it doesn't matter."`)
	testSpace(t, `I said 'I am sorry.', and he said "it doesn't matter."`)
	testSemicolon(t, "SET @safe_uuid := UUID();INSERT INTO sys_user SET ID=@safe_uuid, CODE='1;2', EMAIL=?, PASSWORD=ENCRYPT(?, CONCAT('$6$', SUBSTRING(SHA(RAND()), -16)));")
}

func testSpace(t *testing.T, i string) {
	o, err := SplitArgs(i, "", false)
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

func testSemicolon(t *testing.T, i string) {
	o, err := SplitArgs(i, ";", true)
	assert.Nil(t, err)
	assert.Equal(t, o[0], "SET @safe_uuid := UUID()")
	assert.Equal(t, o[1], "INSERT INTO sys_user SET ID=@safe_uuid, CODE='1;2', EMAIL=?, PASSWORD=ENCRYPT(?, CONCAT('$6$', SUBSTRING(SHA(RAND()), -16)))")
}
