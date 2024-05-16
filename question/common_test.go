package question

import (
	"testing"
)

func TestRandint(t *testing.T) {
	for range 10000 {
		n := randint(20, 30)
		if n < 20 {
			t.Error(n)
		}
		if n >= 30 {
			t.Error(n)
		}
	}
}
