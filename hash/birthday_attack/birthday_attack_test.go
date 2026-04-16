package birthday_attack

import (
	"fmt"
	"math"
	"testing"
)

func TestBirthdayAttackProbability(t *testing.T) {
	// 示例：
	//	经典生日问题：365天，随即选取23人即有概率得到碰撞。
	fmt.Printf("经典生日问题：365天，随即选取23人即有概率得到碰撞。%f\n",
		BirthdayAttackProbability(365, 23, 100000))
	//	模拟 16-bit 截断哈希的生日攻击，305次即有一半概率得到碰撞
	fmt.Printf("模拟 16-bit 截断哈希的生日攻击，305次即有一半概率得到碰撞。%f\n",
		BirthdayAttackProbability(1<<16, 305, 100000))
	//	模拟 24-bit 截断哈希的生日攻击，4850次即有一半概率得到碰撞
	fmt.Printf("模拟 24-bit 截断哈希的生日攻击，4850次即有一半概率得到碰撞。%f\n",
		BirthdayAttackProbability(1<<24, 4850, 100000))
}

// tolerance 是模拟值与期望值之间允许的最大绝对误差。
// 10万次试验下，蒙特卡洛标准误差约为 0.5/√100000 ≈ 0.0016，
// 设置 0.03 的容差已留有足够余量。
const tolerance = 0.03

// exactCollisionProbability 用精确的连乘公式计算碰撞概率，
// 作为测试的参考期望值（不依赖近似）：
//
//	P(collision) = 1 - ∏_{i=0}^{q-1} (N-i)/N
//
// 该公式对任意 N 均精确，适合 N 较小时替代指数近似。
func exactCollisionProbability(universeSize, sampleSize int) float64 {
	n := float64(universeSize)
	p := 1.0
	for i := 0; i < sampleSize; i++ {
		p *= (n - float64(i)) / n
	}
	return 1 - p
}

// withinTolerance 判断两个概率值之差是否在允许误差内。
func withinTolerance(got, want float64) bool {
	return math.Abs(got-want) <= tolerance
}

// ----- 核心功能测试 -----

// TestClassicBirthdayProblem 验证最著名的生日问题实例：
// 365天，23人时碰撞概率约为 50.7%。
func TestClassicBirthdayProblem(t *testing.T) {
	got := BirthdayAttackProbability(365, 23, 100_000)
	want := exactCollisionProbability(365, 23) // ≈ 0.5073

	if !withinTolerance(got, want) {
		t.Errorf("Classic birthday: got %.4f, want ≈ %.4f (±%.2f)", got, want, tolerance)
	}
}

// TestNearCertainCollision 验证当样本量足够大时碰撞概率接近 1。
// 365天取 57 人，精确概率约 99%。
func TestNearCertainCollision(t *testing.T) {
	got := BirthdayAttackProbability(365, 57, 100_000)
	want := exactCollisionProbability(365, 57) // ≈ 0.9901

	if !withinTolerance(got, want) {
		t.Errorf("Near-certain collision: got %.4f, want ≈ %.4f (±%.2f)", got, want, tolerance)
	}
}

// TestSmallUniverse 使用极小全域（N=10）进行验证。
// 此时指数近似误差较大，故使用精确公式作为期望值。
// q=5 时精确值约为 69.8%（由精确公式计算，非近似）。
func TestSmallUniverse(t *testing.T) {
	got := BirthdayAttackProbability(10, 5, 200_000)
	// 精确值：1 - (10*9*8*7*6)/10^5 = 1 - 30240/100000 ≈ 0.6976
	want := exactCollisionProbability(10, 5)

	if !withinTolerance(got, want) {
		t.Errorf("Small universe N=10, q=5: got %.4f, want ≈ %.4f (±%.2f)", got, want, tolerance)
	}
}

