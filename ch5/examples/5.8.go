package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	double(2)
	fmt.Println(triple(4)) // 12
}

// defer 的匿名函数可得到外层函数作用域内的变量，包括命名返回值  // 匿名函数牛逼
func double(x int) (result int) {
	defer func() {
		fmt.Printf("[result]: %d\n", result) // 4
	}()
	return x + x
}

// 既然能获取，那就能修改
func triple(x int) (result int) {
	defer func() {
		result += x // 2*x += x
	}()
	return x + x
}

// 小心 for 遍历导致 defer 无法被执行的问题
func traverse(fNames []string) {
	handleFile := func(f *os.File) {
		defer f.Close() // ok
		// ... use f
	}
	for _, fName := range fNames {
		f, err := os.Open(fName)
		if err != nil {
			log.Fatal(err)
		}
		// defer f.Close() // not ok // 无休止地打开文件，可能导致文件描述符资源耗尽，有资源泄漏的风险
		// ... use f

		// 将 defer 操作放到函数中保证资源可以正确关闭
		handleFile(f) // ok
	}
}
