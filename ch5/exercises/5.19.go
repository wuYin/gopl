package main

import "fmt"

func main() {
	fmt.Println(gg()) // 1
}

func gg() (x int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[panic]:", err)
			x = 1
		}
	}()
	panic("good game")
}
