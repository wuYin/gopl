package circle

type point struct {
	X, Y int
}

// 嵌入，将成员直接递归取出
type Circle struct {
	point
	Radius int
}
