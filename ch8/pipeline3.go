package main

import "fmt"

func main() {
	n := make(chan int)
	s := make(chan int)

	// 数据从 n->s 关闭 n
	go counter(n)
	// 数据从 n 流出，计算后发送给 s，关闭 s
	go square(s, n)
	// 数据从 s 流出，打印
	printer(s)
}

// out 发送通道
func counter(out chan<- int) {
	for i := 0; i <= 100; i++ {
		out <- i
	}
	close(out)
}

// in 接收通道
// out 发送通道
func square(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
	// close(in) // 发送通道只能在发送到关闭
}

func printer(in <-chan int) {
	for res := range in {
		fmt.Println(res)
	}
}
