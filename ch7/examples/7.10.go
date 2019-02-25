package main

import (
	"io"
	"os"
)

func main() {
	// 类型断言
	// x.(T)
	var w io.Writer
	w = os.Stdout
	// T 是具体类型，则提取值
	// f := w.(*os.File)      // f == *os.File
	// c := w.(*bytes.Buffer) // panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer

	// T 是接口类型则判断是否满足接口
	rw := w.(io.ReadWriter)
	_ = rw
}
