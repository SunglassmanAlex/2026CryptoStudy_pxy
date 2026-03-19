package secret_share

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
)

// TestRecoverWithExactShares 测试：恰好提供 t 个份额，应该成功恢复秘密
func TestRecoverWithExactShares(t *testing.T) {
	var secret fr.Element
	secret.SetRandom() // 生成随机秘密

	n, threshold := 5, 3
	shares := SecretSplit(secret, n, threshold)

	// 取前 t (3) 个份额
	subset := shares[:threshold]
	recoveredSecret := SecretCombine(subset)

	if !recoveredSecret.Equal(&secret) {
		t.Errorf("恢复失败：期望 %s, 得到 %s", secret.String(), recoveredSecret.String())
	}
}

// TestRecoverWithMoreShares 测试：提供多于 t 个份额，也应该成功恢复秘密
func TestRecoverWithMoreShares(t *testing.T) {
	var secret fr.Element
	secret.SetUint64(123456789)

	n, threshold := 10, 4
	shares := SecretSplit(secret, n, threshold)

	// 取前 7 个份额 (n > 7 > threshold)
	subset := shares[:7]
	recoveredSecret := SecretCombine(subset)

	if !recoveredSecret.Equal(&secret) {
		t.Error("提供冗余份额时恢复秘密失败")
	}
}

// TestFailWithFewerShares 测试：提供少于 t 个份额，恢复出的秘密应该是错误的
func TestFailWithFewerShares(t *testing.T) {
	var secret fr.Element
	secret.SetRandom()

	n, threshold := 5, 3
	shares := SecretSplit(secret, n, threshold)

	// 只取 2 个份额 (少于门槛 3)
	subset := shares[:threshold-1]
	recoveredSecret := SecretCombine(subset)

	if recoveredSecret.Equal(&secret) {
		t.Error("安全漏洞：份额不足时竟然恢复了正确的秘密")
	}
}
