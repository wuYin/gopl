package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := []byte("I said 北京您早")
	fmt.Printf("%q\n", reverseUTF8(s))
}

// 原地反转 UTF8 字符 slice
// 按照字符算而非字节
func reverseUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		rev(b[i : i+size]) // 局部字符反转
		i += size
	}
	rev(b)
	return b
}

func rev(b []byte) {
	l := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[l-1-i] = b[l-1-i], b[i]
	}
}
