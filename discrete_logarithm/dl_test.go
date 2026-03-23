package discrete_logarithm

import (
	"math/big"
	"testing"
)

func TestDLAlgorithm1(t *testing.T) {
	// 1e5
	n := big.NewInt(100559)
	a := big.NewInt(1926)
	b := big.NewInt(744)
	expectedX := big.NewInt(2022)
	actualX, err := ComputeDiscreteLogarithm(a, b, n)
	if err != nil {
		t.Errorf("Error in ComputeDiscreteLogarithm: %v", err)
	}
	if actualX.Cmp(expectedX) != 0 {
		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
	}
}

func TestDLAlgorithm2(t *testing.T) {
	// n ~ 1e12
	n := big.NewInt(1000000001339)
	a := big.NewInt(24551578114)
	b := big.NewInt(804730211716)
	expectedX := big.NewInt(61209541009)
	actualX, err := ComputeDiscreteLogarithm(a, b, n)
	if err != nil {
		t.Errorf("Error in ComputeDiscreteLogarithm: %v", err)
	}
	if actualX.Cmp(expectedX) != 0 {
		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
	}
}

func TestDLAlgorithm3(t *testing.T) {
	// n ~ 1e13
	n := big.NewInt(10000000000343)
	a := big.NewInt(9876543210001)
	b := big.NewInt(6205365581375)
	expectedX := big.NewInt(61209541009)
	//b := new(big.Int).Exp(a, expectedX, n)
	//fmt.Println(b)
	actualX, err := ComputeDiscreteLogarithm(a, b, n)
	if err != nil {
		t.Errorf("Error in ComputeDiscreteLogarithm: %v", err)
	}
	if actualX.Cmp(expectedX) != 0 {
		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
	}
}

func TestDLAlgorithm4(t *testing.T) {
	// n ~ 1e14
	n := big.NewInt(100000000000541)
	a := big.NewInt(189623411976)
	b := big.NewInt(8731820393861)
	expectedX := big.NewInt(111111111)
	//b := new(big.Int).Exp(a, expectedX, n)
	//fmt.Println(b)
	actualX, err := ComputeDiscreteLogarithm(a, b, n)
	if err != nil {
		t.Errorf("Error in ComputeDiscreteLogarithm: %v", err)
	}
	if actualX.Cmp(expectedX) != 0 {
		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
	}
}

func TestDLAlgorithm5(t *testing.T) {
	// n ~ 1e15
	n := big.NewInt(1000000000000741)
	a := big.NewInt(777777777777)
	b := big.NewInt(931130992341475)
	expectedX := big.NewInt(6666666666)
	//b := new(big.Int).Exp(a, expectedX, n)
	//fmt.Println(b)
	actualX, err := ComputeDiscreteLogarithm(a, b, n)
	if err != nil {
		t.Errorf("Error in ComputeDiscreteLogarithm: %v", err)
	}
	if actualX.Cmp(expectedX) != 0 {
		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
	}
}

func TestDLAlgorithm6(t *testing.T) {
	// n ~ 1e16
	n := big.NewInt(10000000000002137)
	a := big.NewInt(5566778899)
	b := big.NewInt(9615731354586746)
	expectedX := big.NewInt(11223344556677)
	actualX, err := ComputeDiscreteLogarithm(a, b, n)
	if err != nil {
		t.Errorf("Error in ComputeDiscreteLogarithm: %v", err)
	}
	if actualX.Cmp(expectedX) != 0 {
		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
	}
}

func TestDLAlgorithm7(t *testing.T) {
	// n ~ 1e17
	n := big.NewInt(100000000000001221)
	a := big.NewInt(95418289160624112)
	b := big.NewInt(33854480675090530)
	expectedX := big.NewInt(4514521452)
	//b := new(big.Int).Exp(a, expectedX, n)
	//fmt.Println(b)
	actualX, err := ComputeDiscreteLogarithm(a, b, n)
	if err != nil {
		t.Errorf("Error in ComputeDiscreteLogarithm: %v", err)
	}
	if actualX.Cmp(expectedX) != 0 {
		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
	}
}

func TestDLAlgorithm8(t *testing.T) {
	// n ~ 1e18
	n := big.NewInt(1000000000000001743)
	a := big.NewInt(542145148568468414)
	b := big.NewInt(133125798632790839)
	expectedX := big.NewInt(99189874567)
	//b := new(big.Int).Exp(a, expectedX, n)
	//fmt.Println(b)
	actualX, err := ComputeDiscreteLogarithm(a, b, n)
	if err != nil {
		t.Errorf("Error in ComputeDiscreteLogarithm: %v", err)
	}
	if actualX.Cmp(expectedX) != 0 {
		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
	}
}
