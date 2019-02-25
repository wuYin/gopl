package main

import (
	"fmt"
	"syscall"
)

func main() {
	err := syscall.Errno(2)
	fmt.Println(err.Error())
}
