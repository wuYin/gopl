package main

import "fmt"

func main() {

	var 无0 = 0 // 变量声明开头是 Unicode 字符即可
	fmt.Println(无0)

	// true, make := 0, 1
	// fmt.Println(true, make) 	// 内置常量和函数不是预留的
	// m := make(int)			// 重用的风险： cannot call non-function make (type int)

	// i := 0
	// i++ 是语句：语句是一段可执行的代码，不一定有值，不能放在 = 右侧
	// j := i++		// syntax error: unexpected ++ at end of statement
	// i+1 是表达式：表达式是可求值的代码，一定有值，可放在 = 右侧
	// _ = i + 1	// ok

	// var s string
	// 在函数内部使用 var 的场景
	// 1. 局部变量的类型与默认类型不一致，必须使用显式类型声明：var id int64
	// 2. 表示变量的初始值不重要，后边再对变量赋值

	// 结构体、数组、slice（聚合类型）中的元素是可寻址的
	names := []string{"ken", "robe"}
	fmt.Printf("%v\n", &names[1])

	type User struct {
		Name string
		Age  int
	}
	ken := User{"ken", 18}
	fmt.Printf("%v\n", &ken.Age)

	// map 是不可寻址的，元素可能会随着 map 增长而动态改变地址
	// dic := make(map[string]string)
	// dic["one"] = "一"
	// fmt.Printf("%v\n", &dic["one"]) // cannot take the address of dic["one"]

	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil) // true false false

	var x1 struct{}
	var y1 struct{}
	fmt.Printf("%t\n", &x1 == &y1) // false // 书中描述有误

	fmt.Println(gcd(64, 88))
	fmt.Println(fib(10))

	type pint1 *int
	type pint2 *int
	p1 := pint1(&x)
	p2 := pint2(&y)
	fmt.Println(*pint1(p2)) // 两个类型，有相同的底层类型、指向相同的底层类型，则可以相互进行类型转换
	fmt.Println(*pint2(p1))

	type newint1 int
	type newint2 int
	n1 := newint1(1)
	// n2 := newint2(1)
	fmt.Println(n1 == 1) // 自定义类型只能向下比较
	// fmt.Println(n1 == n2) // invalid operation: n1 == n2 (mismatched types newint1 and newint2)

	// 编译器从最内层开始寻找变量的声明
	// 外层的同名声明会被覆盖
	s := "hello!" // string
	for i := 0; i < len(s); i++ {
		s := s[i] // rune
		if s != '!' {
			s := s + 'A' - 'a' // rune
			fmt.Printf("%c", s)
		}
	}

	// switch 同 if，可带初始化语句
	switch x := 0; x > -1 {
	case true:
		sn := 1
		fmt.Println(sn)
	default:
		// fmt.Println(sn)	// 各 case 间相互独立
	}

}

// 多个 init 函数按照其声明顺序执行
func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

// 计算两个数的最大公约数
// 一直给对方取余数，交换值再取余，直到对方值为 0
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// 计算斐波那契数列的第 n 个数
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y // 向后顺延
	}
	return x
}

var global *int

func f() {
	var x int
	x = 1
	global = &x // x 从 f() 中逃逸，在堆空间上分配
}
func g() {
	var y *int
	*y = 1 // y 在栈空间上分配
}
