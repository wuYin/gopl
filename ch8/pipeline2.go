package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i <= 100; i++ {
			naturals <- i
		}
		close(naturals) // 手动通知接收方 channel 关闭 // 这一步是必须的，不然下方 range naturals 的操作会一直阻塞
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}

		// range 循环可以迭代 channel 直到 channel 中没有值再可以接收
		// for {
		// 	x, ok := <-naturals
		// 	if !ok {
		// 		break
		// 	}
		// 	squares <- x * x
		// }

		close(squares)
	}()

	for res := range squares {
		fmt.Println(res)
	}
}
