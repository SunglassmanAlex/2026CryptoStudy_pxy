package small_space_birthday_attack

import (
	"bytes"
	"hash"
)

// SmallSpaceBirthdayAttack 实现 Pollard's Rho 算法
func SmallSpaceBirthdayAttack(h func() hash.Hash) (x, y []byte) {
	// step 函数：执行一次哈希并返回结果
	hasher := h() // 创建实例

	step := func(input []byte) []byte {
		hasher.Reset() // 重置状态，而不是新建对象
		hasher.Write(input)
		// 建议：预先分配好一个固定长度的 slice，避免 Sum(nil) 产生新内存分配
		return hasher.Sum(nil)
	}

	// 初始种子（可以是任意值）
	start := []byte("seed")

	// --- 第一阶段：寻找环（快慢指针相遇） ---
	// TODO

	// --- 第二阶段：寻找碰撞入口 ---
	// TODO
	return nil, nil
}
