package discrete_logarithm

import (
	"fmt"
	"math/big"
	"testing"
)

func TestDiscreteLogarithmBF(t *testing.T) {
	// 定义测试用例集
	tests := []struct {
		name      string
		n         *big.Int
		a         *big.Int
		b         *big.Int
		expectedX *big.Int
	}{
		{"1e5", big.NewInt(100559), big.NewInt(1926), big.NewInt(744), big.NewInt(2022)},
		{"1e6", big.NewInt(1000183), big.NewInt(1874), big.NewInt(705290), big.NewInt(111111)},
		{"1e7", big.NewInt(10000103), big.NewInt(1874), big.NewInt(2897553), big.NewInt(111111)},
		{"1e8", big.NewInt(100000007), big.NewInt(18741), big.NewInt(65235972), big.NewInt(11111111)},
		{"1e9", big.NewInt(1000000007), big.NewInt(18741), big.NewInt(928049376), big.NewInt(111111111)},
		{"1e10", big.NewInt(10000000147), big.NewInt(187412), big.NewInt(2483365300), big.NewInt(1111111)},
		{"1e11", big.NewInt(100000000129), big.NewInt(1187412), big.NewInt(73008291574), big.NewInt(111111111)},
		{"1e12", big.NewInt(1000000000189), big.NewInt(666666667), big.NewInt(616642048409), big.NewInt(666666666)},
		{"1e13", big.NewInt(10000000000343), big.NewInt(9876543210001), big.NewInt(6205365581375), big.NewInt(61209541009)},
		{"1e14", big.NewInt(100000000000541), big.NewInt(189623411976), big.NewInt(8731820393861), big.NewInt(111111111)},
		{"1e15", big.NewInt(1000000000000741), big.NewInt(777777777777), big.NewInt(931130992341475), big.NewInt(6666666666)},
		{"1e16", big.NewInt(10000000000002137), big.NewInt(5566778899), big.NewInt(9615731354586746), big.NewInt(11223344556677)},
		{"1e17", big.NewInt(100000000000001221), big.NewInt(95418289160624112), big.NewInt(33854480675090530), big.NewInt(4514521452)},
		{"1e18", big.NewInt(1000000000000001743), big.NewInt(542145148568468414), big.NewInt(133125798632790839), big.NewInt(99189874567)},
	}

	for _, tt := range tests {
		// 使用 t.Run 创建子测试条目
		t.Run(fmt.Sprintf("Scale_%s", tt.name), func(t *testing.T) {
			actualX, err := ComputeDiscreteLogarithm(tt.a, tt.b, tt.n, DLBruteForce)

			if err != nil {
				t.Fatalf("算法执行出错: %v", err)
			}

			if actualX.Cmp(tt.expectedX) != 0 {
				t.Errorf("\n结果不匹配!\n期望: %v\n实际: %v", tt.expectedX, actualX)
			}
		})
	}
}

