package main

import (
	"fmt"
	"strings"
)

func main() {
	f := square
	fmt.Println(f(3)) // 9

	f = negative
	fmt.Println(f(3))     // -3
	fmt.Printf("%T\n", f) // func(int) int

	// f = product // 编译错误：cannot use product (type func(int, int) int) as type func(int) int in assignment

	var f1 func(int int) // func 的零值是 nil
	// f1(3)             // panic: runtime error: invalid memory address or nil pointer dereference
	if f1 != nil {
		f1(3)
	}
	// if f1 != f  // invalid operation: f1 != f (mismatched types func(int) and func(int) int) // 函数之间不可比较

	fmt.Println(strings.Map(add, "abc")) // bcd
}

func add(r rune) rune {
	return r + 1
}

func square(x int) int {
	return x * x
}

func negative(x int) int {
	return -x
}

func product(x, y int) int {
	return x * y
}
