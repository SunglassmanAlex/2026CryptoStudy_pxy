package discrete_logarithm

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

// ComputeDiscreteLogarithm finds x such that a^x ≡ b (mod n) using Baby-step Giant-step.
// Assumes n is prime (needed for modular inverse of a^m mod n).
// Returns (x, nil) on success, or (nil, error) if no solution exists.
// Reference: https://en.wikipedia.org/wiki/Baby-step_giant-step
//func ComputeDiscreteLogarithm(a, b, n *big.Int) (*big.Int, error) {
//	// m = ceil(sqrt(n))
//	m := ceilSqrt(n)
//	one := big.NewInt(1)
//	//for i := big.NewInt(0); i.Cmp(n) < 0; i.Add(i, one) {
//	//	if mul.Cmp(b) == 0 {
//	//		return i, nil
//	//	}
//	//	mul.Mul(mul, a)
//	//	mul.Mod(mul, n)
//	//}
//	// --- Baby steps: build table of a^j mod n for j in [0, m) ---
//	// TODO: finish Baby steps
//	set := make(map[string]*big.Int)
//	mul := new(big.Int).Set(b)
//	for i := big.NewInt(0); i.Cmp(m) <= 0; i.Add(i, one) {
//		set[mul.String()] = new(big.Int).Set(i)
//		mul.Mul(mul, a)
//		mul.Mod(mul, n)
//	}
//	// --- Giant steps: for i in [0, m], check if b * (a^-m)^i mod n is in table ---
//	// Compute a^m mod n, then its modular inverse: invAm = (a^m)^(-1) mod n
//	// TODO: finish Giant steps
//	mul.Set(one)
//	p := new(big.Int).Set(one) // p = a^m
//	for i := big.NewInt(0); i.Cmp(m) < 0; i.Add(i, one) {
//		p.Mul(p, a)
//		p.Mod(p, n)
//	}
//
//	for i := big.NewInt(0); i.Cmp(m) <= 0; i.Add(i, one) {
//		num, ok := set[mul.String()]
//		if ok {
//			res := new(big.Int).Set(i)
//			res.Mul(res, m)
//			res.Sub(res, num)
//			return res, nil
//		}
//		mul.Mul(mul, p)
//		mul.Mod(mul, n)
//	}
//
//	return nil, fmt.Errorf("no solution: %v^x ≡ %v (mod %v) has no solution", a, b, n)
//}

type State struct {
	X, a, b *big.Int
}

func f(s *State, g, h, mod, order *big.Int) {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)
	three := big.NewInt(3)
	r := new(big.Int).Mod(s.X, three)
	if r.Cmp(zero) == 0 {
		s.X.Mul(s.X, g)
		s.X.Mod(s.X, mod)
		s.a.Add(s.a, one)
		s.a.Mod(s.a, order)
	} else if r.Cmp(one) == 0 {
		s.X.Mul(s.X, h)
		s.X.Mod(s.X, mod)
		s.b.Add(s.b, one)
		s.b.Mod(s.b, order)
	} else {
		s.X.Mul(s.X, s.X)
		s.X.Mod(s.X, mod)
		s.a.Mul(s.a, two)
		s.b.Mul(s.b, two)
		s.a.Mod(s.a, order)
		s.b.Mod(s.b, order)
	}
}

func solveLinearCongruence(r, s, m *big.Int) (ans0, g *big.Int, ok bool) {
	g = new(big.Int).GCD(nil, nil, r, m)
	rem := new(big.Int).Mod(s, g)
	if rem.Sign() != 0 {
		return nil, nil, false
	}
	r1 := new(big.Int).Div(r, g)
	s1 := new(big.Int).Div(s, g)
	m1 := new(big.Int).Div(m, g)
	inv := new(big.Int).ModInverse(r1, m1)
	if inv == nil {
		return nil, nil, false
	}
	x := new(big.Int).Mul(s1, inv)
	x.Mod(x, m1)
	return x, g, true
}

func randomState(g, h, n, order *big.Int) (State, error) {
	a0, err := rand.Int(rand.Reader, order)
	if err != nil {
		return State{}, err
	}
	b0, err := rand.Int(rand.Reader, order)
	if err != nil {
		return State{}, err
	}
	ga := new(big.Int).Exp(g, a0, n)
	hb := new(big.Int).Exp(h, b0, n)
	x := new(big.Int).Mul(ga, hb)
	x.Mod(x, n)
	return State{
		X: x,
		a: new(big.Int).Set(a0),
		b: new(big.Int).Set(b0),
	}, nil
}

