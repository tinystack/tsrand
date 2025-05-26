package rand

import (
	cRand "crypto/rand"
	"math"
	"math/big"
)

func randUint64() uint64 {
	r, err := cRand.Int(cRand.Reader, new(big.Int).SetUint64(math.MaxUint64))
	if err != nil {
		return newSimpleRand.Uint64()
	}
	return r.Uint64()
}

func randInt64() int64 {
	r, err := cRand.Int(cRand.Reader, new(big.Int).SetInt64(math.MaxInt64))
	if err != nil {
		return newSimpleRand.Int63()
	}
	return r.Int64()
}

// Int32 returns a uniform random value in [0, math.MaxInt32)
func Int32() int32 {
	return int32(randInt64() >> 32)
}

// Int64 returns a uniform random value in [0, math.MaxInt64)
func Int64() int64 {
	return randInt64()
}

// Uint32 returns a uniform random value in [0, math.MaxUint32)
func Uint32() uint32 {
	return uint32(randUint64() >> 32)
}

// Uint64 returns a uniform random value in [0, math.MaxUint64)
func Uint64() uint64 {
	return randUint64()
}

// Int returns a uniform random value in [0, math.MaxInt)
func Int() int {
	n := randInt64()
	if n <= math.MaxInt {
		return int(n)
	}
	return int(n % math.MaxInt)
}

// Uint returns a uniform random value in [0, math.MaxUint)
func Uint() uint {
	n := randUint64()
	if n <= math.MaxUint {
		return uint(n)
	}
	return uint(n % math.MaxUint)
}
