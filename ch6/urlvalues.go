package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": []string{"zh-cn"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // zh-cn
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // 1
	fmt.Println(m["item"])     // [1 2]

	m = nil                    // nil 也是合法的 receiver
	fmt.Println(m.Get("item")) // 但无法通过编译（旧版），因为 nil 的类型是不定的
	// m.Add("item", "3") // panic: assignment to entry in nil map
}
