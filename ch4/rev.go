// 就地反转一个整型 slice 中的元素
package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	reverse(a)
	fmt.Println(a)

	b := []int{0, 1, 2, 3, 4, 5}
	// 将 slice 中的元素左移 2 位
	reverse(b[:2])
	fmt.Println(b) // [1 0 2 3 4 5]	// 反转前 n 个元素
	reverse(b[2:])
	fmt.Println(b) // [1 0 5 4 3 2]	// 反转剩余元素
	reverse(b)
	fmt.Println(b) // [2 3 4 5 0 1]	// 全部反转
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
