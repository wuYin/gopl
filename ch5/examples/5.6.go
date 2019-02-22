package main

import (
	"fmt"
	"os"
)

func main() {
	f := squares()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9 // x 的状态是跟随着 f 走的，函数变量 f 内部含有变量 x，这也是函数不可比较的原因

	// 捕获迭代变量的陷进
	localVarTrap()
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func localVarTrap() {
	var rmDirs []func() error
	tempDirs := []string{"temp_a", "temp_b", "temp_c"}
	for _, dir := range tempDirs {
		tmp := dir          // ok
		os.Mkdir(dir, 0755) // always ok
		rmDirs = append(rmDirs, func() error {
			return os.Remove(tmp) // ok
		})
	}

	// 使用临时目录
	// ...

	for _, rmdir := range rmDirs {
		if err := rmdir(); err != nil {
			fmt.Println(err) // remove temp_c: no such file or directory
		}
	}
}
