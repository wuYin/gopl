// 初始化嵌入匿名成员的结构体
package main

import "fmt"

func main() {
	w1 := Wheel{Circle{Point{8, 8}, 5}, 20}
	w2 := Wheel{
		Circle: Circle{
			Point: Point{
				X: 8,
				Y: 8,
			},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w1)
	fmt.Println(w1 == w2) // true
}

type Circle struct {
	Point // 嵌入的匿名成员
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

type Point struct {
	X int
	Y int
}
