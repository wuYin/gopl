package main

import (
	"bufio"
	"fmt"
	"strings"
)

var buf = []byte(`
see you
again
`)

func main() {
	var lc LineCounter
	lc.Write(buf)
	fmt.Println(lc) // 4

	var wc WordCounter
	wc.Write(buf)   // 其实 buf 此时不该再使用
	fmt.Println(wc) // 3
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (n int, err error) {
	*c = 1 // 多 1 行
	for _, b := range p {
		if b == '\n' {
			*c++
		}
	}
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords) // 为 scanner 设置 split 的方式
	*c = 0
	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}
