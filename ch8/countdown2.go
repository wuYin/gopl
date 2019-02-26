package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 一直等待从 stdin 读取 1 个字符数据
		abort <- struct{}{}
	}()

	fmt.Println("counting down")

	select {
	case <-time.After(10 * time.Second):
	case <-abort: // 10s 内收到从 os.Stdin 发来的数据，取消发射
		fmt.Println("lunch abort")
		return
	}

	fmt.Println("lunched")
}
