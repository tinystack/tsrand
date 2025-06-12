package rand

import (
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"
)

// randStringFromCharset generates a cryptographically secure random string of specified length
// using characters from the given charset.
//
// This function uses crypto/rand for secure random generation with fallback to math/rand.
// It uses an optimized approach with pre-computed charset length and efficient string building.
//
// Parameters:
//   - charset: the character set to choose from
//   - length: the desired length of the generated string
//
// Returns:
//   - A random string of the specified length using characters from the charset
func randStringFromCharset(charset string, length int) string {
	if length <= 0 {
		return ""
	}

	charsetLen := utf8.RuneCountInString(charset)
	if charsetLen == 0 {
		return ""
	}

	// Convert charset to rune slice for proper Unicode support
	charsetRunes := []rune(charset)
	charsetLen = len(charsetRunes)

	// Pre-allocate the result slice for better performance
	result := make([]rune, length)

	maxBig := getBigInt()
	defer putBigInt(maxBig)
	maxBig.SetInt64(int64(charsetLen))

	for i := 0; i < length; i++ {
		if randResult, ok := secureRandomBigInt(maxBig); ok {
			result[i] = charsetRunes[randResult.Int64()]
		} else {
			// Fallback to pseudo-random generation
			result[i] = charsetRunes[getFallbackRand().Intn(charsetLen)]
		}
	}

	return string(result)
}

// String generates a cryptographically secure random string of the specified length
// using all alphanumeric characters (0-9, a-z, A-Z).
//
// The function uses crypto/rand for secure random generation with fallback to math/rand.
// The character set includes all digits and both uppercase and lowercase letters.
//
// Parameters:
//   - length: the desired length of the generated string
//
// Returns:
//   - A random alphanumeric string of the specified length
//
// Example:
//
//	s := rand.String(10) // Returns something like "aB3xY9mK7Q"
func String(length int) string {
	return randStringFromCharset(NormalLetters, length)
}

// VisibleString generates a cryptographically secure random string of the specified length
// using only visually distinct characters.
//
// This function excludes characters that can be easily confused in many fonts:
//   - 0 (zero) vs O (capital O)
//   - 1 (one) vs I (capital i) vs l (lowercase L)
//
// The function uses crypto/rand for secure random generation with fallback to math/rand.
// This is particularly useful for generating codes that users need to read and type.
//
// Parameters:
//   - length: the desired length of the generated string
//
// Returns:
//   - A random string of visually distinct characters with the specified length
//
// Example:
//
//	s := rand.VisibleString(8) // Returns something like "a8Bx9mKQ" (no confusing chars)
func VisibleString(length int) string {
	return randStringFromCharset(VisibleLetters, length)
}

// CustomString generates a cryptographically secure random string of the specified length
// using a custom character set.
//
// This function allows you to specify your own character set for random string generation.
// The function uses crypto/rand for secure random generation with fallback to math/rand.
//
// Parameters:
//   - charset: the custom character set to use for generation
//   - length: the desired length of the generated string
//
// Returns:
//   - A random string using characters from the custom charset
//
// Example:
//
//	s := rand.CustomString("ABCDEF0123456789", 8) // Hexadecimal-style string
func CustomString(charset string, length int) string {
	return randStringFromCharset(charset, length)
}

// AlphaString generates a cryptographically secure random string of the specified length
// using only alphabetic characters (a-z, A-Z).
//
// Parameters:
//   - length: the desired length of the generated string
//
// Returns:
//   - A random alphabetic string of the specified length
//
// Example:
//
//	s := rand.AlphaString(10) // Returns something like "aBxYmKqWeR"
func AlphaString(length int) string {
	const alphaChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return randStringFromCharset(alphaChars, length)
}

// NumericString generates a cryptographically secure random string of the specified length
// using only numeric characters (0-9).
//
// Parameters:
//   - length: the desired length of the generated string
//
// Returns:
//   - A random numeric string of the specified length
//
// Example:
//
//	s := rand.NumericString(6) // Returns something like "138947"
func NumericString(length int) string {
	const numericChars = "0123456789"
	return randStringFromCharset(numericChars, length)
}

// LowercaseString generates a cryptographically secure random string of the specified length
// using only lowercase alphabetic characters (a-z).
//
// Parameters:
//   - length: the desired length of the generated string
//
// Returns:
//   - A random lowercase string of the specified length
//
// Example:
//
//	s := rand.LowercaseString(8) // Returns something like "abxymkqw"
func LowercaseString(length int) string {
	const lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
	return randStringFromCharset(lowercaseChars, length)
}

// UppercaseString generates a cryptographically secure random string of the specified length
// using only uppercase alphabetic characters (A-Z).
//
// Parameters:
//   - length: the desired length of the generated string
//
// Returns:
//   - A random uppercase string of the specified length
//
// Example:
//
//	s := rand.UppercaseString(8) // Returns something like "ABXYMKQW"
func UppercaseString(length int) string {
	const uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return randStringFromCharset(uppercaseChars, length)
}

// UUID generates a cryptographically secure Version 4 UUID string.
//
// This function attempts to generate a Version 4 (random) UUID using crypto/rand.
// If that fails, it falls back to a Version 1 (time-based) UUID.
// The returned UUID follows the standard format: xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
//
// Returns:
//   - A UUID string in standard format
//
// Example:
//
//	id := rand.UUID() // Returns something like "550e8400-e29b-41d4-a716-446655440000"
func UUID() string {
	// Attempt to generate a Version 4 (random) UUID
	if u, err := uuid.NewRandom(); err == nil {
		return u.String()
	}

	// Fallback to Version 1 (time-based) UUID
	// This should rarely happen as uuid.NewRandom() is very reliable
	if u, err := uuid.NewUUID(); err == nil {
		return u.String()
	}

	// Final fallback: generate a pseudo-UUID using our own random functions
	// This maintains the UUID format but may not be standards-compliant
	return generateFallbackUUID()
}

// generateFallbackUUID creates a pseudo-UUID when the uuid package fails entirely.
// This is a rare fallback scenario that maintains UUID format.
func generateFallbackUUID() string {
	const hexChars = "0123456789abcdef"

	var parts []string
	parts = append(parts, randStringFromCharset(hexChars, 8))     // 8 hex chars
	parts = append(parts, randStringFromCharset(hexChars, 4))     // 4 hex chars
	parts = append(parts, "4"+randStringFromCharset(hexChars, 3)) // Version 4 + 3 hex chars

	// Generate the variant bits (8, 9, A, or B)
	variantChars := "89ab"
	variant := randStringFromCharset(variantChars, 1)
	parts = append(parts, variant+randStringFromCharset(hexChars, 3)) // Variant + 3 hex chars
	parts = append(parts, randStringFromCharset(hexChars, 12))        // 12 hex chars

	return strings.Join(parts, "-")
}
