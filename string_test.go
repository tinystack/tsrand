package rand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert.Equal(t, 200, len(String(200)))
	assert.NotEqual(t, String(100), String(100))
}

func TestVisibleString(t *testing.T) {
	assert.Equal(t, 200, len(VisibleString(200)))

	for i := 0; i < 1000; i++ {
		assert.NotEqual(t, VisibleString(100), VisibleString(100))
	}

	assert.NotContains(t, VisibleString(1000), "1")
	assert.NotContains(t, VisibleString(1000), "i")
	assert.NotContains(t, VisibleString(1000), "I")
	assert.NotContains(t, VisibleString(1000), "l")
	assert.NotContains(t, VisibleString(1000), "0")
	assert.NotContains(t, VisibleString(1000), "O")
	assert.NotContains(t, VisibleString(1000), "o")
}

func TestUUID(t *testing.T) {
	assert.Equal(t, 36, len(UUID()))
	for i := 0; i < 1000; i++ {
		assert.NotEqual(t, UUID(), UUID())
	}
}
