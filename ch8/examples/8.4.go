package main

import "fmt"

func main() {

	// 缓冲通道是一个指定类型的元素队列
	ch := make(chan string, 3)
	ch <- "A"
	ch <- "B"
	fmt.Println(len(ch), cap(ch)) // 2 3 // 调试时有用，一般存取数据都很快，价值不高
	ch <- "C"
	// ch <- "D" // 缓冲通道满了之后再发送数据会将 goroutine 死锁阻塞
	fmt.Println(<-ch) // A
}
