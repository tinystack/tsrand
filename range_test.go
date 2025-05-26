package rand

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeInt(t *testing.T) {
	assert.Equal(t, RangeInt(0, 0), 0)
	assert.Equal(t, RangeInt(1, 0), 0)
	assert.Equal(t, RangeInt(1, 1), 1)
	assert.Equal(t, RangeInt(1, 2), 1)

	var testData = map[int]int{
		100:     200,
		101:     202,
		1010:    2020,
		10101:   20202,
		1010101: 2020202,
		1024:    4096,
		1024000: 2048000,
	}
	for min, max := range testData {
		n := RangeInt(min, max)
		assert.Less(t, n, max)
		assert.LessOrEqual(t, min, n)
	}

	for i := 0; i < 1000; i++ {
		n := RangeInt(0, math.MaxInt)
		assert.Less(t, n, math.MaxInt)
		assert.LessOrEqual(t, 0, n)
	}
}

func TestRangeInt64(t *testing.T) {
	assert.Equal(t, RangeInt64(int64(0), int64(0)), int64(0))
	assert.Equal(t, RangeInt64(int64(1), int64(0)), int64(0))
	assert.Equal(t, RangeInt64(int64(1), int64(1)), int64(1))
	assert.Equal(t, RangeInt64(int64(1), int64(2)), int64(1))

	var testData = map[int64]int64{
		100:     200,
		101:     202,
		1010:    2020,
		10101:   20202,
		1010101: 2020202,
		1024:    4096,
		1024000: 2048000,
	}
	for min, max := range testData {
		n := RangeInt64(min, max)
		assert.Less(t, n, max)
		assert.LessOrEqual(t, min, n)
	}

	for i := 0; i < 1000; i++ {
		n := RangeInt64(int64(0), math.MaxInt64)
		assert.Less(t, n, int64(math.MaxInt64))
		assert.LessOrEqual(t, int64(0), n)
	}
}
