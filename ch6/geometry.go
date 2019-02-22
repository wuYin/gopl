package main

import (
	"fmt"
	"math"
)

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // 5
	fmt.Println(p.Distance(q))  // 5

	path := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(path.Distance()) // 12

	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(r) // &{2 4}

	x := Point{1, 2}
	(&x).ScaleBy(2) // ok
	y := &x
	y.ScaleBy(2)      // ok
	fmt.Println(r, x) // {4 8} // x 始终是值

	// 不能在无法获取地址的变量上进行 Point -> *Point 的 receiver 转换。地址都取不到转啥转
	// Point{1, 2}.ScaleBy(2) // cannot take the address of composite literal

	fmt.Println(y.Distance(*r)) // ok // *Point->Point 解引用，直接取值调用，不会失败的
}

type Point struct {
	X, Y float64
}

// 普通的函数
func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y) // 平方和函数
}

// Point 类型的方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

// 更新变量
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type Path []Point

// receiver 不同，方法就不同
func (path Path) Distance() float64 {
	var sum float64
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
