package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 验证 m:n / 协程:OS 线程调度
	// runtime.GOMAXPROCS(1) // 靠调度器来切换两个 goroutine
	runtime.GOMAXPROCS(2) // Go 调度器可支配 2 个线程调度切换
	for {
		go fmt.Print("1")
		fmt.Print("0")
	}

}
