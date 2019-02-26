package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rfib(%d)=%d\n", n, fibN)
}

func spinner(d time.Duration) {
	for {
		for _, r := range `-\|/` { // 很酷的终端等待显示
			fmt.Printf("\r%c", r)
			time.Sleep(d)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
