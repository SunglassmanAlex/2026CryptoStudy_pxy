package secret_share

import "math/big"

type Equation struct {
	Remainder *big.Int
	Modulus   *big.Int
}

func SolveSimpleCRT(equations []Equation) *big.Int {
	// TODO: finish SolveSimpleCRT. You can always assume that the moduli are pairwise coprime.
	mul := big.NewInt(1)
	n := len(equations)
	for i := 0; i < n; i++ {
		mul.Mul(mul, equations[i]->Modulus)
	}
	return nil
}
