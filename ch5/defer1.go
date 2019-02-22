package main

import "fmt"

func main() {
	f(3)
	// f(3)
	// f(2)
	// f(1)
	// defer x= 1
	// defer x= 2
	// defer x= 3
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // 0/0 发生 panic
	defer fmt.Println("defer x=", x)
	f(x - 1)
}
