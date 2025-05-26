package rand

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt32(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := Int32()
		assert.LessOrEqual(t, n, int32(math.MaxInt32))
		assert.Less(t, int32(0), n)
	}
}

func TestInt64(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := Int64()
		assert.LessOrEqual(t, n, int64(math.MaxInt64))
		assert.Less(t, int64(0), n)
	}
}

func TestInt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := Int()
		assert.LessOrEqual(t, n, math.MaxInt64)
		assert.Less(t, 0, n)
	}
}

func TestUint32(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := Uint32()
		assert.LessOrEqual(t, n, uint32(math.MaxUint32))
		assert.Less(t, uint32(0), n)
	}
}

func TestUint64(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := Uint64()
		assert.LessOrEqual(t, n, uint64(math.MaxUint64))
		assert.Less(t, uint64(0), n)
	}
}

func TestUint(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := Uint()
		assert.LessOrEqual(t, n, uint(math.MaxUint))
		assert.Less(t, uint(0), n)
	}
}
