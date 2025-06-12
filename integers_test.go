package rand

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestInt32 validates the Int32 function
func TestInt32(t *testing.T) {
	const iterations = 10000

	for i := 0; i < iterations; i++ {
		n := Int32()

		// Test bounds
		assert.GreaterOrEqual(t, n, int32(0), "Int32() should return non-negative values")
		assert.Less(t, n, int32(math.MaxInt32), "Int32() should be less than MaxInt32")
	}

	// Test uniqueness with multiple calls
	values := make(map[int32]bool, 1000)
	duplicates := 0

	for i := 0; i < 1000; i++ {
		n := Int32()
		if values[n] {
			duplicates++
		}
		values[n] = true
	}

	// There should be very few duplicates in 1000 random int32 values
	assert.Less(t, duplicates, 10, "Too many duplicates in Int32 generation")
}

// TestInt64 validates the Int64 function
func TestInt64(t *testing.T) {
	const iterations = 10000

	for i := 0; i < iterations; i++ {
		n := Int64()

		// Test bounds
		assert.GreaterOrEqual(t, n, int64(0), "Int64() should return non-negative values")
		assert.Less(t, n, int64(math.MaxInt64), "Int64() should be less than MaxInt64")
	}

	// Test uniqueness
	values := make(map[int64]bool, 1000)
	duplicates := 0

	for i := 0; i < 1000; i++ {
		n := Int64()
		if values[n] {
			duplicates++
		}
		values[n] = true
	}

	// There should be very few duplicates in 1000 random int64 values
	assert.Less(t, duplicates, 10, "Too many duplicates in Int64 generation")
}

// TestInt validates the Int function
func TestInt(t *testing.T) {
	const iterations = 10000

	for i := 0; i < iterations; i++ {
		n := Int()

		// Test bounds
		assert.GreaterOrEqual(t, n, 0, "Int() should return non-negative values")
		assert.Less(t, n, math.MaxInt, "Int() should be less than MaxInt")
	}
}

// TestUint32 validates the Uint32 function
func TestUint32(t *testing.T) {
	const iterations = 10000

	for i := 0; i < iterations; i++ {
		n := Uint32()

		// Test bounds
		assert.GreaterOrEqual(t, n, uint32(0), "Uint32() should return non-negative values")
		assert.Less(t, n, uint32(math.MaxUint32), "Uint32() should be less than MaxUint32")
	}

	// Test uniqueness
	values := make(map[uint32]bool, 1000)
	duplicates := 0

	for i := 0; i < 1000; i++ {
		n := Uint32()
		if values[n] {
			duplicates++
		}
		values[n] = true
	}

	assert.Less(t, duplicates, 10, "Too many duplicates in Uint32 generation")
}

// TestUint64 validates the Uint64 function
func TestUint64(t *testing.T) {
	const iterations = 10000

	for i := 0; i < iterations; i++ {
		n := Uint64()

		// Test bounds
		assert.GreaterOrEqual(t, n, uint64(0), "Uint64() should return non-negative values")
		assert.Less(t, n, uint64(math.MaxUint64), "Uint64() should be less than MaxUint64")
	}

	// Test uniqueness
	values := make(map[uint64]bool, 1000)
	duplicates := 0

	for i := 0; i < 1000; i++ {
		n := Uint64()
		if values[n] {
			duplicates++
		}
		values[n] = true
	}

	assert.Less(t, duplicates, 10, "Too many duplicates in Uint64 generation")
}

// TestUint validates the Uint function
func TestUint(t *testing.T) {
	const iterations = 10000

	for i := 0; i < iterations; i++ {
		n := Uint()

		// Test bounds
		assert.GreaterOrEqual(t, n, uint(0), "Uint() should return non-negative values")
		assert.Less(t, n, uint(math.MaxUint), "Uint() should be less than MaxUint")
	}
}

// BenchmarkInt32 benchmarks the Int32 function
func BenchmarkInt32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Int32()
	}
}

// BenchmarkInt64 benchmarks the Int64 function
func BenchmarkInt64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Int64()
	}
}

// BenchmarkUint32 benchmarks the Uint32 function
func BenchmarkUint32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Uint32()
	}
}

// BenchmarkUint64 benchmarks the Uint64 function
func BenchmarkUint64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Uint64()
	}
}

// TestDistribution validates that the generated numbers have reasonable distribution
func TestDistribution(t *testing.T) {
	const iterations = 100000
	const buckets = 10

	// Test Int32 distribution
	int32Counts := make([]int, buckets)
	bucketSize := math.MaxInt32 / buckets

	for i := 0; i < iterations; i++ {
		n := Int32()
		bucket := int(n) / bucketSize
		if bucket >= buckets {
			bucket = buckets - 1 // Handle edge case
		}
		int32Counts[bucket]++
	}

	// Each bucket should have roughly iterations/buckets values
	expected := iterations / buckets
	tolerance := expected / 5 // 20% tolerance

	for i, count := range int32Counts {
		assert.InDelta(t, expected, count, float64(tolerance),
			"Int32 distribution bucket %d has count %d, expected ~%d", i, count, expected)
	}
}

// TestConcurrency validates thread safety
func TestConcurrency(t *testing.T) {
	const goroutines = 100
	const iterations = 1000

	results := make(chan int32, goroutines*iterations)

	// Launch multiple goroutines
	for g := 0; g < goroutines; g++ {
		go func() {
			for i := 0; i < iterations; i++ {
				results <- Int32()
			}
		}()
	}

	// Collect all results
	values := make(map[int32]bool)
	for i := 0; i < goroutines*iterations; i++ {
		n := <-results
		values[n] = true

		// Basic bounds check
		require.GreaterOrEqual(t, n, int32(0))
		require.Less(t, n, int32(math.MaxInt32))
	}

	// Should have generated many unique values
	assert.Greater(t, len(values), goroutines*iterations/2,
		"Concurrent generation should produce many unique values")
}
