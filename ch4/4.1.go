// 统计 2 个 sha256 散列中不同的位数
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println(diff([]byte{1, 2, 3}, []byte{4, 5, 6})) // 2+3+2 = 7
	s1 := sha256.Sum256([]byte("x"))
	s2 := sha256.Sum256([]byte("X"))
	fmt.Println(diff(s1[:], s2[:])) // 125
}

func diff(b1, b2 []byte) (count int) {
	l1, l2 := len(b1), len(b2)
	maxL := l1
	if l2 > maxL {
		maxL = l2
	}
	// 按字节遍历并累计相异位数
	for i := 0; i < maxL; i++ {
		switch {
		case i > l1:
			count += popCount(b2[i])
		case i > l2:
			count += popCount(b1[i])
		default:
			count += popCount(b1[i] ^ b2[i])
		}
	}
	return
}

// 累积计算二进制数 b 中 1 的个数
func popCount(b byte) (count int) {
	for b != 0 {
		b = b & (b - 1) // 与自己减 1 相与，消除最低位的 1
		count++
	}
	return
}
