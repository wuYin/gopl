package main

import (
	"fmt"
	"time"
)

var balance int // 银行的所有存款

func Save(amount int) {
	balance += amount
}

func Balance() int {
	return balance
}

func main() {
	for i := 1; i <= 10; i++ {
		go func() {
			Save(10)
		}()
		fmt.Println(i, balance) // 很显然，balance 和 i 是对不上的
	}
	// // Alice
	// go func() {
	// 	Save(200)
	// 	fmt.Println("bank balance", Balance())
	// }()
	//
	// // Bob
	// go Save(100)

	time.Sleep(100 * time.Millisecond)
}

// 串行受限
// 在 goroutine 之间传递数据并处理后，变不再访问该变量
// 也是并发安全的
type Cake struct {
	state string
}

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake // 之后便不再访问该 cake 变量，cake 变量在多个 goroutine 之间虽然共享，但是读写是受限的
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for c := range cooked {
		c.state = "iced"
		iced <- c // icer 方法不再访问 c
	}
}


// 总结一下
// 避免数据竞态
// 1. 直接就不要修改，数据只读，肯定是并发安全的：比如 map 只进行一次初始化，后边只读数据
// 2. monitor goroutine：将共享数据限制在一个 goroutine 内部进行读写，其他 goroutine 想读写只能通过 channel 来进行数据通信
// 3. 串行受限：多个 goroutine 的确能共享变量，但是每个变量在每个函数内部只操作一次
// 4. 互斥机制：sync.Mutex / sync.RWMutex / sync.WaitGroup