package main

import (
	"fmt"
)

func main() {
	fmt.Println(min(1, 3, -4, 9))
	fmt.Println(max(1, 3, -1, 9))
}

// 要求至少有一个参数：显式将参数定为 1 个
func min(base int, nums ...int) int {
	for _, num := range nums {
		if num < base {
			base = num
		}
	}

	return base
}

func max(base int, nums ...int) int {
	for _, num := range nums {
		if num > base {
			base = num
		}
	}

	return base
}
