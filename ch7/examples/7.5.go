package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {

	// 接口值是可以比较的
	// 动态类型和动态值都完全一致，则相等
	var x, y interface{}
	file := os.Stdout
	x, y = file, file
	fmt.Println(x == y)         // true
	fmt.Printf("%T %v\n", x, x) // *os.File &{0xc420084050}

	// 空接口值与 nil 不同
	// nil 值是值动态类型和动态值都是不定的，都是 nil
	// 下方在 f 中使用 out 变量时，out 有动态类型是 *bytes.Buffer，但值依旧是 nil，发生 panic
	debug := true
	// var buf *bytes.Buffer
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf) // panic: runtime error: invalid memory address or nil pointer dereference
	if debug {
		// ...
	}
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte(time.Now().Format("2006-01-02") + " debug done "))
	}
}
