package rand

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestRangeInt validates the RangeInt function
func TestRangeInt(t *testing.T) {
	// Test edge cases
	assert.Equal(t, 0, RangeInt(0, 0), "RangeInt(0, 0) should return 0")
	assert.Equal(t, 1, RangeInt(1, 1), "RangeInt(1, 1) should return 1")
	assert.Equal(t, 0, RangeInt(10, 5), "RangeInt(10, 5) should return 0 for invalid range")

	// Test single value range
	assert.Equal(t, 5, RangeInt(5, 6), "RangeInt(5, 6) should return 5")

	// Test various ranges
	testCases := map[int]int{
		100:     200,
		101:     202,
		1010:    2020,
		10101:   20202,
		1010101: 2020202,
		1024:    4096,
		1024000: 2048000,
	}

	for min, max := range testCases {
		for i := 0; i < 1000; i++ {
			n := RangeInt(min, max)
			assert.GreaterOrEqual(t, n, min, "RangeInt(%d, %d) returned %d, should be >= %d", min, max, n, min)
			assert.Less(t, n, max, "RangeInt(%d, %d) returned %d, should be < %d", min, max, n, max)
		}
	}

	// Test large range
	for i := 0; i < 1000; i++ {
		n := RangeInt(0, math.MaxInt32)
		assert.GreaterOrEqual(t, n, 0)
		assert.Less(t, n, math.MaxInt32)
	}
}

// TestRangeIntSafe validates the RangeIntSafe function
func TestRangeIntSafe(t *testing.T) {
	// Test valid ranges
	n, err := RangeIntSafe(10, 20)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, n, 10)
	assert.Less(t, n, 20)

	// Test invalid ranges
	_, err = RangeIntSafe(20, 10)
	assert.ErrorIs(t, err, ErrInvalidRange)

	// Test equal values - should return the value itself
	n, err = RangeIntSafe(10, 10)
	require.NoError(t, err)
	assert.Equal(t, 10, n)

	n, err = RangeIntSafe(0, 0)
	require.NoError(t, err)
	assert.Equal(t, 0, n)
}

// TestRangeInt64 validates the RangeInt64 function
func TestRangeInt64(t *testing.T) {
	// Test edge cases
	assert.Equal(t, int64(0), RangeInt64(0, 0), "RangeInt64(0, 0) should return 0")
	assert.Equal(t, int64(1), RangeInt64(1, 1), "RangeInt64(1, 1) should return 1")
	assert.Equal(t, int64(0), RangeInt64(10, 5), "RangeInt64(10, 5) should return 0 for invalid range")

	// Test various ranges
	testCases := map[int64]int64{
		100:     200,
		101:     202,
		1010:    2020,
		10101:   20202,
		1010101: 2020202,
		1024:    4096,
		1024000: 2048000,
	}

	for min, max := range testCases {
		for i := 0; i < 1000; i++ {
			n := RangeInt64(min, max)
			assert.GreaterOrEqual(t, n, min, "RangeInt64(%d, %d) returned %d, should be >= %d", min, max, n, min)
			assert.Less(t, n, max, "RangeInt64(%d, %d) returned %d, should be < %d", min, max, n, max)
		}
	}

	// Test large range
	for i := 0; i < 1000; i++ {
		n := RangeInt64(0, math.MaxInt64)
		assert.GreaterOrEqual(t, n, int64(0))
		assert.Less(t, n, int64(math.MaxInt64))
	}
}

// TestRangeInt64Safe validates the RangeInt64Safe function
func TestRangeInt64Safe(t *testing.T) {
	// Test valid ranges
	n, err := RangeInt64Safe(1000, 2000)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, n, int64(1000))
	assert.Less(t, n, int64(2000))

	// Test invalid ranges
	_, err = RangeInt64Safe(2000, 1000)
	assert.ErrorIs(t, err, ErrInvalidRange)

	// Test equal values - should return the value itself
	n, err = RangeInt64Safe(1000, 1000)
	require.NoError(t, err)
	assert.Equal(t, int64(1000), n)
}

