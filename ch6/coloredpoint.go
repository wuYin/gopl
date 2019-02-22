package main

import (
	"fmt"
	"image/color"
	"math"
)

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // 1
	cp.Y = 2
	fmt.Println(cp.Y) // 2 // 直接越过 Point

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint{Point{1, 1}, red}
	q := ColoredPoint{Point{5, 4}, blue}

	//  cannot use q (type ColoredPoint) as type Point in argument to p.Point.Distance
	// 虽然 q 组合了 Point，但是作为 Point 的形参类型传入时候必须显式指定使用 Point 字段
	fmt.Println(p.Distance(q.Point)) // 5
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // 10
}

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
