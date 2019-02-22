package main

import (
	"fmt"
	"math"
)

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance   // 方法变量，在调用时候只用传递实参即可，不必指定 receiver
	fmt.Println(distanceFromP(q)) // 5

	scaleP := p.ScaleBy // 方法变量的值是类型的方法
	scaleP(2)           // ok
	fmt.Println(p)      // {2 4}

	distance := Point.Distance
	fmt.Println(distance(p, q)) // 2.8

	scale := (*Point).ScaleBy
	scale(&p, 2)   // 第一个参数跟随 receiver
	fmt.Println(p) // {4 8}

	path := Path{
		{1, 1},
		{2, 2},
	}
	path.Move(Point{10, 10}, false) // 很酷
	fmt.Println(path)               // [{-9 -9} {-8 -8}]
}

type Path []Point

func (path Path) Move(offset Point, add bool) {
	var operate func(p, q Point) Point // p 参数是 receiver // 方法变量 operate 配合方法表达式使用
	if add {
		operate = Point.Add
	} else {
		operate = Point.Sub
	}
	for i := range path {
		path[i] = operate(path[i], offset)
	}
}

func (p Point) Add(q Point) Point {
	return Point{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}

func (p Point) Sub(q Point) Point {
	return Point{
		X: p.X - q.X,
		Y: p.Y - q.Y,
	}
}

type Point struct {
	X, Y float64
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
