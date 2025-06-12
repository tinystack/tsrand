package rand

import (
	"regexp"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestString validates the String function
func TestString(t *testing.T) {
	// Test length
	lengths := []int{0, 1, 10, 50, 100, 1000}
	for _, length := range lengths {
		s := String(length)
		assert.Equal(t, length, len(s), "String(%d) should return string of length %d", length, length)

		if length > 0 {
			// Test that all characters are from NormalLetters
			for _, char := range s {
				assert.Contains(t, NormalLetters, string(char),
					"String() should only contain characters from NormalLetters")
			}
		}
	}

	// Test uniqueness
	strings1 := make([]string, 100)
	for i := range strings1 {
		strings1[i] = String(20)
	}

	// Count unique strings
	uniqueStrings := make(map[string]bool)
	for _, s := range strings1 {
		uniqueStrings[s] = true
	}

	// Should have many unique strings
	assert.Greater(t, len(uniqueStrings), 95, "String() should generate mostly unique strings")
}

// TestVisibleString validates the VisibleString function
func TestVisibleString(t *testing.T) {
	// Test length
	lengths := []int{0, 1, 10, 50, 100}
	for _, length := range lengths {
		s := VisibleString(length)
		assert.Equal(t, length, len(s), "VisibleString(%d) should return string of length %d", length, length)

		if length > 0 {
			// Test that all characters are from VisibleLetters
			for _, char := range s {
				assert.Contains(t, VisibleLetters, string(char),
					"VisibleString() should only contain characters from VisibleLetters")
			}
		}
	}

	// Test that ambiguous characters are excluded
	excludedChars := []string{"0", "1", "I", "l", "O", "o"}
	for i := 0; i < 1000; i++ {
		s := VisibleString(100)
		for _, excluded := range excludedChars {
			assert.NotContains(t, s, excluded,
				"VisibleString() should not contain ambiguous character: %s", excluded)
		}
	}

	// Test uniqueness
	for i := 0; i < 100; i++ {
		s1 := VisibleString(20)
		s2 := VisibleString(20)
		assert.NotEqual(t, s1, s2, "VisibleString() should generate unique strings")
	}
}

// TestCustomString validates the CustomString function
func TestCustomString(t *testing.T) {
	// Test with custom charset
	charset := "ABC123"
	length := 20

	s := CustomString(charset, length)
	assert.Equal(t, length, len(s), "CustomString should return string of specified length")

	// Test that all characters are from the custom charset
	for _, char := range s {
		assert.Contains(t, charset, string(char),
			"CustomString() should only contain characters from custom charset")
	}

	// Test with empty charset
	s = CustomString("", 10)
	assert.Equal(t, "", s, "CustomString with empty charset should return empty string")

	// Test with zero length
	s = CustomString("ABC", 0)
	assert.Equal(t, "", s, "CustomString with zero length should return empty string")

	// Test with negative length
	s = CustomString("ABC", -5)
	assert.Equal(t, "", s, "CustomString with negative length should return empty string")
}

// TestAlphaString validates the AlphaString function
func TestAlphaString(t *testing.T) {
	s := AlphaString(50)
	assert.Equal(t, 50, len(s), "AlphaString should return string of specified length")

	// Test that all characters are alphabetic
	for _, char := range s {
		assert.True(t, unicode.IsLetter(char),
			"AlphaString() should only contain alphabetic characters, found: %c", char)
	}
}

// TestNumericString validates the NumericString function
func TestNumericString(t *testing.T) {
	s := NumericString(20)
	assert.Equal(t, 20, len(s), "NumericString should return string of specified length")

	// Test that all characters are numeric
	for _, char := range s {
		assert.True(t, unicode.IsDigit(char),
			"NumericString() should only contain numeric characters, found: %c", char)
	}

	// Test with regex pattern
	matched, err := regexp.MatchString("^[0-9]+$", s)
	require.NoError(t, err)
	assert.True(t, matched, "NumericString should match numeric regex pattern")
}

// TestLowercaseString validates the LowercaseString function
func TestLowercaseString(t *testing.T) {
	s := LowercaseString(30)
	assert.Equal(t, 30, len(s), "LowercaseString should return string of specified length")

	// Test that all characters are lowercase letters
	for _, char := range s {
		assert.True(t, unicode.IsLetter(char) && unicode.IsLower(char),
			"LowercaseString() should only contain lowercase letters, found: %c", char)
	}
}

// TestUppercaseString validates the UppercaseString function
func TestUppercaseString(t *testing.T) {
	s := UppercaseString(25)
	assert.Equal(t, 25, len(s), "UppercaseString should return string of specified length")

	// Test that all characters are uppercase letters
	for _, char := range s {
		assert.True(t, unicode.IsLetter(char) && unicode.IsUpper(char),
			"UppercaseString() should only contain uppercase letters, found: %c", char)
	}
}

// TestUUID validates the UUID function
func TestUUID(t *testing.T) {
	// Test basic properties
	uuid := UUID()
	assert.Equal(t, 36, len(uuid), "UUID should be 36 characters long")

	// Test UUID format (8-4-4-4-12)
	parts := strings.Split(uuid, "-")
	assert.Len(t, parts, 5, "UUID should have 5 parts separated by hyphens")
	assert.Len(t, parts[0], 8, "First UUID part should be 8 characters")
	assert.Len(t, parts[1], 4, "Second UUID part should be 4 characters")
	assert.Len(t, parts[2], 4, "Third UUID part should be 4 characters")
	assert.Len(t, parts[3], 4, "Fourth UUID part should be 4 characters")
	assert.Len(t, parts[4], 12, "Fifth UUID part should be 12 characters")

	// Test UUID regex pattern
	uuidPattern := `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`
	matched, err := regexp.MatchString(uuidPattern, uuid)
	require.NoError(t, err)
	assert.True(t, matched, "UUID should match standard format pattern")

	// Test uniqueness
	uuids := make(map[string]bool, 10000)
	duplicates := 0

	for i := 0; i < 10000; i++ {
		u := UUID()
		if uuids[u] {
			duplicates++
		}
		uuids[u] = true
	}

	// There should be no duplicates in UUID generation
	assert.Equal(t, 0, duplicates, "UUIDs should be unique, found %d duplicates", duplicates)
}

// TestStringConcurrency validates thread safety of string functions
func TestStringConcurrency(t *testing.T) {
	const goroutines = 50
	const iterations = 100

	results := make(chan string, goroutines*iterations)

	// Launch multiple goroutines
	for g := 0; g < goroutines; g++ {
		go func() {
			for i := 0; i < iterations; i++ {
				results <- String(20)
			}
		}()
	}

	// Collect all results
	strings := make(map[string]bool)
	for i := 0; i < goroutines*iterations; i++ {
		s := <-results
		strings[s] = true

		// Basic validation
		require.Equal(t, 20, len(s), "Concurrent String() should return correct length")

		// Check all characters are valid
		for _, char := range s {
			require.Contains(t, NormalLetters, string(char),
				"Concurrent String() should only contain valid characters")
		}
	}

	// Should have generated many unique strings
	assert.Greater(t, len(strings), goroutines*iterations/2,
		"Concurrent string generation should produce many unique values")
}

// TestUnicodeSupport validates Unicode character support
func TestUnicodeSupport(t *testing.T) {
	// Test with Unicode characters in custom charset
	unicodeCharset := "αβγδεζηθικλμνξοπρστυφχψω"
	s := CustomString(unicodeCharset, 10)

	assert.Equal(t, 10, len([]rune(s)), "Unicode string should have correct rune length")

	// Test that all characters are from the Unicode charset
	for _, char := range s {
		assert.Contains(t, unicodeCharset, string(char),
			"CustomString with Unicode should only contain characters from charset")
	}
}

// BenchmarkString benchmarks the String function
func BenchmarkString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = String(50)
	}
}

// BenchmarkVisibleString benchmarks the VisibleString function
func BenchmarkVisibleString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = VisibleString(50)
	}
}

// BenchmarkCustomString benchmarks the CustomString function
func BenchmarkCustomString(b *testing.B) {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CustomString(charset, 50)
	}
}

// BenchmarkUUID benchmarks the UUID function
func BenchmarkUUID(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = UUID()
	}
}
