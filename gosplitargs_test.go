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
	testEmpty(t, `,,,`)
}

func testSpace(t *testing.T, i string) {
	o, err := SplitArgs(i, "", false)
	assert.Nil(t, err)
	assert.Equal(t, 7, len(o))
	assert.Equal(t, "I", o[0])
	assert.Equal(t, "said", o[1])
	assert.Equal(t, "I am sorry.,", o[2])
	assert.Equal(t, "and", o[3])
	assert.Equal(t, "he", o[4])
	assert.Equal(t, "said", o[5])
	assert.Equal(t, "it doesn't matter.", o[6])
}

func testSemicolon(t *testing.T, i string) {
	o, err := SplitArgs(i, ";", true)
	assert.Nil(t, err)
	assert.Equal(t, "SET @safe_uuid := UUID()", o[0])
	assert.Equal(t, "INSERT INTO sys_user SET ID=@safe_uuid, CODE='1;2', EMAIL=?, PASSWORD=ENCRYPT(?, CONCAT('$6$', SUBSTRING(SHA(RAND()), -16)))", o[1])
}

func testEmpty(t *testing.T, i string) {
	o, err := SplitArgs(i, ",", true)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(o))

}
