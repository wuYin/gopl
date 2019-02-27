package main

import (
	"fmt"
	"time"
)

func main() {
	var x, y int
	go func() {
		x = 1
		fmt.Println("y", y)
	}()

	go func() {
		y = 1
		fmt.Println("x", x)
	}()

	time.Sleep(10 * time.Millisecond)
}
