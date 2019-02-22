package main

import "fmt"

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3, 4))

	var f func(...int)
	var g func([]int)
	fmt.Printf("%T\n", f) // func(...int)
	fmt.Printf("%T\n", g) // func([]int) // 类型并不相同
}

// 变长函数
func sum(nums ...int) int {
	var res int
	for _, num := range nums {
		res += num
	}
	return res
}
