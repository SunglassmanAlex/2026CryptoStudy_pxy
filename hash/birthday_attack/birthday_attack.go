package birthday_attack

// BirthdayAttackProbability 通过蒙特卡洛模拟估计生日碰撞概率。
//
// 参数：
//   - universeSize：全域大小 N，即元素可能取值的数量（类比：哈希输出空间 2^n）。
//   - sampleSize：每次试验中有放回地随机抽取的元素个数 q。
//   - numTrials：独立重复试验的次数 t，次数越多结果越稳定。
//
// 返回值：
//   - 经验碰撞概率，即在 numTrials 次试验中发生碰撞的比例 x/t。
//     返回值范围为 [0.0, 1.0]。
//
// 碰撞的定义：一次试验中，sampleSize 个随机抽取的整数里，
// 存在至少一对取值相同的元素。
func BirthdayAttackProbability(universeSize, sampleSize, numTrials int) float64 {
	// TODO
	// 提示，你应该使用rand.IntN方法。它来自"math/rand/v2"。如果路径不对，请重新导入。
	// 参考：https://go.dev/blog/randv2
	return 0.0
}
