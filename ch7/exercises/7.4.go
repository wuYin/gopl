package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

func main() {
	doc, _ := html.Parse(NewReader(`<html><head></head><body><div>DIV Content</div></body></html>`))
	fmt.Println(doc.LastChild.FirstChild.Data)
}

// 实现 io.Reader 接口，只是将传入的 []byte 处理了即可
// 返回你成功处理了的字节数，及可能遇到的错误
// 至于 Read 内部的逻辑，接口才不会管，简单复杂都与 io.Reader 无关
type StringReader struct {
	s string
}

func (r *StringReader) Read(p []byte) (n int, err error) {
	n = len(p)
	err = io.EOF         // 手动设为读取结束
	copy(p, []byte(r.s)) // 读取数据
	return
}

// 自己实现 strings.NewReader(s string) Reader
// 这里的 Reader 只是将 string 读取到内部
func NewReader(s string) *StringReader {
	return &StringReader{s: s}
}
