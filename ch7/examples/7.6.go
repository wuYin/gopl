// 使用 sort.Interface 来对任意序列，任意排序方式进行原地排序
package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"k8s", "ali", "elem", "docker"}
	sort.Sort(StringSlice(names))
	fmt.Println(names) // ok
}

type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}
