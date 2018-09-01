package main

// 更为简洁的容量声明
const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

// 也不错
const (
	_  = 1 << (iota * 10) // 2 ^ 0
	kb                    // 2 ^ 10
	mb                    // 2 ^ 20
	gb
)

func main() {

}
