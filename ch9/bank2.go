package main

func main() {
	// sync.Mutex 可理解为是容量为 1 的缓冲通道实现的变量访问控制机制
}

var (
	sema    = make(chan struct{}, 1) // 容量只有 1 的缓冲 channel，保证同一时刻只能有一个 goroutine 访问 balance
	balance int
)

func Save(n int) {
	sema <- struct{}{} // 获得锁
	balance += n
	<-sema // 释放锁
}

func Balance() int {
	sema <- struct{}{}
	v := balance
	<-sema
	return v
}
