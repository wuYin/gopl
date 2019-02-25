package main

import "fmt"

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // 5

	c = 0
	fmt.Fprintf(&c, "I have 16 length") // 注意形参是接口类型时，实参是实现该接口的任意类型指针
	fmt.Println(c)                      // 16
}

type ByteCounter int

// 实现 io.Writer 接口
func (c *ByteCounter) Write(p []byte) (n int, err error) {
	*c += ByteCounter(len(p)) // 对数据不作处理
	return len(p), nil
}
