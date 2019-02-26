package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch: // 输出偶数，此时 i++ 会跳过奇数
			fmt.Println(x)
		case ch <- i: // 先执行
		}
	}
}
