package discrete_logarithm

import (
	"fmt"
	"math"
	"math/big"
)

type DLOption int

const (
	DLBruteForce DLOption = iota
	DLBSGS
	DLNFS
)

func ComputeDiscreteLogarithm(a, b, n *big.Int, op DLOption) (*big.Int, error) {
	switch op {
	case DLBruteForce:
		return computeDiscreteLogarithmBruteForced(a, b, n)
	case DLBSGS:
		return computeDiscreteLogarithmBSGS(a, b, n)
	case DLNFS:
		return computeDiscreteLogarithmNFS(a, b, n)
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

// ─── Index Calculus (conceptual foundation of NFS-DL) ─────────────────────────
//
// The true Number Field Sieve for DLP operates inside an algebraic number field
// and is prohibitively complex to implement from scratch. Index Calculus captures
// the identical three-phase structure — relation collection, linear algebra,
// individual logarithm — at a tractable level, and is the direct ancestor NFS
// generalises. It runs in sub-exponential time L_p[½, ½].
//
// Reference: https://en.wikipedia.org/wiki/Index_calculus_algorithm

// computeDiscreteLogarithmNFS finds x such that a^x ≡ b (mod p) via Index Calculus.
// Requires p to be prime.
//
// Phase 1 – Relation collection:
//
//	Sample r until a^r mod p factors over the factor base {p₁,…,pₖ}.
//	Each smooth value gives a linear equation over ℤ/(p−1)ℤ:
//	    e₁·log_a(p₁) + … + eₖ·log_a(pₖ) ≡ r  (mod p−1)
//
// Phase 2 – Linear algebra:
//
//	Gaussian elimination over ℤ/(p−1)ℤ recovers log_a(pᵢ) for each pᵢ.
//
// Phase 3 – Individual logarithm:
//
//	Find s ≥ 0 such that b·aˢ mod p is B-smooth, then:
//	    log_a(b) ≡ Σ eᵢ·log_a(pᵢ) − s  (mod p−1)
func computeDiscreteLogarithmNFS(a, b, p *big.Int) (*big.Int, error) {
	if !p.ProbablyPrime(20) {
		return nil, fmt.Errorf("index calculus requires a prime modulus")
	}

	ord := elementOrder(a, p)

	// Phase 0: build factor base of all primes up to the smoothness bound B.
	bound := nfsSmoothnessBound(p)
	factorBase := primeSieve(bound)
	fbSize := len(factorBase)
	if fbSize == 0 {
		return nil, fmt.Errorf("factor base is empty (bound=%d)", bound)
	}

	// Phase 1: collect at least fbSize + slack linearly independent relations.
	const slack = 20
	relMatrix, relExps, err := nfsCollectRelations(a, p, ord, factorBase, fbSize+slack)
	if err != nil {
		return nil, fmt.Errorf("relation collection: %w", err)
	}

	// Phase 2: solve for the discrete log of each factor-base element.
	fbLogs, err := nfsSolveLinearSystem(relMatrix, relExps, ord, fbSize)
	if err != nil {
		return nil, fmt.Errorf("linear algebra: %w", err)
	}

	// Phase 3: express log_a(b) in terms of the factor-base logs.
	return nfsIndividualLog(a, b, p, ord, factorBase, fbLogs)
}

// nfsSmoothnessBound returns a smoothness bound B ≈ L_p[½, ½] = exp(½·√(ln p · ln ln p)).
func nfsSmoothnessBound(p *big.Int) int64 {
	lnP := float64(p.BitLen()) * math.Log(2)
	lnlnP := math.Log(math.Max(lnP, math.E)) // guard: ln(ln p) > 0
	bound := math.Exp(math.Sqrt(lnP * lnlnP))
	switch {
	case bound < 200:
		return 200
	case bound > 500_000:
		return 500_000
	default:
		return int64(bound)
	}
}

// primeSieve returns all primes ≤ limit via the Sieve of Eratosthenes.
func primeSieve(limit int64) []int64 {
	if limit < 2 {
		return nil
	}
	composite := make([]bool, limit+1)
	var primes []int64
	for i := int64(2); i <= limit; i++ {
		if !composite[i] {
			primes = append(primes, i)
			for j := i * i; j <= limit; j += i {
				composite[j] = true
			}
		}
	}
	return primes
}

// trySmoothFactor attempts to fully factor v over factorBase by trial division.
// Returns the exponent vector if v is B-smooth, nil otherwise.
func trySmoothFactor(v *big.Int, factorBase []int64) []int64 {
	if v.Sign() <= 0 {
		return nil
	}
	rem := new(big.Int).Set(v)
	exps := make([]int64, len(factorBase))
	for i, p := range factorBase {
		pBig := big.NewInt(p)
		for new(big.Int).Mod(rem, pBig).Sign() == 0 {
			rem.Div(rem, pBig)
			exps[i]++
		}
	}
	if rem.Cmp(big.NewInt(1)) != 0 {
		return nil // a factor larger than B remains
	}
	return exps
}

// nfsCollectRelations gathers at least `needed` smooth relations of the form
//
//	a^r ≡ p₁^e₁ · p₂^e₂ · … (mod p)
//
// and returns the exponent matrix alongside the corresponding r-values.
func nfsCollectRelations(a, p, order *big.Int, factorBase []int64, needed int) ([][]int64, []*big.Int, error) {
	var matrix [][]int64
	var rVals []*big.Int

	r := big.NewInt(1)
	one := big.NewInt(1)
	maxAttempts := needed * 500

	for attempt := 0; attempt < maxAttempts && len(matrix) < needed; attempt++ {
		aExpR := new(big.Int).Exp(a, r, p)
		if exps := trySmoothFactor(aExpR, factorBase); exps != nil {
			matrix = append(matrix, exps)
			rVals = append(rVals, new(big.Int).Set(r))
		}
		r.Add(r, one)
		if r.Cmp(order) >= 0 {
			r.SetInt64(1)
		}
	}

	if len(matrix) < len(factorBase) {
		return nil, nil, fmt.Errorf("only %d/%d relations found; consider increasing the smoothness bound",
			len(matrix), len(factorBase))
	}
	return matrix, rVals, nil
}

// nfsSolveLinearSystem solves rows·x ≡ rhs (mod m) via Gaussian elimination
// over ℤ/mℤ, where m = p−1 is composite.
//
// Because ℤ/mℤ is not a field, not every column will have an invertible pivot.
// The implementation prefers rows where the column entry is coprime to m; columns
// without an invertible pivot are skipped and the corresponding unknown defaults
// to zero. The individual-log phase verifies the final answer and naturally
// compensates for any under-determined variable.
func nfsSolveLinearSystem(rows [][]int64, rhs []*big.Int, m *big.Int, cols int) ([]*big.Int, error) {
	n := len(rows)

	// Build the augmented matrix [A | b] with entries reduced mod m.
	aug := make([][]*big.Int, n)
	for i, row := range rows {
		aug[i] = make([]*big.Int, cols+1)
		for j := 0; j < cols; j++ {
			aug[i][j] = new(big.Int).Mod(big.NewInt(row[j]), m)
		}
		aug[i][cols] = new(big.Int).Mod(rhs[i], m)
	}

	one := big.NewInt(1)
	pivotOfCol := make([]int, cols) // pivotOfCol[j] = row index of pivot for column j
	for j := range pivotOfCol {
		pivotOfCol[j] = -1
	}

	cur := 0 // next unassigned pivot row
	for col := 0; col < cols && cur < n; col++ {
		// Choose the row whose entry in this column has the smallest GCD with m.
		// A GCD of 1 means the entry is invertible — the ideal pivot.
		bestRow, bestGCD := -1, new(big.Int).Set(m)
		for r := cur; r < n; r++ {
			if aug[r][col].Sign() == 0 {
				continue
			}
			g := new(big.Int).GCD(nil, nil, aug[r][col], m)
			if g.Cmp(bestGCD) < 0 {
				bestRow, bestGCD = r, new(big.Int).Set(g)
				if g.Cmp(one) == 0 {
					break // invertible pivot found; stop searching
				}
			}
		}
		if bestRow == -1 || bestGCD.Cmp(one) != 0 {
			continue // no invertible pivot in this column; skip
		}

		aug[cur], aug[bestRow] = aug[bestRow], aug[cur]

		// Normalise: scale the pivot row so aug[cur][col] ≡ 1 (mod m).
		inv := new(big.Int).ModInverse(aug[cur][col], m)
		for k := col; k <= cols; k++ {
			aug[cur][k].Mul(aug[cur][k], inv)
			aug[cur][k].Mod(aug[cur][k], m)
		}

		// Eliminate column `col` from every other row.
		for r := 0; r < n; r++ {
			if r == cur || aug[r][col].Sign() == 0 {
				continue
			}
			f := new(big.Int).Set(aug[r][col])
			for k := col; k <= cols; k++ {
				sub := new(big.Int).Mul(f, aug[cur][k])
				aug[r][k].Sub(aug[r][k], sub)
				aug[r][k].Mod(aug[r][k], m)
			}
		}

		pivotOfCol[col] = cur
		cur++
	}

	sol := make([]*big.Int, cols)
	for j := 0; j < cols; j++ {
		if pivotOfCol[j] >= 0 {
			sol[j] = new(big.Int).Set(aug[pivotOfCol[j]][cols])
		} else {
			sol[j] = big.NewInt(0) // free variable; individual-log step verifies
		}
	}
	return sol, nil
}

// nfsIndividualLog computes log_a(b) given precomputed factor-base logs.
//
// It scans s = 0, 1, 2, … until γ = b·aˢ mod p is B-smooth, then uses:
//
//	log_a(b·aˢ) ≡ Σ eᵢ·log_a(pᵢ)  (mod p−1)
//	log_a(b)    ≡ Σ eᵢ·log_a(pᵢ) − s  (mod p−1)
//
// The result is verified against the original equation before returning.
func nfsIndividualLog(a, b, p, order *big.Int, factorBase []int64, fbLogs []*big.Int) (*big.Int, error) {
	const maxAttempts = 500_000

	gamma := new(big.Int).Set(b) // b·aˢ mod p; s = 0 initially
	aModP := new(big.Int).Mod(a, p)

	for s := int64(0); s < maxAttempts; s++ {
		if exps := trySmoothFactor(gamma, factorBase); exps != nil {
			logB := big.NewInt(0)
			for i, e := range exps {
				if e == 0 {
					continue
				}
				contrib := new(big.Int).Mul(big.NewInt(e), fbLogs[i])
				logB.Add(logB, contrib)
			}
			logB.Sub(logB, big.NewInt(s))
			logB.Mod(logB, order)

			if new(big.Int).Exp(a, logB, p).Cmp(b) == 0 {
				return logB, nil
			}
			// Verification failed: the linear system left some logs undetermined.
			// Continue scanning; a different smooth value may give the right answer.
		}
		gamma.Mul(gamma, aModP)
		gamma.Mod(gamma, p)
	}

	return nil, fmt.Errorf("individual log: no B-smooth value found within %d attempts", maxAttempts)
}

// elementOrder 计算 a 在 (ℤ/pℤ)* 中的乘法阶。
// 算法：从 p-1 出发，对每个素因子 f 反复除，只要商仍是 a 的指数就继续缩小。
// p ≤ 1e10 时，对 p-1 试除到 sqrt(p-1) ≈ 1e5，速度可接受。
func elementOrder(a, p *big.Int) *big.Int {
	ord := new(big.Int).Sub(p, big.NewInt(1))
	one := big.NewInt(1)
	for _, f := range distinctPrimeFactors(ord) {
		for {
			q := new(big.Int).Div(ord, f)
			if new(big.Int).Mod(ord, f).Sign() != 0 {
				break // f 已不整除 ord
			}
			if new(big.Int).Exp(a, q, p).Cmp(one) == 0 {
				ord.Set(q) // 可以再缩小
			} else {
				break
			}
		}
	}
	return ord
}

// distinctPrimeFactors 对 n 做试除，返回不重复的素因子列表。
func distinctPrimeFactors(n *big.Int) []*big.Int {
	m := new(big.Int).Set(n)
	rem := new(big.Int)
	var factors []*big.Int
	for i := int64(2); i*i <= m.Int64(); i++ {
		bi := big.NewInt(i)
		if rem.Mod(m, bi).Sign() == 0 {
			factors = append(factors, bi)
			for rem.Mod(m, bi).Sign() == 0 {
				m.Div(m, bi)
			}
		}
	}
	if m.Cmp(big.NewInt(1)) > 0 {
		factors = append(factors, new(big.Int).Set(m))
	}
	return factors
}
