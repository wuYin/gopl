package main

import (
	"fmt"
)

func main() {
	x := false
	fmt.Println(assert(x))
}

func assert(x interface{}) string {
	switch x := x.(type) { // 这里左侧的 x 是仅在 switch 作用域范围内的局部变量
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return x
	default:
		panic(fmt.Sprintf("unknown type: %T: %v", x, x))
	}
}
