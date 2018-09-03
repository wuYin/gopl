package export

import (
	"fmt"
	"gopl/ch4/circle"
)

func Demo() {
	// 嵌套的结构体必须指定字段名来初始化
	c1 := circle.Circle{}
	fmt.Printf("%#v\n", c1) // main.Circle{Point:main.Point{X:0, Y:0}, Radius:1}
	fmt.Println(c1.X)       // ok
	// fmt.Println(c1.point.X) // not ok
}
