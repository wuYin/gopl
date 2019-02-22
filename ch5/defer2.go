package main

import (
	"fmt"
	"runtime"
)

func main() {
	defer printStack() // defer 语句在程序栈清理之前得以调用
	g(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Println(string(buf[:n]))
}


func g(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // 0/0 发生 panic
	defer fmt.Println("defer x=", x)
	g(x - 1)
}
