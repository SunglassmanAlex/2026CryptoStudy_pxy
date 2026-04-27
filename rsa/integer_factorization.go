package rsa

import (
	"math/big"
)

func f(x *big.Int) *big.Int {
	ans := new(big.Int).Set(x)
	ans.Mul(ans, x)
	ans.Add(ans, big.NewInt(1))
	return ans
}

// IntegerFactorization
// Give n, return a and b satisfies a*b=n and both a and b are not equal to 1
func IntegerFactorization(n *big.Int) (*big.Int, *big.Int) {
	// TODO
	x := big.NewInt(1)
	y := big.NewInt(1)
	one := big.NewInt(1)
	d := new(big.Int)
	for d.Cmp(one) != 0 {
		x = f(x)
		y = f(f(y))
		diff := new(big.Int).Sub(x, y)
		d.GCD(nil, nil, diff, n)
	}
	e := new(big.Int).Div(n, d)
	return d, e
}
