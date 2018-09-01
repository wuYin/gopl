package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {

	// byte 是 uint8 的别名，强调值是原始字节数据
	var b1 uint8
	var b2 byte
	fmt.Println(b1 == b2) // true

	// rune 是 int32 的别名，强调值是 Unicode 码点值
	var r1 rune
	var i1 int32
	fmt.Println(i1 == r1) // true

	// int 与 int32 在绝大部分平台上是一致的，但编译器不认为是同一类型
	// var i2 int
	// fmt.Println(i1 == i2)	// invalid operation: i1 == i2 (mismatched types int32 and int)

	fmt.Println(-5 % -3) // -2	// php -r 'echo -5 % -3;'	// 负数给负数取余，结果还是负数
	fmt.Println(7 / 3)   // 2	// 整数除法舍弃小数位
	fmt.Println(7 / 3.0) // 2.333

	var i2 uint8 = 127
	fmt.Println(strconv.FormatInt(127*127, 2)) // 11111100000001
	fmt.Println(i2 * i2)                       // 1 // 发生溢出后，高于 8 位的直接截取掉，乘积剩下 0000 0001
	fmt.Println(i2*i2 + 2)                     // 3 // 11111100000011

	fmt.Println(100 >> 3) // 12	// 左移运算符结果向下取整 12.5 -> 12

	// 位运算
	var x uint8 = 1<<1 | 1<<5  // 00100010
	var y uint8 = 1<<1 | 1<<2  // 00000110
	fmt.Printf("%08b\n", x^y)  // 00100100	// 异或，相异为真
	fmt.Printf("%08b\n", x&^y) // 00100000	// 按 y 位清零

	names := []string{"pike", "ken"}
	fmt.Printf("%T\n", len(names)) // len 返回值是 int 而非看起来更合理的 uint（向下减会溢出）
	// 出于可能溢出的考虑，无符号类型一般很少用来表示非负数（有点奇怪，但会更安全）

	f1, f2 := 2.8, -2.8
	fmt.Println(int(f1), int(f2)) // 2 -2	// 负数向上取整，浮点转整型去掉小数部分

	o := 0666
	fmt.Printf("%d %[1]o %#[1]x\n", o) // [1] 重复使用第一个操作数		# 输出前缀
	// 整数进制之间的转换直接使用标识符，fmt.Sprinf("%x %o") 等

	// float32 有效数字6位，会迅速累积误差
	var f3 float32 = 1 << 24
	fmt.Println(1<<24, f3 == f3+1) // 16777216, true

	var f4 float64
	fmt.Println(f4 / f4) // NaN	// 数学上无意义的值

	var s string
	if s != "" && s[0] != 'a' { // 短路求值，s[0] 取值安全
		fmt.Println(s)
	}

	s1 := "中"
	fmt.Println(len(s1)) // 3 字节

	// 字面量
	// \r 为回车符，去到本行开头
	// \n 为换行符，读到下一行
	// \n\r 为键盘 Enter 信号

	fmt.Println(utf8.RuneCountInString(s1)) // 1
	fmt.Println(string(65))                 // A

	const (
		a int = iota
		b
		c = 0
		d = iota // 3	// 注意不是 0
		e        // 4
	)
	fmt.Println(a, b, c, d, e)

	var f5 float64 = 100
	fmt.Printf("%T %[1]v\n", (f5-32)*5/9) // float64  37.77
	// 书中翻译描述有误
	fmt.Printf("%T %[1]v\n", 5/9*(f5-32)) // float64  0
}
