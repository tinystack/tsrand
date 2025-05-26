package rand

import (
	cRand "crypto/rand"
	"math/big"
)

// RangeInt returns a uniform random value in [min, max). It returns 0 If min <=0 or min >max
func RangeInt(min, max int) int {
	if min == max {
		return min
	}
	if min > max {
		return 0
	}
	rag := max - min
	r, err := cRand.Int(cRand.Reader, new(big.Int).SetInt64(int64(rag)))
	if err != nil {
		return newSimpleRand.Intn(rag) + min
	}
	return int(r.Int64()) + min
}

// RangeInt64 returns a uniform random value in [min, max). It returns 0 If min <=0 or min >max
func RangeInt64(min, max int64) int64 {
	if min == max {
		return min
	}
	if min > max {
		return 0
	}
	rag := max - min
	r, err := cRand.Int(cRand.Reader, new(big.Int).SetInt64(rag))
	if err != nil {
		return newSimpleRand.Int63n(rag) + min
	}
	return r.Int64() + min
}
