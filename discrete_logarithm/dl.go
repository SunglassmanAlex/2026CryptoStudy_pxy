package discrete_logarithm

import (
	"fmt"
	"math"
	"math/big"
)

// ComputeDiscreteLogarithm finds x such that a^x ≡ b (mod n) using Baby-step Giant-step.
// Assumes n is prime (needed for modular inverse of a^m mod n).
// Returns (x, nil) on success, or (nil, error) if no solution exists.
// Reference: https://en.wikipedia.org/wiki/Baby-step_giant-step
func ComputeDiscreteLogarithm(a, b, n *big.Int) (*big.Int, error) {
	// m = ceil(sqrt(n))
	m := ceilSqrt(n)
	one := big.NewInt(1)

	// --- Baby steps: build table of a^j mod n for j in [0, m) ---
	// TODO: finish Baby steps

	// --- Giant steps: for i in [0, m], check if b * (a^-m)^i mod n is in table ---
	// Compute a^m mod n, then its modular inverse: invAm = (a^m)^(-1) mod n
	// TODO: finish Giant steps

	return nil, fmt.Errorf("no solution: %v^x ≡ %v (mod %v) has no solution", a, b, n)
}

// ceilSqrt returns ceil(sqrt(n)) as a *big.Int.
func ceilSqrt(n *big.Int) *big.Int {
	if n.Sign() <= 0 {
		return big.NewInt(0)
	}
	// Start with float64 estimate
	nF, _ := new(big.Float).SetInt(n).Float64()
	sqrtF := new(big.Int).SetInt64(int64(math.Sqrt(nF)))

	// Newton's method refinement for accuracy with large integers
	one := big.NewInt(1)
	for {
		// next = (sqrtF + n/sqrtF) / 2
		next := new(big.Int).Div(n, sqrtF)
		next.Add(next, sqrtF)
		next.Rsh(next, 1) // divide by 2

		if next.Cmp(sqrtF) >= 0 {
			break
		}
		sqrtF.Set(next)
	}

	// Ensure we have ceil: if sqrtF^2 < n, add 1
	sq := new(big.Int).Mul(sqrtF, sqrtF)
	if sq.Cmp(n) < 0 {
		sqrtF.Add(sqrtF, one)
	}
	return sqrtF
}
