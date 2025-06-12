# TSRand

[![English](https://img.shields.io/badge/Language-English-blue.svg)](README.md) [![中文](https://img.shields.io/badge/Language-中文-red.svg)](README_CN.md)

A cryptographically secure random number generation library for Go, providing thread-safe random generators for integers, strings, and UUIDs with fallback mechanisms for better reliability.

[![Go Report Card](https://goreportcard.com/badge/github.com/tinystack/tsrand)](https://goreportcard.com/report/github.com/tinystack/tsrand)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.18-61CFDD.svg?style=flat-square)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/tinystack/tsrand)](https://pkg.go.dev/mod/github.com/tinystack/tsrand)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

## ✨ Features

- 🔐 **Cryptographically Secure**: Uses `crypto/rand` as primary source with `math/rand` fallback
- 🚀 **High Performance**: Optimized with object pooling and efficient algorithms (~350 ns/op for integers)
- 🧵 **Thread Safe**: All functions support concurrent access without additional locking
- 🎯 **Rich API**: Comprehensive set of functions for different use cases
- 🌍 **Unicode Support**: Full support for international characters via `CustomString`
- 🛡️ **Visual Safety**: `VisibleString` excludes ambiguous characters (0, O, I, l, 1)
- ⚡ **Memory Optimized**: Object pooling reduces GC pressure
- 🔄 **Reliable Fallback**: Automatic fallback ensures generation never fails

## 📦 Installation

```bash
go get -u github.com/tinystack/tsrand
```

## 🚀 Quick Start

```go
package main

import (
    "fmt"
    "github.com/tinystack/tsrand"
)

func main() {
    // Generate random integers
    fmt.Printf("Int32: %d\n", rand.Int32())
    fmt.Printf("Range: %d\n", rand.RangeInt(1, 100))

    // Generate random strings
    fmt.Printf("String: %s\n", rand.String(10))
    fmt.Printf("Visible: %s\n", rand.VisibleString(8))
    fmt.Printf("UUID: %s\n", rand.UUID())
}
```

## 📚 API Reference

### Integer Generation

| Function   | Description                     | Example         |
| ---------- | ------------------------------- | --------------- |
| `Int32()`  | Random int32 in [0, MaxInt32)   | `rand.Int32()`  |
| `Int64()`  | Random int64 in [0, MaxInt64)   | `rand.Int64()`  |
| `Uint32()` | Random uint32 in [0, MaxUint32) | `rand.Uint32()` |
| `Uint64()` | Random uint64 in [0, MaxUint64) | `rand.Uint64()` |
| `Int()`    | Random int in [0, MaxInt)       | `rand.Int()`    |
| `Uint()`   | Random uint in [0, MaxUint)     | `rand.Uint()`   |

### Range Generation

| Function                   | Description                   | Example                                 |
| -------------------------- | ----------------------------- | --------------------------------------- |
| `RangeInt(min, max)`       | Random int in [min, max)      | `rand.RangeInt(1, 100)`                 |
| `RangeInt64(min, max)`     | Random int64 in [min, max)    | `rand.RangeInt64(1000, 9999)`           |
| `RangeUint32(min, max)`    | Random uint32 in [min, max)   | `rand.RangeUint32(10, 50)`              |
| `RangeUint64(min, max)`    | Random uint64 in [min, max)   | `rand.RangeUint64(100, 1000)`           |
| `RangeIntSafe(min, max)`   | Range int with error return   | `n, err := rand.RangeIntSafe(1, 100)`   |
| `RangeInt64Safe(min, max)` | Range int64 with error return | `n, err := rand.RangeInt64Safe(1, 100)` |

### String Generation

| Function                        | Description          | Character Set      | Example                           |
| ------------------------------- | -------------------- | ------------------ | --------------------------------- |
| `String(length)`                | Alphanumeric string  | 0-9, a-z, A-Z      | `rand.String(10)`                 |
| `VisibleString(length)`         | No ambiguous chars   | Excludes 0,O,I,l,1 | `rand.VisibleString(8)`           |
| `AlphaString(length)`           | Alphabetic only      | a-z, A-Z           | `rand.AlphaString(10)`            |
| `NumericString(length)`         | Numeric only         | 0-9                | `rand.NumericString(6)`           |
| `LowercaseString(length)`       | Lowercase only       | a-z                | `rand.LowercaseString(8)`         |
| `UppercaseString(length)`       | Uppercase only       | A-Z                | `rand.UppercaseString(8)`         |
| `CustomString(charset, length)` | Custom character set | User-defined       | `rand.CustomString("ABC123", 10)` |
| `UUID()`                        | Standard UUID v4     | Hex + hyphens      | `rand.UUID()`                     |

## 🎯 Use Cases

### 🔐 Security Applications

```go
// Password generation (avoid confusing characters)
password := rand.VisibleString(16)

// API key generation
apiKey := fmt.Sprintf("ak_%s_%s",
    rand.LowercaseString(8), rand.String(16))

// Session token
sessionToken := rand.String(32)

// Cryptographic salt
salt := rand.String(24)
```

### 💻 Web Development

```go
// Verification code
verifyCode := rand.NumericString(6)

// Short URL ID (no confusing characters)
shortID := rand.VisibleString(8)

// User ID
userID := fmt.Sprintf("U%s%s",
    rand.UppercaseString(2), rand.NumericString(8))

// Invitation code
inviteCode := fmt.Sprintf("%s-%s",
    rand.UppercaseString(4), rand.UppercaseString(4))
```

### 🎮 Game Development

```go
// Dice simulation
dice1 := rand.RangeInt(1, 7)
dice2 := rand.RangeInt(1, 7)

// Random equipment attributes
attack := rand.RangeInt(50, 150)
durability := rand.RangeInt(80, 100)

// Player ID
playerID := fmt.Sprintf("P%s%s",
    rand.UppercaseString(2), rand.NumericString(6))

// Map seed
mapSeed := rand.Int64()
```

### 🌍 Internationalization

```go
// Chinese characters
chineseStr := rand.CustomString("天地玄黄宇宙洪荒", 5)

// Greek letters
greekStr := rand.CustomString("αβγδεζηθικλμν", 8)

// Custom emoji
emojiStr := rand.CustomString("😀😁😂🤣😃😄", 3)
```

## 📊 Performance

| Operation    | Performance | Memory    | Allocations   |
| ------------ | ----------- | --------- | ------------- |
| `Int32()`    | ~350 ns/op  | 48 B/op   | 3 allocs/op   |
| `Int64()`    | ~344 ns/op  | 48 B/op   | 3 allocs/op   |
| `RangeInt()` | ~350 ns/op  | 48 B/op   | 3 allocs/op   |
| `String(10)` | ~3.6 μs/op  | ~688 B/op | ~27 allocs/op |
| `UUID()`     | ~460 ns/op  | 64 B/op   | 2 allocs/op   |

_Benchmarks run on Intel Core i7-9750H @ 2.60GHz_

### Performance Scaling

- **Integer generation**: Constant time ~350 ns/op
- **String generation**: Linear scaling ~200-400 ns/character
- **Concurrent access**: Full thread safety with no performance penalty
- **Memory optimization**: Object pooling reduces GC pressure

## 🔧 Advanced Usage

### Error Handling

```go
// Safe range functions return errors for invalid ranges
if n, err := rand.RangeIntSafe(10, 100); err != nil {
    log.Printf("Invalid range: %v", err)
} else {
    fmt.Printf("Generated: %d", n)
}
```

### Complex Password Generation

```go
// Layered password with guaranteed character types
upper := rand.UppercaseString(4)
lower := rand.LowercaseString(4)
numbers := rand.NumericString(4)
special := rand.CustomString("!@#$%^&*", 2)

// Combine and shuffle
parts := []string{upper, lower, numbers, special}
// ... shuffle logic ...
password := strings.Join(parts, "")
```

### Business Applications

```go
// Order number
orderNum := fmt.Sprintf("ORD%s%s",
    rand.NumericString(8), rand.UppercaseString(4))

// Transaction ID
txnID := fmt.Sprintf("TXN_%s_%s",
    rand.NumericString(10), rand.String(8))

// Product serial number
serialNum := fmt.Sprintf("%s-%s-%s",
    rand.UppercaseString(4),
    rand.NumericString(4),
    rand.UppercaseString(4))
```

## 🧪 Examples

See the [`examples/`](examples/) directory for comprehensive examples:

- **[`main.go`](examples/main.go)**: Complete feature demonstration
- **[`performance_demo.go`](examples/performance_demo.go)**: Performance analysis
- **[`README.md`](examples/README.md)**: Detailed examples documentation

```bash
# Run complete examples
cd examples
go run main.go

# Run performance demo
go run performance_demo.go
```

## 🔍 Testing

```bash
# Run all tests
go test -v

# Run benchmarks
go test -bench=. -benchmem

# Run with race detection
go test -race

# Test coverage
go test -cover
```

## 🛡️ Security

- **Primary source**: `crypto/rand` for cryptographically secure random generation
- **Fallback mechanism**: Automatic fallback to `math/rand` when `crypto/rand` is unavailable
- **Thread safety**: All functions are safe for concurrent use
- **No blocking**: Never blocks even when system entropy is low

## 🎨 Design Principles

1. **Security First**: Cryptographic security by default
2. **Performance Optimized**: Efficient algorithms and memory management
3. **Developer Friendly**: Intuitive API with comprehensive documentation
4. **Reliability**: Multiple fallback mechanisms ensure functionality
5. **Extensibility**: Easy to extend with custom character sets

## 📋 Requirements

- Go 1.18 or later
- No external dependencies beyond Go standard library

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Go team for the excellent `crypto/rand` and `math/rand` packages
- Contributors and users of this library
- Open source community for inspiration and feedback

## 📞 Support

- 📖 [Documentation](https://pkg.go.dev/github.com/tinystack/tsrand)
- 🐛 [Issue Tracker](https://github.com/tinystack/tsrand/issues)
- 💬 [Discussions](https://github.com/tinystack/tsrand/discussions)

---

**Made with ❤️ for the Go community**
