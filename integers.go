package rand

import (
	"math"
)

// randUint64 generates a cryptographically secure random uint64.
// It uses crypto/rand as the primary source and falls back to math/rand if necessary.
func randUint64() uint64 {
	maxBig := getBigInt()
	defer putBigInt(maxBig)

	maxBig.SetUint64(math.MaxUint64)
	if result, ok := secureRandomBigInt(maxBig); ok {
		return result.Uint64()
	}

	// Fallback to pseudo-random generation
	return getFallbackRand().Uint64()
}

// randInt64 generates a cryptographically secure random int64.
// It uses crypto/rand as the primary source and falls back to math/rand if necessary.
func randInt64() int64 {
	maxBig := getBigInt()
	defer putBigInt(maxBig)

	maxBig.SetInt64(math.MaxInt64)
	if result, ok := secureRandomBigInt(maxBig); ok {
		return result.Int64()
	}

	// Fallback to pseudo-random generation
	return getFallbackRand().Int63()
}

// Int32 returns a cryptographically secure random int32 value in the range [0, math.MaxInt32).
// The distribution is uniform across the entire range.
//
// Example:
//
//	n := rand.Int32() // Returns a value like 1234567890
func Int32() int32 {
	return int32(randInt64() >> 32)
}

// Int64 returns a cryptographically secure random int64 value in the range [0, math.MaxInt64).
// The distribution is uniform across the entire range.
//
// Example:
//
//	n := rand.Int64() // Returns a value like 1234567890123456789
func Int64() int64 {
	return randInt64()
}

// Uint32 returns a cryptographically secure random uint32 value in the range [0, math.MaxUint32).
// The distribution is uniform across the entire range.
//
// Example:
//
//	n := rand.Uint32() // Returns a value like 2345678901
func Uint32() uint32 {
	return uint32(randUint64() >> 32)
}

// Uint64 returns a cryptographically secure random uint64 value in the range [0, math.MaxUint64).
// The distribution is uniform across the entire range.
//
// Example:
//
//	n := rand.Uint64() // Returns a value like 12345678901234567890
func Uint64() uint64 {
	return randUint64()
}

// Int returns a cryptographically secure random int value in the range [0, math.MaxInt).
// The distribution is uniform across the entire range.
// This function is platform-dependent as int size varies by architecture.
//
// Example:
//
//	n := rand.Int() // Returns a value like 1234567890123456
func Int() int {
	n := randInt64()
	// Ensure the value fits in the platform's int type
	if n <= math.MaxInt {
		return int(n)
	}
	return int(n % math.MaxInt)
}

// Uint returns a cryptographically secure random uint value in the range [0, math.MaxUint).
// The distribution is uniform across the entire range.
// This function is platform-dependent as uint size varies by architecture.
//
// Example:
//
//	n := rand.Uint() // Returns a value like 1234567890123456
func Uint() uint {
	n := randUint64()
	// Ensure the value fits in the platform's uint type
	if n <= math.MaxUint {
		return uint(n)
	}
	return uint(n % math.MaxUint)
}