func TestDiscreteLogarithmBSGS(t *testing.T) {
	// 定义测试用例集
	tests := []struct {
		name      string
		n         *big.Int
		a         *big.Int
		b         *big.Int
		expectedX *big.Int
	}{
		{"1e5", big.NewInt(100559), big.NewInt(1926), big.NewInt(744), big.NewInt(2022)},
		{"1e6", big.NewInt(1000183), big.NewInt(1874), big.NewInt(705290), big.NewInt(111111)},
		{"1e7", big.NewInt(10000103), big.NewInt(1874), big.NewInt(2897553), big.NewInt(111111)},
		{"1e8", big.NewInt(100000007), big.NewInt(18741), big.NewInt(65235972), big.NewInt(11111111)},
		{"1e9", big.NewInt(1000000007), big.NewInt(18741), big.NewInt(928049376), big.NewInt(111111111)},
		{"1e10", big.NewInt(10000000147), big.NewInt(187412), big.NewInt(2483365300), big.NewInt(1111111)},
		{"1e11", big.NewInt(100000000129), big.NewInt(1187412), big.NewInt(73008291574), big.NewInt(111111111)},
		{"1e12", big.NewInt(1000000000189), big.NewInt(666666667), big.NewInt(616642048409), big.NewInt(666666666)},
		{"1e13", big.NewInt(10000000000343), big.NewInt(9876543210001), big.NewInt(6205365581375), big.NewInt(61209541009)},
		{"1e14", big.NewInt(100000000000541), big.NewInt(189623411976), big.NewInt(8731820393861), big.NewInt(111111111)},
		{"1e15", big.NewInt(1000000000000741), big.NewInt(777777777777), big.NewInt(931130992341475), big.NewInt(6666666666)},
		{"1e16", big.NewInt(10000000000002137), big.NewInt(5566778899), big.NewInt(9615731354586746), big.NewInt(11223344556677)},
		{"1e17", big.NewInt(100000000000001221), big.NewInt(95418289160624112), big.NewInt(33854480675090530), big.NewInt(4514521452)},
		{"1e18", big.NewInt(1000000000000001743), big.NewInt(542145148568468414), big.NewInt(133125798632790839), big.NewInt(99189874567)},
	}

	for _, tt := range tests {
		// 使用 t.Run 创建子测试条目
		t.Run(fmt.Sprintf("Scale_%s", tt.name), func(t *testing.T) {
			actualX, err := ComputeDiscreteLogarithm(tt.a, tt.b, tt.n, DLBSGS)

			if err != nil {
				t.Fatalf("算法执行出错: %v", err)
			}

			if actualX.Cmp(tt.expectedX) != 0 {
				t.Errorf("\n结果不匹配!\n期望: %v\n实际: %v", tt.expectedX, actualX)
			}
		})
	}
}

func TestDiscreteLogarithmPR(t *testing.T) {
	// 定义测试用例集
	tests := []struct {
		name      string
		n         *big.Int
		a         *big.Int
		b         *big.Int
		expectedX *big.Int
	}{
		{"1e5", big.NewInt(100559), big.NewInt(1926), big.NewInt(744), big.NewInt(2022)},
		{"1e6", big.NewInt(1000183), big.NewInt(1874), big.NewInt(705290), big.NewInt(111111)},
		{"1e7", big.NewInt(10000103), big.NewInt(1874), big.NewInt(2897553), big.NewInt(111111)},
		{"1e8", big.NewInt(100000007), big.NewInt(18741), big.NewInt(65235972), big.NewInt(11111111)},
		{"1e9", big.NewInt(1000000007), big.NewInt(18741), big.NewInt(928049376), big.NewInt(111111111)},
		{"1e10", big.NewInt(10000000147), big.NewInt(187412), big.NewInt(2483365300), big.NewInt(1111111)},
		{"1e11", big.NewInt(100000000129), big.NewInt(1187412), big.NewInt(73008291574), big.NewInt(111111111)},
		{"1e12", big.NewInt(1000000000189), big.NewInt(666666667), big.NewInt(616642048409), big.NewInt(666666666)},
		{"1e13", big.NewInt(10000000000343), big.NewInt(9876543210001), big.NewInt(6205365581375), big.NewInt(61209541009)},
		{"1e14", big.NewInt(100000000000541), big.NewInt(189623411976), big.NewInt(8731820393861), big.NewInt(111111111)},
		{"1e15", big.NewInt(1000000000000741), big.NewInt(777777777777), big.NewInt(931130992341475), big.NewInt(6666666666)},
		{"1e16", big.NewInt(10000000000002137), big.NewInt(5566778899), big.NewInt(9615731354586746), big.NewInt(11223344556677)},
		{"1e17", big.NewInt(100000000000001221), big.NewInt(95418289160624112), big.NewInt(33854480675090530), big.NewInt(4514521452)},
		{"1e18", big.NewInt(1000000000000001743), big.NewInt(542145148568468414), big.NewInt(133125798632790839), big.NewInt(99189874567)},
	}

	for _, tt := range tests {
		// 使用 t.Run 创建子测试条目
		t.Run(fmt.Sprintf("Scale_%s", tt.name), func(t *testing.T) {
			actualX, err := ComputeDiscreteLogarithm(tt.a, tt.b, tt.n, DLPollardRho)

			if err != nil {
				t.Fatalf("算法执行出错: %v", err)
			}

			if actualX.Cmp(tt.expectedX) != 0 {
				t.Errorf("\n结果不匹配!\n期望: %v\n实际: %v", tt.expectedX, actualX)
			}
		})
	}
}

