// append 为 []int 数组 slice 实现的 append
package main

import "fmt"

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d %v\n", i, cap(y), y)
		x = y
	}
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zLen := len(x) + len(y)
	if zLen < cap(x) {
		// slice x 仍然有增长空间
		z = x[:zLen] // [:]分割: z 与 x 有相同的底层数组
	} else {
		// slice x 无增长空间
		// 为达到增长线性，成倍增长（新 slice 的长度必须是整数）
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		z = make([]int, zLen, zCap)
		copy(z, x) // copy: z 与 x 的底层数组不同
	}
	copy(z[len(x):], y)
	return z
}
