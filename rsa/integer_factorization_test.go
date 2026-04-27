package rsa

import (
	"math/big"
	"testing"
)

func TestIntegerFactorization(t *testing.T) {
	tests := []struct {
		name string
		n    *big.Int
	}{
		{
			name: "Test1:1e6",
			n:    big.NewInt(4080319),
		},
		{
			name: "Test2:1e7",
			n:    big.NewInt(12594917),
		},
		{
			name: "Test3:1e8",
			n:    big.NewInt(360004171),
		},
		{
			name: "Test4:1e9",
			n:    big.NewInt(1361728873),
		},
		{
			name: "Test5:1e10",
			n:    big.NewInt(11994220423),
		},
		{
			name: "Test6:1e11",
			n:    big.NewInt(872461181791),
		},
		{
			name: "Test7:1e12",
			n:    big.NewInt(1000582001737),
		},
		{
			name: "Test8:1e13",
			n:    big.NewInt(10006529088821),
		},
		{
			name: "Test9:1e14",
			n:    big.NewInt(100011980296301),
		},
		{
			name: "Test10:1e15",
			n:    big.NewInt(1000027530056419),
		},
		{
			name: "Test11:1e16",
			n:    big.NewInt(10000075200127687),
		},
		{
			name: "Test12:1e17",
			n:    big.NewInt(100000052300001617),
		},
		{
			name: "Test13:1e18",
			n:    big.NewInt(1000000820300008343),
		},
		{
			name: "Test14:1e19",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("10000001377000018081", 10); return v }(),
		},
		{
			name: "Test15:1e20",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("100000018271000030951", 10); return v }(),
		},
		{
			name: "Test16:1e21",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("1000000001930000000057", 10); return v }(),
		},
		{
			name: "Test17:1e22",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("1000000013810000020691", 10); return v }(),
		},
		{
			name: "Test18:1e23",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("1000000010090000022011", 10); return v }(),
		},
		{
			name: "Test19:1e24",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("1000000000100000000002379", 10); return v }(),
		},
		{
			name: "Test20:1e25",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("100000000003580000000025641", 10); return v }(),
		},
		{
			name: "Test21:1e26",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("1000000000004370000000002479", 10); return v }(),
		},
		{
			name: "Test22:1e27",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("10000000000039400000000034713", 10); return v }(),
		},
		{
			name: "Test23:1e28",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("10000000000046400000000035599", 10); return v }(),
		},
		{
			name: "Test24:1e29",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("1000000000000314000000000020293", 10); return v }(),
		},
		{
			name: "Test25:1e30",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("1000000000000128000000000003367", 10); return v }(),
		},
		{
			name: "Test26:1e31",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("100000000000011260000000000314469", 10); return v }(),
		},
		{
			name: "Test27:1e32",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("100000000000001300000000000004209", 10); return v }(),
		},
		{
			name: "Test28:1e33",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("100000000000015000000000000550179", 10); return v }(),
		},
		{
			name: "Test29:1e34",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("10000000000000001600000000000000039", 10); return v }(),
		},
		{
			name: "Test30:1e35",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("10000000000000006800000000000000931", 10); return v }(),
		},
		{
			name: "Test31:1e36",
			n:    func() *big.Int { v, _ := new(big.Int).SetString("1000000000000000034000000000000000093", 10); return v }(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p, q := IntegerFactorization(tc.n)
			actual := new(big.Int)
			if actual.Mul(p, q).Cmp(tc.n) != 0 {
				t.Errorf("%s * %s got %s, want %s", p, q, actual, tc.n)
			}
		})
	}
}
