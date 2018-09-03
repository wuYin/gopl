package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []byte("  name is  wuYin")
	res := subSpace(s)
	fmt.Printf("%q %d %d\n", res, len(res), cap(res))
}

func subSpace(s []byte) []byte {
	count := 0
	for i, b := range s {
		if unicode.IsSpace(rune(b)) {
			if i > 0 && unicode.IsSpace(rune(s[i-1])) {
				copy(s[i:], s[i+1:])
				count++
			}
		}
	}
	return s[:len(s)-count]
}
