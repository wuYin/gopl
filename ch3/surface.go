// surface 函数根据一个三维曲面函数计算并生成 svg 图
package main

import (
	"math"
	"fmt"
)

const (
	width, height = 600, 320            // 画布像素大小
	cells         = 100                 // 格子数
	xyRange       = 30.0                // x,y 坐标轴范围
	xyScale       = width / 2 / xyRange // x,y 轴单位长度
	zScale        = height * 0.4        // z 轴单位长度
	angle         = math.Pi / 6         // x,y 轴角度：30°
)

var sin30, cos30 = math.Sin(30), math.Cos(30)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			// 参数列表较长时，参数列表换行是好习惯
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyScale
	sy := height/2 + (x+y)*sin30*xyScale - z*zScale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
