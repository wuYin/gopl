// basename 移除路径部分和 . 后缀
// eg:	a => a, a.go => a, a/b/c.go => c, a/b.c.go	=> b.c
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("a"))
	fmt.Println(basename("a.go"))
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("a/b.c.go"))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // "/" 不存在与 s 中会返回 -1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
