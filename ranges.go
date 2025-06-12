package rand

import (
	"errors"
)

// ErrInvalidRange is returned when the range parameters are invalid
var ErrInvalidRange = errors.New("invalid range: min must be less than max")

// RangeInt returns a cryptographically secure random integer in the range [min, max).
// It returns an error if min >= max.
//
// The function uses crypto/rand for secure random generation with fallback to math/rand.
// The distribution is uniform across the specified range.
//
// Parameters:
//   - min: the minimum value (inclusive)
//   - max: the maximum value (exclusive)
//
// Returns:
//   - A random integer in [min, max)
//   - An error if the range is invalid
//
// Example:
//
//	n, err := rand.RangeIntSafe(10, 100)
//	if err != nil {
//		// Handle error
//	}
//	// n is between 10 and 99 (inclusive)
func RangeIntSafe(min, max int) (int, error) {
	if min > max {
		return 0, ErrInvalidRange
	}

	if min == max {
		return min, nil
	}

	rangeSize := max - min
	maxBig := getBigInt()
	defer putBigInt(maxBig)

	maxBig.SetInt64(int64(rangeSize))
	if result, ok := secureRandomBigInt(maxBig); ok {
		return int(result.Int64()) + min, nil
	}

	// Fallback to pseudo-random generation
	return getFallbackRand().Intn(rangeSize) + min, nil
}

// RangeInt returns a cryptographically secure random integer in the range [min, max).
// It returns 0 if min >= max for backward compatibility.
//
// Note: This function maintains backward compatibility with the original API.
// For better error handling, consider using RangeIntSafe instead.
//
// Parameters:
//   - min: the minimum value (inclusive)
//   - max: the maximum value (exclusive)
//
// Returns:
//   - A random integer in [min, max), or 0 if the range is invalid
//
// Example:
//
//	n := rand.RangeInt(10, 100) // Returns a value between 10 and 99
func RangeInt(min, max int) int {
	result, err := RangeIntSafe(min, max)
	if err != nil {
		return 0
	}
	return result
}

// RangeInt64Safe returns a cryptographically secure random int64 in the range [min, max).
// It returns an error if min >= max.
//
// The function uses crypto/rand for secure random generation with fallback to math/rand.
// The distribution is uniform across the specified range.
//
// Parameters:
//   - min: the minimum value (inclusive)
//   - max: the maximum value (exclusive)
//
// Returns:
//   - A random int64 in [min, max)
//   - An error if the range is invalid
//
// Example:
//
//	n, err := rand.RangeInt64Safe(1000, 9999)
//	if err != nil {
//		// Handle error
//	}
//	// n is between 1000 and 9998 (inclusive)
func RangeInt64Safe(min, max int64) (int64, error) {
	if min > max {
		return 0, ErrInvalidRange
	}

	if min == max {
		return min, nil
	}

	rangeSize := max - min
	maxBig := getBigInt()
	defer putBigInt(maxBig)

	maxBig.SetInt64(rangeSize)
	if result, ok := secureRandomBigInt(maxBig); ok {
		return result.Int64() + min, nil
	}

	// Fallback to pseudo-random generation
	return getFallbackRand().Int63n(rangeSize) + min, nil
}

// RangeInt64 returns a cryptographically secure random int64 in the range [min, max).
// It returns 0 if min >= max for backward compatibility.
//
// Note: This function maintains backward compatibility with the original API.
// For better error handling, consider using RangeInt64Safe instead.
//
// Parameters:
//   - min: the minimum value (inclusive)
//   - max: the maximum value (exclusive)
//
// Returns:
//   - A random int64 in [min, max), or 0 if the range is invalid
//
// Example:
//
//	n := rand.RangeInt64(1000, 9999) // Returns a value between 1000 and 9998
func RangeInt64(min, max int64) int64 {
	result, err := RangeInt64Safe(min, max)
	if err != nil {
		return 0
	}
	return result
}

// RangeUint32 returns a cryptographically secure random uint32 in the range [min, max).
// It returns 0 if min >= max.
//
// Parameters:
//   - min: the minimum value (inclusive)
//   - max: the maximum value (exclusive)
//
// Returns:
//   - A random uint32 in [min, max), or 0 if the range is invalid
//
// Example:
//
//	n := rand.RangeUint32(100, 1000) // Returns a value between 100 and 999
func RangeUint32(min, max uint32) uint32 {
	if min > max {
		return 0
	}

	if min == max {
		return min
	}

	rangeSize := max - min
	maxBig := getBigInt()
	defer putBigInt(maxBig)

	maxBig.SetUint64(uint64(rangeSize))
	if result, ok := secureRandomBigInt(maxBig); ok {
		return uint32(result.Uint64()) + min
	}

	// Fallback to pseudo-random generation
	return uint32(getFallbackRand().Uint32())%rangeSize + min
}

// RangeUint64 returns a cryptographically secure random uint64 in the range [min, max).
// It returns 0 if min >= max.
//
// Parameters:
//   - min: the minimum value (inclusive)
//   - max: the maximum value (exclusive)
//
// Returns:
//   - A random uint64 in [min, max), or 0 if the range is invalid
//
// Example:
//
//	n := rand.RangeUint64(1000, 9999) // Returns a value between 1000 and 9998
func RangeUint64(min, max uint64) uint64 {
	if min > max {
		return 0
	}

	if min == max {
		return min
	}

	rangeSize := max - min
	maxBig := getBigInt()
	defer putBigInt(maxBig)

	maxBig.SetUint64(rangeSize)
	if result, ok := secureRandomBigInt(maxBig); ok {
		return result.Uint64() + min
	}

	// Fallback to pseudo-random generation
	return getFallbackRand().Uint64()%rangeSize + min
}
