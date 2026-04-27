package rsa

import (
	"math/big"
	"math/rand"
)

var one = big.NewInt(1)

// f computes (x*x + c) mod n
func f(x, c, n *big.Int) *big.Int {
	x2 := new(big.Int).Mul(x, x)
	x2.Add(x2, c)
	return x2.Mod(x2, n)
}

// pollardRho finds a non-trivial factor of n using Pollard's rho algorithm.
func pollardRho(n *big.Int) *big.Int {
	for {
		x := big.NewInt(rand.Int63n(100000) + 2)
		c := big.NewInt(rand.Int63n(100000) + 1)
		y := new(big.Int).Set(x)
		d := new(big.Int).Set(one)

		for d.Cmp(one) == 0 {
			x = f(x, c, n)
			y = f(f(y, c, n), c, n)

			sub := new(big.Int).Sub(x, y)
			sub.Abs(sub)
			d.GCD(nil, nil, sub, n)
		}

		if d.Cmp(n) != 0 {
			return d
		}
	}
}

// IntegerFactorization
// Give n, return a and b satisfies a*b=n and both a and b are not equal to 1
func IntegerFactorization(n *big.Int) (*big.Int, *big.Int) {
	if n.Bit(0) == 0 {
		return big.NewInt(2), new(big.Int).Div(n, big.NewInt(2))
	}

	a := pollardRho(n)
	b := new(big.Int).Div(n, a)
	return a, b
}
