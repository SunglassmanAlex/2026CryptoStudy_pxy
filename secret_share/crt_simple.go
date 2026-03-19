package secret_share

import "math/big"

type Equation struct {
	Remainder *big.Int
	Modulus   *big.Int
}

func SolveSimpleCRT(equations []Equation) *big.Int {
	// finish SolveSimpleCRT. You can always assume that the moduli are pairwise coprime.
	mul := big.NewInt(1)
	for _, eq := range equations {
		mul.Mul(mul, eq.Modulus)
	}
	sum := big.NewInt(0)
	for _, eq := range equations {
		otherMul := new(big.Int).Div(mul, eq.Modulus)
		cur := new(big.Int).Mul(otherMul, eq.Remainder)
		t := new(big.Int).ModInverse(otherMul, eq.Modulus)
		cur.Mul(cur, t)
		sum.Add(sum, cur)
		sum.Mod(sum, mul)
	}
	return sum
}

func SolveSimple(equations []Equation) *big.Int {
	// TODO: finish SolveSimpleCRT. You should check whether modulus is prime.
	return nil
}
