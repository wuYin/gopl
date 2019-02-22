package main

import "fmt"

func main() {
	noReturn()
	fmt.Printf("%T\n", sum) // func(int, int) int
}

// _ 表示形参在函数内未使用
func sum(x, _ int) int {
	return x
}

// 函数明确不会执行完，如死循环或 panic
// 有返回值也可以不显式 return
func noReturn() int {
	// panic("")
	for {
		x := 0
		_ = x
	}
}
