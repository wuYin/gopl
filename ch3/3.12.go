// 判断两个字符串是否同文异构
// 两个字符串含有相同的字母但顺序不同
package main

import "fmt"

func main() {
	s1, s2 := "golang", "langgo"
	fmt.Println(hasSameUnit(s1, s2))
}

func hasSameUnit(s1, s2 string) bool {
	u1 := make(map[rune]int)
	for _, r := range s1 {
		u1[r]++
	}

	u2 := make(map[rune]int)
	for _, r := range s2 {
		u2[r]++
	}

	for r, n := range u1 {
		if u2[r] != n {
			return false
		}
	}

	for r, n := range u2 {
		if u1[r] != n {
			return false
		}
	}
	return true
}
