package rand

import (
	"math/rand"
	"time"
)

var (
	newSimpleRand  = rand.New(rand.NewSource(time.Now().UnixNano()))
	visibleLetters = "23456789abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	normalLetters  = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)