func TestDiscreteLogarithmNFS(t *testing.T) {
	// 定义测试用例集
	tests := []struct {
		name      string
		n         *big.Int
		a         *big.Int
		b         *big.Int
		expectedX *big.Int
	}{
		{"1e5", big.NewInt(100559), big.NewInt(1926), big.NewInt(744), big.NewInt(2022)},
		{"1e6", big.NewInt(1000183), big.NewInt(1874), big.NewInt(705290), big.NewInt(111111)},
		{"1e7", big.NewInt(10000103), big.NewInt(1874), big.NewInt(2897553), big.NewInt(111111)},
		{"1e8", big.NewInt(100000007), big.NewInt(18741), big.NewInt(65235972), big.NewInt(11111111)},
		{"1e9", big.NewInt(1000000007), big.NewInt(18741), big.NewInt(928049376), big.NewInt(111111111)},
		{"1e10", big.NewInt(10000000147), big.NewInt(187412), big.NewInt(2483365300), big.NewInt(1111111)},
		{"1e11", big.NewInt(100000000129), big.NewInt(1187412), big.NewInt(73008291574), big.NewInt(111111111)},
		{"1e12", big.NewInt(1000000000189), big.NewInt(666666667), big.NewInt(616642048409), big.NewInt(666666666)},
		{"1e13", big.NewInt(10000000000343), big.NewInt(9876543210001), big.NewInt(6205365581375), big.NewInt(61209541009)},
		{"1e14", big.NewInt(100000000000541), big.NewInt(189623411976), big.NewInt(8731820393861), big.NewInt(111111111)},
		{"1e15", big.NewInt(1000000000000741), big.NewInt(777777777777), big.NewInt(931130992341475), big.NewInt(6666666666)},
		{"1e16", big.NewInt(10000000000002137), big.NewInt(5566778899), big.NewInt(9615731354586746), big.NewInt(11223344556677)},
		{"1e17", big.NewInt(100000000000001221), big.NewInt(95418289160624112), big.NewInt(33854480675090530), big.NewInt(4514521452)},
		{"1e18", big.NewInt(1000000000000001743), big.NewInt(542145148568468414), big.NewInt(133125798632790839), big.NewInt(99189874567)},
	}

	for _, tt := range tests {
		// 使用 t.Run 创建子测试条目
		t.Run(fmt.Sprintf("Scale_%s", tt.name), func(t *testing.T) {
			actualX, err := ComputeDiscreteLogarithm(tt.a, tt.b, tt.n, DLNFS)

			if err != nil {
				t.Fatalf("算法执行出错: %v", err)
			}

			if actualX.Cmp(tt.expectedX) != 0 {
				t.Errorf("\n结果不匹配!\n期望: %v\n实际: %v", tt.expectedX, actualX)
			}
		})
	}
}

