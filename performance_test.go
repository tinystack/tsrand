package rand

import (
	"fmt"
	"testing"
	"time"
)

// TestPerformanceComparison compares performance between different approaches
func TestPerformanceComparison(t *testing.T) {
	const iterations = 100000

	// Test integer generation performance
	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = Int32()
	}
	int32Duration := time.Since(start)
	t.Logf("Int32 generation: %d operations in %v (%.2f ns/op)",
		iterations, int32Duration, float64(int32Duration.Nanoseconds())/float64(iterations))

	// Test string generation performance
	start = time.Now()
	for i := 0; i < iterations/10; i++ { // Fewer iterations for strings as they're more expensive
		_ = String(10)
	}
	stringDuration := time.Since(start)
	t.Logf("String(10) generation: %d operations in %v (%.2f ns/op)",
		iterations/10, stringDuration, float64(stringDuration.Nanoseconds())/float64(iterations/10))

	// Test UUID generation performance
	start = time.Now()
	for i := 0; i < iterations; i++ {
		_ = UUID()
	}
	uuidDuration := time.Since(start)
	t.Logf("UUID generation: %d operations in %v (%.2f ns/op)",
		iterations, uuidDuration, float64(uuidDuration.Nanoseconds())/float64(iterations))

	// Test range generation performance
	start = time.Now()
	for i := 0; i < iterations; i++ {
		_ = RangeInt(0, 1000)
	}
	rangeDuration := time.Since(start)
	t.Logf("RangeInt(0, 1000) generation: %d operations in %v (%.2f ns/op)",
		iterations, rangeDuration, float64(rangeDuration.Nanoseconds())/float64(iterations))
}

// TestMemoryUsage provides insights into memory allocation patterns
func TestMemoryUsage(t *testing.T) {
	const samples = 1000

	// Test basic integer functions - should have minimal allocations
	t.Run("IntegerMemoryUsage", func(t *testing.T) {
		for i := 0; i < samples; i++ {
			_ = Int32()
			_ = Int64()
			_ = Uint32()
			_ = Uint64()
		}
	})

	// Test range functions - should have minimal allocations due to object pooling
	t.Run("RangeMemoryUsage", func(t *testing.T) {
		for i := 0; i < samples; i++ {
			_ = RangeInt(0, 1000)
			_ = RangeInt64(0, 1000)
		}
	})

	// Test string functions - will have more allocations due to string building
	t.Run("StringMemoryUsage", func(t *testing.T) {
		for i := 0; i < samples/10; i++ { // Fewer samples as strings are more expensive
			_ = String(20)
			_ = VisibleString(20)
		}
	})
}

// BenchmarkVsStandardLib compares our implementation with standard library
func BenchmarkVsStandardLib(b *testing.B) {
	b.Run("OurInt32", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Int32()
		}
	})

	b.Run("StdLibInt31", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = getFallbackRand().Int31()
		}
	})

	b.Run("OurRangeInt", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = RangeInt(0, 1000)
		}
	})

	b.Run("StdLibIntn", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = getFallbackRand().Intn(1000)
		}
	})
}

// BenchmarkStringLengths tests performance across different string lengths
func BenchmarkStringLengths(b *testing.B) {
	lengths := []int{8, 16, 32, 64, 128, 256}

	for _, length := range lengths {
		b.Run(fmt.Sprintf("String-%d", length), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = String(length)
			}
		})

		b.Run(fmt.Sprintf("VisibleString-%d", length), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = VisibleString(length)
			}
		})
	}
}
