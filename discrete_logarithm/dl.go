package discrete_logarithm

import (
	"fmt"
	"math"
	"math/big"
)

type DLOption int

const (
	DLBruteForce DLOption = iota
	DLBabyStepGiantStep
	DLPollardRho
)

func ComputeDiscreteLogarithm(a, b, n *big.Int, op DLOption) (*big.Int, error) {
	switch op {
	case DLBruteForce:
		return computeDiscreteLogarithmBruteForced(a, b, n)
	case DLBabyStepGiantStep:
		return computeDiscreteLogarithmBSGS(a, b, n)
	case DLPollardRho:
		return computeDiscreteLogarithmPollardRho(a, b, n)
	}
	return nil, fmt.Errorf("unsupported operation")
}

func computeDiscreteLogarithmBruteForced(a, b, n *big.Int) (*big.Int, error) {
	aExpI := big.NewInt(1)
	for i := big.NewInt(0); i.Cmp(n) <= 0; i.Add(i, big.NewInt(1)) {
		if aExpI.Cmp(b) == 0 {
			return i, nil
		}
		aExpI.Mul(aExpI, a)
		aExpI.Mod(aExpI, n)
	}
	return nil, fmt.Errorf("no solution: %v^x ≡ %v (mod %v) has no solution", a, b, n)
}

// computeDiscreteLogarithmBSGS finds x such that a^x ≡ b (mod n) using Baby-step Giant-step.
// Assumes n is prime (needed for modular inverse of a^m mod n).
// Returns (x, nil) on success, or (nil, error) if no solution exists.
// Reference: https://en.wikipedia.org/wiki/Baby-step_giant-step
func computeDiscreteLogarithmBSGS(a, b, n *big.Int) (*big.Int, error) {
	// m = ceil(sqrt(n))
	m := ceilSqrt(n)
	one := big.NewInt(1)

	// --- Baby steps: build table of a^j mod n for j in [0, m) ---
	// finish Baby steps
	table := make(map[string]*big.Int)
	aExpJ := big.NewInt(1)
	for j := big.NewInt(0); j.Cmp(m) <= 0; j.Add(j, one) {
		// a^j => j
		table[aExpJ.String()] = new(big.Int).Set(j)
		aExpJ.Mul(aExpJ, a)
		aExpJ.Mod(aExpJ, n)
	}

	// --- Giant steps: for i in [0, m], check if b * (a^-m)^i mod n is in table ---
	// Compute a^m mod n, then its modular inverse: invAm = (a^m)^(-1) mod n
	// finish Giant steps
	aModM := new(big.Int).Exp(a, m, n)
	aModM.Mod(aModM, n)
	aModInverseM := new(big.Int).ModInverse(aModM, n) // a^(-m)
	if aModInverseM == nil {
		return nil, fmt.Errorf("(a^m) does not have a modulus inverse")
	}
	right := new(big.Int).Set(b) // right = a^(-m) * b
	for i := big.NewInt(0); i.Cmp(m) <= 0; i.Add(i, one) {
		if j, ok := table[right.String()]; ok {
			x := new(big.Int).Mul(i, m)
			x.Add(x, j)
			return x, nil
		}
		right.Mul(right, aModInverseM)
		right.Mod(right, n)
	}

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

// ─── Pollard's Rho for Discrete Logarithm ────────────────────────────────────
// computeDiscreteLogarithmPollardRho inds x such that a^x ≡ b (mod n) using Baby-step Giant-step.
// // Assumes n is prime (needed for modular inverse of a^m mod n).
// // Returns (x, nil) on success, or (nil, error) if no solution exists.
// // Reference: https://en.wikipedia.org/wiki/Pollard%27s_rho_algorithm_for_logarithms
func computeDiscreteLogarithmPollardRho(a, b, n *big.Int) (*big.Int, error) {
	// TODO: write your code here
	return nil, fmt.Errorf("Pollard's rho failed to converge: no collision found ")
}
