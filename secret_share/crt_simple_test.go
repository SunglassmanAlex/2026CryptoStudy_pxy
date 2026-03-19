package secret_share

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func TestSolveSimpleCRT_Classic(t *testing.T) {
	// x ≡ 2 (mod 3), x ≡ 3 (mod 5), x ≡ 2 (mod 7) => x = 23 (mod 105)
	equations := []Equation{
		{big.NewInt(2), big.NewInt(3)},
		{big.NewInt(3), big.NewInt(5)},
		{big.NewInt(2), big.NewInt(7)},
	}

	expected := big.NewInt(23)
	got := SolveSimpleCRT(equations)

	if got == nil || got.Cmp(expected) != 0 {
		t.Errorf("Classic case failed: got %v, want %v", got, expected)
	}
}

func TestSolveSimpleCRT_LargeNumbers(t *testing.T) {
	n1, _ := new(big.Int).SetString("1000000000000000003", 10)
	n2, _ := new(big.Int).SetString("1000000000000000009", 10)
	a1 := big.NewInt(12345)
	a2 := big.NewInt(67890)

	equations := []Equation{
		{a1, n1},
		{a2, n2},
	}

	got := SolveSimpleCRT(equations)

	if got == nil {
		t.Fatal("Large numbers case failed: got nil")
	}

	for _, eq := range equations {
		rem := new(big.Int).Mod(got, eq.Modulus)
		if rem.Cmp(eq.Remainder) != 0 {
			t.Errorf("Large numbers failed: x ≡ %v (mod %v) is not satisfied, got remainder %v",
				eq.Remainder, eq.Modulus, rem)
		}
	}
}

func generateRandomPrime(bits int) *big.Int {
	p, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	return p
}

func TestSolveSimpleCRT_PropertyBased(t *testing.T) {
	numRounds := 1000
	numEquations := 5
	bitSize := 64

	for i := 0; i < numRounds; i++ {
		var equations []Equation
		expectedM := big.NewInt(1)

		usedPrimes := make(map[string]bool)
		for len(equations) < numEquations {
			m := generateRandomPrime(bitSize)

			if usedPrimes[m.String()] {
				continue
			}
			usedPrimes[m.String()] = true

			a, _ := rand.Int(rand.Reader, m)

			equations = append(equations, Equation{
				Remainder: a,
				Modulus:   m,
			})

			expectedM.Mul(expectedM, m)
		}

		gotX := SolveSimpleCRT(equations)

		if gotX == nil {
			t.Fatalf("Round %d: SolveSimpleCRT returned nil", i)
		}

		if gotX.Sign() < 0 || gotX.Cmp(expectedM) >= 0 {
			t.Errorf("Round %d: Result x is out of range [0, M). got %v, M %v", i, gotX, expectedM)
		}

		for _, eq := range equations {
			rem := new(big.Int).Mod(gotX, eq.Modulus)
			if rem.Cmp(eq.Remainder) != 0 {
				t.Errorf("Round %d: Equation not satisfied. x ≡ %v (mod %v) failed, got remainder %v",
					i, eq.Remainder, eq.Modulus, rem)
			}
		}
	}
}
