// intsToString 与 fmt.Sprint(values) 类似，但插入了逗号
package main

import (
	"bytes"
	"fmt"
)

func main() {
	arr := []int{1, 2, 2, 3}
	s := fmt.Sprint(arr)
	fmt.Println(s)
	fmt.Println(intsToString(arr))
}

func intsToString(vals []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, val := range vals {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", val)
	}
	buf.WriteByte(']')
	return buf.String()
}
