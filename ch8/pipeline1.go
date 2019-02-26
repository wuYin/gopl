package main

import "fmt"

func main() {
	counter := make(chan int) // 都是同步阻塞的无缓冲 channel
	square := make(chan int)

	go func() {
		for x := 0; ; x++ {
			counter <- x
		}
	}()

	go func() {
		for {
			x := <-counter
			square <- x * x
		}
	}()

	for {
		fmt.Println(<-square)
	}
}
