package small_space_birthday_attack

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"testing"
)

// truncatedHash 包装了底层的 Hash，只取最后的 N 个字节
type truncatedHash struct {
	hash.Hash
	size int
}

func (t *truncatedHash) Sum(b []byte) []byte {
	fullHash := t.Hash.Sum(nil)
	// 取后 size 个字节 (例如 40位 = 5字节)
	return fullHash[len(fullHash)-t.size:]
}

// 构造函数，满足 func() hash.Hash 签名
func Newtruncatedsha25640() hash.Hash {
	return &truncatedHash{
		Hash: sha256.New(),
		size: 5, // 40 bits = 5 bytes
	}
}

func Newtruncatedsha25648() hash.Hash {
	return &truncatedHash{
		Hash: sha256.New(),
		size: 6, // 48 bits = 6 bytes
	}
}

func Newtruncatedsha25656() hash.Hash {
	return &truncatedHash{
		Hash: sha256.New(),
		size: 7, // 54 bits = 7 bytes
	}
}

func Newtruncatedsha25664() hash.Hash {
	return &truncatedHash{
		Hash: sha256.New(),
		size: 8, // 64 bits = 8 bytes
	}
}

func TestSHA256_40(t *testing.T) {
	fmt.Println("正在尝试对 SHA-256 (截断至后40位) 进行碰撞攻击...")

	x, y := SmallSpaceBirthdayAttack(Newtruncatedsha25640)

	// 验证结果
	h1 := Newtruncatedsha25640()
	h1.Write(x)
	resX := h1.Sum(nil)

	h2 := Newtruncatedsha25640()
	h2.Write(y)
	resY := h2.Sum(nil)

	fmt.Printf("找到碰撞！\n")
	fmt.Printf("x (hex): %s\n", hex.EncodeToString(x))
	fmt.Printf("y (hex): %s\n", hex.EncodeToString(y))
	fmt.Printf("H(x) (hex): %s\n", hex.EncodeToString(resX))
	fmt.Printf("H(y) (hex): %s\n", hex.EncodeToString(resY))

	if bytes.Equal(resX, resY) && !bytes.Equal(x, y) {
		fmt.Println("验证成功：输入不同且哈希值相等。")
	} else {
		fmt.Println("验证失败。")
	}
}

func TestSHA256_48(t *testing.T) {
	fmt.Println("正在尝试对 SHA-256 (截断至后48位) 进行碰撞攻击...")

	x, y := SmallSpaceBirthdayAttack(Newtruncatedsha25648)

	// 验证结果
	h1 := Newtruncatedsha25648()
	h1.Write(x)
	resX := h1.Sum(nil)

	h2 := Newtruncatedsha25648()
	h2.Write(y)
	resY := h2.Sum(nil)

	fmt.Printf("找到碰撞！\n")
	fmt.Printf("x (hex): %s\n", hex.EncodeToString(x))
	fmt.Printf("y (hex): %s\n", hex.EncodeToString(y))
	fmt.Printf("H(x) (hex): %s\n", hex.EncodeToString(resX))
	fmt.Printf("H(y) (hex): %s\n", hex.EncodeToString(resY))

	if bytes.Equal(resX, resY) && !bytes.Equal(x, y) {
		fmt.Println("验证成功：输入不同且哈希值相等。")
	} else {
		fmt.Println("验证失败。")
	}
}

func TestSHA256_56(t *testing.T) {
	fmt.Println("正在尝试对 SHA-256 (截断至后56位) 进行碰撞攻击...")

	x, y := SmallSpaceBirthdayAttack(Newtruncatedsha25656)

	// 验证结果
	h1 := Newtruncatedsha25656()
	h1.Write(x)
	resX := h1.Sum(nil)

	h2 := Newtruncatedsha25656()
	h2.Write(y)
	resY := h2.Sum(nil)

	fmt.Printf("找到碰撞！\n")
	fmt.Printf("x (hex): %s\n", hex.EncodeToString(x))
	fmt.Printf("y (hex): %s\n", hex.EncodeToString(y))
	fmt.Printf("H(x) (hex): %s\n", hex.EncodeToString(resX))
	fmt.Printf("H(y) (hex): %s\n", hex.EncodeToString(resY))

	if bytes.Equal(resX, resY) && !bytes.Equal(x, y) {
		fmt.Println("验证成功：输入不同且哈希值相等。")
	} else {
		fmt.Println("验证失败。")
	}
}

func TestSHA256_64(t *testing.T) {
	fmt.Println("正在尝试对 SHA-256 (截断至后64位) 进行碰撞攻击...")

	x, y := SmallSpaceBirthdayAttack(Newtruncatedsha25664)

	// 验证结果
	h1 := Newtruncatedsha25664()
	h1.Write(x)
	resX := h1.Sum(nil)

	h2 := Newtruncatedsha25664()
	h2.Write(y)
	resY := h2.Sum(nil)

	fmt.Printf("找到碰撞！\n")
	fmt.Printf("x (hex): %s\n", hex.EncodeToString(x))
	fmt.Printf("y (hex): %s\n", hex.EncodeToString(y))
	fmt.Printf("H(x) (hex): %s\n", hex.EncodeToString(resX))
	fmt.Printf("H(y) (hex): %s\n", hex.EncodeToString(resY))

	if bytes.Equal(resX, resY) && !bytes.Equal(x, y) {
		fmt.Println("验证成功：输入不同且哈希值相等。")
	} else {
		fmt.Println("验证失败。")
	}
}
