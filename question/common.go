package question

import (
	"crypto/rand"
	"math/big"
)

func randint(min, max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		panic(err)
	}

	return min + int(n.Int64())
}
