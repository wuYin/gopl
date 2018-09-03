// 去除字符串 slice 中相邻重复元素
package main

import "fmt"

func main() {
	s := []string{"am", "is", "is", "are", "is"}
	fmt.Println(noneRepeat(s))
}

func noneRepeat(strings []string) []string {
	i := 0
	for _, s := range strings {
		// 下一字符串重复则跳过
		if strings[i] == s {
			continue
		}
		// 不重复，则复制
		i++
		strings[i] = s
	}
	return strings[:i+1]
}
