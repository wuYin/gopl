package main

func main() {
}

// 避免数据竞态：避免多个 goroutine 访问同一个变量
// 衍生出来监控 goroutine，只能由它来修改公用变量。其他的 goroutine 通过 channel 来与它通信，从而共享变量的内存
// 原则：不要通过共享内存来通信，要通过通信来共享内存。说的就是 monitor goroutine
var (
	saveCh    = make(chan int) // 发送存款项
	balanceCh = make(chan int) // 查询存款
)

func Save(n int) {
	saveCh <- n
}

func Balance() int {
	return <-balanceCh
}

func run() {
	var balances int // 控制只有 main goroutine 才能访问
	select {
	case n := <-saveCh:
		balances += n
	case balanceCh <- balances: // 阻塞发送
	}
}

func init() {
	go run() // 启动 monitor goroutine
}
