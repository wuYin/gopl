// 增强的 comma 函数，能正确处理浮点数及正负号
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("-10086.1222"))
	fmt.Println(comma("+123456789.1234"))
}

func comma(s string) string {
	var buf bytes.Buffer
	start := 0
	// 先检查正负号
	if s[0] == '+' || s[0] == '-' {
		start = 1
		buf.WriteByte(s[0])
	}
	// 找出 . 的位置
	end := strings.LastIndex(s, ".")
	if end == -1 {
		end = len(s)
	}

	positiveNum := s[start:end] // 小数去掉符号后的整数部分

	// 找出多余的位数
	pre := len(positiveNum) % 3
	if pre > 0 {
		buf.Write([]byte(positiveNum[:pre]))
		buf.WriteString(",")
	}
	// 处理之后的位数
	for i, r := range positiveNum[pre:] {
		if i > 0 && i%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(r)
	}

	buf.WriteString(s[end:])
	return buf.String()
}
