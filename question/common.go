package question

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

// randomDigitRune returns a random digit other than n
func randomDigitRune(n string) rune {
	nn, err := strconv.Atoi(n)
	if err != nil {
		return []rune(n)[0]
	}
	for {
		i := randint(0, 9)
		if i == nn {
			continue
		}
		return []rune(strconv.Itoa(i))[0]
	}

}

func randint(min, max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		panic(err)
	}

	return min + int(n.Int64())
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
