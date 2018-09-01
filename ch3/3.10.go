// 使用 bytes.Buffer 完成 comma
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%v\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i, c := range s {
		// 向后倒数的数目是 3 的整数倍，则加入 , 号
		if (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(c)
	}
	return buf.String()
}