//package discrete_logarithm
//
//import (
//	"math/big"
//	"testing"
//)
//
//func TestDLAlgorithm1e5(t *testing.T) {
//	// 1e5
//	n := big.NewInt(100559)
//	a := big.NewInt(1926)
//	b := big.NewInt(744)
//	expectedX := big.NewInt(2022)
//	actualX, err := computeDiscreteLogarithmBSGS(a, b, n)
//	if err != nil {
//		t.Errorf("Error in computeDiscreteLogarithmBSGS: %v", err)
//	}
//	if actualX.Cmp(expectedX) != 0 {
//		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
//	}
//}
//
//func TestDLAlgorithm1e6(t *testing.T) {
//	// n ~ 1e6
//	n := big.NewInt(1000183)
//	a := big.NewInt(1874)
//	b := big.NewInt(705290)
//	expectedX := big.NewInt(111111)
//	//b := new(big.Int).Exp(a, expectedX, n)
//	//fmt.Println(b)
//	actualX, err := computeDiscreteLogarithmBSGS(a, b, n)
//	if err != nil {
//		t.Errorf("Error in computeDiscreteLogarithmBSGS: %v", err)
//	}
//	if actualX.Cmp(expectedX) != 0 {
//		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
//	}
//}
//
//func TestDLAlgorithm1e7(t *testing.T) {
//	// n ~ 1e7
//	n := big.NewInt(10000103)
//	a := big.NewInt(1874)
//	b := big.NewInt(2897553)
//	expectedX := big.NewInt(111111)
//	//b := new(big.Int).Exp(a, expectedX, n)
//	//fmt.Println(b)
//	actualX, err := computeDiscreteLogarithmBSGS(a, b, n)
//	if err != nil {
//		t.Errorf("Error in computeDiscreteLogarithmBSGS: %v", err)
//	}
//	if actualX.Cmp(expectedX) != 0 {
//		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
//	}
//}
//
//func TestDLAlgorithm1e8(t *testing.T) {
//	// n ~ 1e8
//	n := big.NewInt(100000007)
//	a := big.NewInt(18741)
//	b := big.NewInt(65235972)
//	expectedX := big.NewInt(11111111)
//	//b := new(big.Int).Exp(a, expectedX, n)
//	//fmt.Println(b)
//	actualX, err := computeDiscreteLogarithmBSGS(a, b, n)
//	if err != nil {
//		t.Errorf("Error in computeDiscreteLogarithmBSGS: %v", err)
//	}
//	if actualX.Cmp(expectedX) != 0 {
//		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
//	}
//}
//
//func TestDLAlgorithm1e9(t *testing.T) {
//	// n ~ 1e9
//	n := big.NewInt(1000000007)
//	a := big.NewInt(18741)
//	b := big.NewInt(928049376)
//	expectedX := big.NewInt(111111111)
//	//b := new(big.Int).Exp(a, expectedX, n)
//	//fmt.Println(b)
//	actualX, err := computeDiscreteLogarithmBSGS(a, b, n)
//	if err != nil {
//		t.Errorf("Error in computeDiscreteLogarithmBSGS: %v", err)
//	}
//	if actualX.Cmp(expectedX) != 0 {
//		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
//	}
//}
//
//func TestDLAlgorithm1e10(t *testing.T) {
//	// n ~ 1e10
//	n := big.NewInt(10000000147)
//	a := big.NewInt(187412)
//	b := big.NewInt(2483365300)
//	expectedX := big.NewInt(1111111)
//	//b := new(big.Int).Exp(a, expectedX, n)
//	//b.Mod(b, n)
//	//fmt.Println(b)
//	actualX, err := computeDiscreteLogarithmBSGS(a, b, n)
//	if err != nil {
//		t.Errorf("Error in computeDiscreteLogarithmBSGS: %v", err)
//	}
//	if actualX.Cmp(expectedX) != 0 {
//		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
//	}
//}
//
//func TestDLAlgorithm1e11(t *testing.T) {
//	// n ~ 1e11
//	n := big.NewInt(100000000129)
//	a := big.NewInt(1187412)
//	b := big.NewInt(73008291574)
//	expectedX := big.NewInt(111111111)
//	//b := new(big.Int).Exp(a, expectedX, n)
//	//b.Mod(b, n)
//	//fmt.Println(b)
//	actualX, err := computeDiscreteLogarithmBSGS(a, b, n)
//	if err != nil {
//		t.Errorf("Error in computeDiscreteLogarithmBSGS: %v", err)
//	}
//	if actualX.Cmp(expectedX) != 0 {
//		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
//	}
//}
//
//func TestDLAlgorithm1e12(t *testing.T) {
//	// n ~ 1e12
//	n := big.NewInt(1000000000189)
//	a := big.NewInt(666666667)
//	b := big.NewInt(616642048409)
//	expectedX := big.NewInt(666666666)
//	//b := new(big.Int).Exp(a, expectedX, n)
//	//b.Mod(b, n)
//	//fmt.Println(b)
//	actualX, err := computeDiscreteLogarithmBSGS(a, b, n)
//	if err != nil {
//		t.Errorf("Error in computeDiscreteLogarithmBSGS: %v", err)
//	}
//	if actualX.Cmp(expectedX) != 0 {
//		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
//	}
//}
//
//func TestDLAlgorithm1e13(t *testing.T) {
//	// n ~ 1e13
//	n := big.NewInt(10000000000343)
//	a := big.NewInt(9876543210001)
//	b := big.NewInt(6205365581375)
//	expectedX := big.NewInt(61209541009)
//	//b := new(big.Int).Exp(a, expectedX, n)
//	//fmt.Println(b)
//	actualX, err := computeDiscreteLogarithmBSGS(a, b, n)
//	if err != nil {
//		t.Errorf("Error in computeDiscreteLogarithmBSGS: %v", err)
//	}
//	if actualX.Cmp(expectedX) != 0 {
//		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
//	}
//}
//
//func TestDLAlgorithm1e14(t *testing.T) {
//	// n ~ 1e14
//	n := big.NewInt(100000000000541)
//	a := big.NewInt(189623411976)
//	b := big.NewInt(8731820393861)
//	expectedX := big.NewInt(111111111)
//	//b := new(big.Int).Exp(a, expectedX, n)
//	//fmt.Println(b)
//	actualX, err := computeDiscreteLogarithmBSGS(a, b, n)
//	if err != nil {
//		t.Errorf("Error in computeDiscreteLogarithmBSGS: %v", err)
//	}
//	if actualX.Cmp(expectedX) != 0 {
//		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
//	}
//}
//
//func TestDLAlgorithm1e15(t *testing.T) {
//	// n ~ 1e15
//	n := big.NewInt(1000000000000741)
//	a := big.NewInt(777777777777)
//	b := big.NewInt(931130992341475)
//	expectedX := big.NewInt(6666666666)
//	//b := new(big.Int).Exp(a, expectedX, n)
//	//fmt.Println(b)
//	actualX, err := computeDiscreteLogarithmBSGS(a, b, n)
//	if err != nil {
//		t.Errorf("Error in computeDiscreteLogarithmBSGS: %v", err)
//	}
//	if actualX.Cmp(expectedX) != 0 {
//		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
//	}
//}
//
//func TestDLAlgorithm1e16(t *testing.T) {
//	// n ~ 1e16
//	n := big.NewInt(10000000000002137)
//	a := big.NewInt(5566778899)
//	b := big.NewInt(9615731354586746)
//	expectedX := big.NewInt(11223344556677)
//	actualX, err := computeDiscreteLogarithmBSGS(a, b, n)
//	if err != nil {
//		t.Errorf("Error in computeDiscreteLogarithmBSGS: %v", err)
//	}
//	if actualX.Cmp(expectedX) != 0 {
//		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
//	}
//}
//
//func TestDLAlgorithm1e17(t *testing.T) {
//	// n ~ 1e17
//	n := big.NewInt(100000000000001221)
//	a := big.NewInt(95418289160624112)
//	b := big.NewInt(33854480675090530)
//	expectedX := big.NewInt(4514521452)
//	//b := new(big.Int).Exp(a, expectedX, n)
//	//fmt.Println(b)
//	actualX, err := computeDiscreteLogarithmBSGS(a, b, n)
//	if err != nil {
//		t.Errorf("Error in computeDiscreteLogarithmBSGS: %v", err)
//	}
//	if actualX.Cmp(expectedX) != 0 {
//		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
//	}
//}
//
//func TestDLAlgorithm1e18(t *testing.T) {
//	// n ~ 1e18
//	n := big.NewInt(1000000000000001743)
//	a := big.NewInt(542145148568468414)
//	b := big.NewInt(133125798632790839)
//	expectedX := big.NewInt(99189874567)
//	//b := new(big.Int).Exp(a, expectedX, n)
//	//fmt.Println(b)
//	actualX, err := computeDiscreteLogarithmBSGS(a, b, n)
//	if err != nil {
//		t.Errorf("Error in computeDiscreteLogarithmBSGS: %v", err)
//	}
//	if actualX.Cmp(expectedX) != 0 {
//		t.Errorf("expected: %v, actual: %v", expectedX, actualX)
//	}
//}
