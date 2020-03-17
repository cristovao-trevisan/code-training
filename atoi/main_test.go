package atoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	assert.Equal(t, +1234, myAtoi("+1234"))
	assert.Equal(t, -1234, myAtoi("  -1234"))
	assert.Equal(t, 0, myAtoi("+-2"))
	assert.Equal(t, 1234, myAtoi(" 1234 with words"))
	assert.Equal(t, 0, myAtoi("words and 1234"))
	assert.Equal(t, -2147483648, myAtoi("-91283472332"))
	assert.Equal(t, 2147483647, myAtoi("91283472332"))
	assert.Equal(t, 0, myAtoi("- 1243"))
	assert.Equal(t, 0, myAtoi("0-1"))
	assert.Equal(t, -5, myAtoi("-5-"))
	assert.Equal(t, maxInt, myAtoi("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"))
	assert.Equal(t, 12345678, myAtoi("  0000000000012345678"))
}
