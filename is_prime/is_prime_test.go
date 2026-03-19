package is_prime

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func TestIsPrime(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(100000000))
		if err != nil {
			panic("Oh my god!")
		}
		expected := n.ProbablyPrime(20)
		got := IsPrime(n)
		if expected != got {
			t.Errorf("恢复失败：n = %s, expected %v, got %v", n.String(), expected, got)
		}
	}
}

func BenchmarkIsPrime(b *testing.B) {
	n := big.NewInt(998244353)
	for i := 0; i < b.N; i++ {
		IsPrime(n)
	}
}

func BenchmarkProbablyPrime(b *testing.B) {
	n := big.NewInt(998244353)
	for i := 0; i < b.N; i++ {
		n.ProbablyPrime(20)
	}
}

// try push
