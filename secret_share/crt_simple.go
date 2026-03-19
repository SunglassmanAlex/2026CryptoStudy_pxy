package secret_share

import "math/big"

type Equation struct {
	Remainder *big.Int
	Modulus   *big.Int
}

func SolveSimpleCRT(equations []Equation) *big.Int {
	// TODO: finish SolveSimpleCRT. You can always assume that the moduli are pairwise coprime.
	return nil
}

func SolveCRT(equations []Equation) *big.Int {
	// TODO: finish SolveSimpleCRT. You should check whether modulus is prime.
	return nil
}
