package is_prime

import (
	"math/big"
)

func IsPrime(n *big.Int) bool {
	// write a method to determine if n is a prime
	bigOne := big.NewInt(1)
	bigZero := big.NewInt(0)
	for i := big.NewInt(2); i.Cmp(n) <= 0; i.Add(i, bigOne) {
		r := new(big.Int).Mod(n, i)
		tmp := new(big.Int).Mul(i, i)
		if tmp.Cmp(n) > 0 {
			break
		}
		if r.Cmp(bigZero) == 0 {
			return false
		}
	}
	return true
}
