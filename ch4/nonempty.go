// nonempty 演示了 slice 的就地修改算法
// 从给定的字符串列表中剔除空字符串并返回
package main

import "fmt"

func main() {
	s := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(s)) // ["one" "three"]
	fmt.Printf("%q\n", s)           // ["one" "three" "three"]
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s // 直接修改底层数组
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
