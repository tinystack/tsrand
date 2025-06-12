// Package rand provides cryptographically secure random number generation utilities.
// It offers thread-safe random number generators for various types including integers,
// strings, and UUIDs with fallback mechanisms for better reliability.
//
// This package prioritizes security by using crypto/rand as the primary source,
// falling back to math/rand when cryptographic randomness is unavailable.
//
// Example usage:
//
//	import "github.com/tinystack/tsrand"
//
//	// Generate random integers
//	n := rand.Int32()
//	u := rand.Uint64()
//
//	// Generate random strings
//	s := rand.String(10)        // Contains all alphanumeric characters
//	v := rand.VisibleString(10) // Excludes ambiguous characters (0, O, I, l, etc.)
//
//	// Generate UUIDs
//	id := rand.UUID()
//
//	// Generate random numbers in range
//	n := rand.RangeInt(10, 100)
package rand

import (
	cRand "crypto/rand"
	"math/big"
	"math/rand"
	"sync"
	"time"
)

// Character sets used for string generation
const (
	// NormalLetters contains all alphanumeric characters (0-9, a-z, A-Z)
	NormalLetters = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// VisibleLetters excludes ambiguous characters that can be easily confused:
	// Excludes: 0 (zero), O (capital o), I (capital i), l (lowercase L), 1 (one)
	VisibleLetters = "23456789abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
)

var (
	// Global pseudo-random generator used as fallback when crypto/rand fails
	fallbackRand     *rand.Rand
	fallbackRandOnce sync.Once

	// Object pools for reducing allocations
	bigIntPool = sync.Pool{
		New: func() interface{} {
			return new(big.Int)
		},
	}
)

// initFallbackRand initializes the fallback random generator
func initFallbackRand() {
	fallbackRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// getFallbackRand returns the fallback random generator, initializing it if necessary
func getFallbackRand() *rand.Rand {
	fallbackRandOnce.Do(initFallbackRand)
	return fallbackRand
}

// getBigInt gets a big.Int from the pool
func getBigInt() *big.Int {
	return bigIntPool.Get().(*big.Int)
}

// putBigInt returns a big.Int to the pool after resetting it
func putBigInt(bi *big.Int) {
	bi.SetInt64(0)
	bigIntPool.Put(bi)
}

// secureRandomBigInt generates a cryptographically secure random number in [0, max)
// Returns the random number and a boolean indicating success
func secureRandomBigInt(max *big.Int) (*big.Int, bool) {
	result, err := cRand.Int(cRand.Reader, max)
	return result, err == nil
}
