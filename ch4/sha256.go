// sha256 比较两个 [32]byte 数组的值是否一致
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	fmt.Printf("%x\n", []int{1, 2, 12}) // slice 按十六进制输出
}
