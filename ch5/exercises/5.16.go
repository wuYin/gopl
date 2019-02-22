package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(join(", ", "see", "you", "again"))
}

// 剔除最后一个字符串的后缀即可
func join(sep string, strs ...string) string {
	var buf bytes.Buffer
	n := len(strs)
	for i, s := range strs {
		buf.WriteString(s)
		if i < n-1 {
			buf.WriteString(sep)
		}
	}
	return buf.String()
}
