package rand

import (
	cRand "crypto/rand"
	"math/big"
	"strings"

	"github.com/google/uuid"
)

func randString(letters string, l int) string {
	var s []string
	b := new(big.Int).SetInt64(int64(len(letters)))
	for i := 0; i < l; i++ {
		if i, err := cRand.Int(cRand.Reader, b); err == nil {
			s = append(s, string(letters[i.Int64()]))
		}
	}
	return strings.Join(s, "")
}

func String(len int) string {
	return randString(normalLetters, len)
}

func VisibleString(len int) string {
	return randString(visibleLetters, len)
}

func UUID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		u, _ = uuid.NewUUID()
	}
	return u.String()
}
