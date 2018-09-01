// basename 移除路径部分和 . 后缀
// eg:	a => a, a.go => a, a/b/c.go => c, a/b.c.go	=> b.c
package main

import "fmt"

func main() {
	fmt.Println(basename("a"))
	fmt.Println(basename("a.go"))
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("a/b.c.go"))
}

func basename(s string) string {
	// 舍弃最后一个 / 之前的内容
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// 保留最后一 . 之前的内容
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
