package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// 任何类型的值都可赋值给空接口
	var any interface{}
	fmt.Printf("%T\n", any) // nil
	any = 1
	fmt.Printf("%T\n", any) // int
	any = "string"
	fmt.Printf("%T\n", any) // string

	// 类型必需实现接口的断言
	var _ io.Writer = (*bytes.Buffer)(nil) // 说明 bytes.Buffer 类型实现了 io.Writer，即使值是 nil
}
