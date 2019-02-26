package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("counting... press return to abort...")
	ticker := time.Tick(1 * time.Second) // 有点像在内部 sleep 定时醒来后发送数据出来的 tick
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		select {
		case <-ticker: // 只是为了继续走流程循环 // 当 abort 后 ticker 依旧不断地发送数据。发生资源泄漏 // 可使用 NewTicker().Stop() 手动处理
		case <-abort:
			fmt.Println("aborted")
			return
		}
	}

	fmt.Println("lunched")
}