// TestHashBirthdayAttack_16bit 模拟 16-bit 截断哈希的生日攻击场景。
// N = 2^16 = 65536，q = 300 时碰撞概率约 50.8%。
func TestHashBirthdayAttack_16bit(t *testing.T) {
	got := BirthdayAttackProbability(1<<16, 300, 50_000)
	want := exactCollisionProbability(1<<16, 300)

	if !withinTolerance(got, want) {
		t.Errorf("16-bit hash attack: got %.4f, want ≈ %.4f (±%.2f)", got, want, tolerance)
	}
}

// TestHalfwayPoint 验证"生日攻击平衡点"：
// q ≈ √N 时碰撞概率应接近 50%。
// 此测试直接体现生日攻击将安全位数折半的核心结论。
func TestHalfwayPoint(t *testing.T) {
	universeSize := 10_000
	// q ≈ √N = 100，使得碰撞概率接近 50%
	sampleSize := int(math.Round(math.Sqrt(float64(universeSize))))

	got := BirthdayAttackProbability(universeSize, sampleSize, 100_000)
	want := exactCollisionProbability(universeSize, sampleSize)

	if !withinTolerance(got, want) {
		t.Errorf("Halfway point (q≈√N): got %.4f, want ≈ %.4f (±%.2f)", got, want, tolerance)
	}

	// 额外断言：碰撞概率应落在 [0.35, 0.65] 的合理区间内
	if got < 0.35 || got > 0.65 {
		t.Errorf("Halfway point: expected probability near 0.50, got %.4f", got)
	}
}

// TestSampleEqualsUniverse 验证 sampleSize == universeSize 时碰撞概率极接近 1。
// 由鸽巢原理，从 N 个槽位中有放回取 N 个元素，碰撞概率非常高（但非严格 100%）。
func TestSampleEqualsUniverse(t *testing.T) {
	got := BirthdayAttackProbability(50, 50, 50_000)
	if got < 0.95 {
		t.Errorf("sampleSize==universeSize: expected p≥0.95, got %.4f", got)
	}
}

// ----- TheoreticalCollisionProbability 测试 -----

// TestTheoreticalVsExact 验证 TheoreticalCollisionProbability（指数近似）
// 与精确公式的差距在大 N 时小于 0.002，在小 N 时误差更大属正常现象。
//func TestTheoreticalVsExact(t *testing.T) {
//	cases := []struct {
//		universeSize, sampleSize int
//		maxError                 float64
//		desc                     string
//	}{
//		// 大 N：近似精度高
//		{365, 23, 0.01, "经典生日问题"},
//		{365, 57, 0.003, "365天 57人"},
//		{1 << 16, 300, 0.002, "16-bit 哈希空间"},
//		// sampleSize=1：精确值和近似值均为 0
//		{1000, 1, 0.0, "sampleSize=1"},
//	}
//
//	for _, tc := range cases {
//		approx := TheoreticalCollisionProbability(tc.universeSize, tc.sampleSize)
//		exact := exactCollisionProbability(tc.universeSize, tc.sampleSize)
//		diff := math.Abs(approx - exact)
//		if diff > tc.maxError+0.001 {
//			t.Errorf("TheoreticalCollisionProbability(%d,%d) [%s]: approx=%.5f exact=%.5f diff=%.5f > maxError=%.3f",
//				tc.universeSize, tc.sampleSize, tc.desc, approx, exact, diff, tc.maxError)
//		}
//	}
//}
//
//// TestTheoreticalInvalidParameters 验证 TheoreticalCollisionProbability 的参数校验。
//func TestTheoreticalInvalidParameters(t *testing.T) {
//	if got := TheoreticalCollisionProbability(0, 10); got != -1 {
//		t.Errorf("universeSize=0: expected -1, got %.4f", got)
//	}
//	if got := TheoreticalCollisionProbability(100, -1); got != -1 {
//		t.Errorf("sampleSize<0: expected -1, got %.4f", got)
//	}
//}
