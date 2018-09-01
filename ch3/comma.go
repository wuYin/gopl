// comma 向表示十进制非负整数的字符串中插入逗号
package main

import (
	"os"
	"fmt"
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
	return comma(s[:n-3]) + "," + s[n-3:]
}