// TestRangeUint32 validates the RangeUint32 function
func TestRangeUint32(t *testing.T) {
	// Test edge cases
	assert.Equal(t, uint32(0), RangeUint32(0, 0), "RangeUint32(0, 0) should return 0")
	assert.Equal(t, uint32(1), RangeUint32(1, 1), "RangeUint32(1, 1) should return 1")
	assert.Equal(t, uint32(0), RangeUint32(10, 5), "RangeUint32(10, 5) should return 0 for invalid range")

	// Test various ranges
	testCases := map[uint32]uint32{
		100:  200,
		1000: 2000,
		1024: 4096,
	}

	for min, max := range testCases {
		for i := 0; i < 1000; i++ {
			n := RangeUint32(min, max)
			assert.GreaterOrEqual(t, n, min, "RangeUint32(%d, %d) returned %d, should be >= %d", min, max, n, min)
			assert.Less(t, n, max, "RangeUint32(%d, %d) returned %d, should be < %d", min, max, n, max)
		}
	}
}

// TestRangeUint64 validates the RangeUint64 function
func TestRangeUint64(t *testing.T) {
	// Test edge cases
	assert.Equal(t, uint64(0), RangeUint64(0, 0), "RangeUint64(0, 0) should return 0")
	assert.Equal(t, uint64(1), RangeUint64(1, 1), "RangeUint64(1, 1) should return 1")
	assert.Equal(t, uint64(0), RangeUint64(10, 5), "RangeUint64(10, 5) should return 0 for invalid range")

	// Test various ranges
	testCases := map[uint64]uint64{
		100:  200,
		1000: 2000,
		1024: 4096,
	}

	for min, max := range testCases {
		for i := 0; i < 1000; i++ {
			n := RangeUint64(min, max)
			assert.GreaterOrEqual(t, n, min, "RangeUint64(%d, %d) returned %d, should be >= %d", min, max, n, min)
			assert.Less(t, n, max, "RangeUint64(%d, %d) returned %d, should be < %d", min, max, n, max)
		}
	}
}

// TestRangeDistribution validates the distribution of range functions
func TestRangeDistribution(t *testing.T) {
	const iterations = 100000
	const min, max = 0, 100
	const buckets = 10
	const bucketSize = (max - min) / buckets

	counts := make([]int, buckets)

	for i := 0; i < iterations; i++ {
		n := RangeInt(min, max)
		bucket := (n - min) / bucketSize
		if bucket >= buckets {
			bucket = buckets - 1
		}
		counts[bucket]++
	}

	// Each bucket should have roughly iterations/buckets values
	expected := iterations / buckets
	tolerance := expected / 5 // 20% tolerance

	for i, count := range counts {
		assert.InDelta(t, expected, count, float64(tolerance),
			"Range distribution bucket %d has count %d, expected ~%d", i, count, expected)
	}
}

// TestRangeConcurrency validates thread safety of range functions
func TestRangeConcurrency(t *testing.T) {
	const goroutines = 50
	const iterations = 1000
	const min, max = 100, 1000

	results := make(chan int, goroutines*iterations)

	// Launch multiple goroutines
	for g := 0; g < goroutines; g++ {
		go func() {
			for i := 0; i < iterations; i++ {
				results <- RangeInt(min, max)
			}
		}()
	}

	// Collect all results
	for i := 0; i < goroutines*iterations; i++ {
		n := <-results

		// Basic bounds check
		require.GreaterOrEqual(t, n, min, "Concurrent RangeInt returned value out of bounds")
		require.Less(t, n, max, "Concurrent RangeInt returned value out of bounds")
	}
}

// BenchmarkRangeInt benchmarks the RangeInt function
func BenchmarkRangeInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = RangeInt(0, 1000)
	}
}

// BenchmarkRangeIntSafe benchmarks the RangeIntSafe function
func BenchmarkRangeIntSafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = RangeIntSafe(0, 1000)
	}
}

// BenchmarkRangeInt64 benchmarks the RangeInt64 function
func BenchmarkRangeInt64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = RangeInt64(0, 1000)
	}
}

// BenchmarkRangeInt64Safe benchmarks the RangeInt64Safe function
func BenchmarkRangeInt64Safe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = RangeInt64Safe(0, 1000)
	}
}