func rhoStep(x, c, mod *big.Int) *big.Int {
	next := new(big.Int).Mul(x, x)
	next.Add(next, c)
	next.Mod(next, mod)
	return next
}

func factorBigInt(n *big.Int, factors map[string]int) error {
	one := big.NewInt(1)
	two := big.NewInt(2)
	if n.Cmp(one) == 0 {
		return nil
	}
	if n.ProbablyPrime(20) {
		factors[n.String()]++
		return nil
	}
	if new(big.Int).Mod(n, two).Sign() == 0 {
		factors["2"]++
		return factorBigInt(new(big.Int).Rsh(new(big.Int).Set(n), 1), factors)
	}

	max := new(big.Int).Sub(n, one)
	for attempt := 0; attempt < 32; attempt++ {
		c, err := rand.Int(rand.Reader, max)
		if err != nil {
			return err
		}
		c.Add(c, one)

		x, err := rand.Int(rand.Reader, max)
		if err != nil {
			return err
		}
		x.Add(x, one)

		y := new(big.Int).Set(x)
		d := big.NewInt(1)
		for d.Cmp(one) == 0 {
			x = rhoStep(x, c, n)
			y = rhoStep(rhoStep(y, c, n), c, n)
			diff := new(big.Int).Sub(x, y)
			diff.Abs(diff)
			d = new(big.Int).GCD(nil, nil, diff, n)
		}
		if d.Cmp(n) == 0 {
			continue
		}
		if err := factorBigInt(d, factors); err != nil {
			return err
		}
		return factorBigInt(new(big.Int).Div(new(big.Int).Set(n), d), factors)
	}

	return fmt.Errorf("failed to factor %v", n)
}

func multiplicativeOrder(base, mod *big.Int) (*big.Int, error) {
	one := big.NewInt(1)
	groupOrder := new(big.Int).Sub(mod, one)

	factors := make(map[string]int)
	if err := factorBigInt(groupOrder, factors); err != nil {
		return nil, err
	}

	order := new(big.Int).Set(groupOrder)
	for factor := range factors {
		p, ok := new(big.Int).SetString(factor, 10)
		if !ok {
			return nil, fmt.Errorf("invalid factor %q", factor)
		}
		for {
			quotient, rem := new(big.Int).QuoRem(order, p, new(big.Int))
			if rem.Sign() != 0 {
				break
			}
			if new(big.Int).Exp(base, quotient, mod).Cmp(one) != 0 {
				break
			}
			order = quotient
		}
	}

	return order, nil
}

func ComputeDiscreteLogarithmPollardRho(a, b, n *big.Int) (*big.Int, error) {
	order, err := multiplicativeOrder(a, n)
	if err != nil {
		return nil, err
	}
	fmt.Println("order", order)
	//order := new(big.Int).Sub(n, big.NewInt(1))

	for restart := 1; restart <= 1000; restart++ {
		start, err := randomState(a, b, n, order)
		if err != nil {
			return nil, err
		}

		x := State{
			X: new(big.Int).Set(start.X),
			a: new(big.Int).Set(start.a),
			b: new(big.Int).Set(start.b),
		}
		y := State{
			X: new(big.Int).Set(start.X),
			a: new(big.Int).Set(start.a),
			b: new(big.Int).Set(start.b),
		}
		maxSteps := 5 * int(ceilSqrt(order).Int64())
		for i := 1; i <= maxSteps; i++ {
			f(&x, a, b, n, order)
			f(&y, a, b, n, order)
			f(&y, a, b, n, order)
			if x.X.Cmp(y.X) == 0 {
				tmp1 := new(big.Int).Sub(x.a, y.a)
				tmp1.Mod(tmp1, order)

				tmp2 := new(big.Int).Sub(y.b, x.b)
				tmp2.Mod(tmp2, order)
				ans0, g, ok := solveLinearCongruence(tmp2, tmp1, order)
				if !ok {
					break
				}
				step := new(big.Int).Div(order, g)
				cand := new(big.Int).Set(ans0)
				var best *big.Int
				for k := big.NewInt(0); k.Cmp(g) < 0; k.Add(k, big.NewInt(1)) {
					check := new(big.Int).Exp(a, cand, n)
					if check.Cmp(b) == 0 {
						if best == nil || cand.Cmp(best) < 0 {
							best = new(big.Int).Set(cand)
						}
					}
					cand.Add(cand, step)
					cand.Mod(cand, order)
				}
				if best != nil {
					return best, nil
				}
				break
			}
		}
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
