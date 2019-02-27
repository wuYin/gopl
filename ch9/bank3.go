package main

import "sync"

func main() {
	// 使用 sync.Mutex 来保护共享变量的访问
	Withdraw2(2) //  all goroutines are asleep - deadlock!
}

var (
	mu      sync.Mutex
	balance int
)

func Save(n int) {
	mu.Lock()
	defer mu.Unlock()
	balance += n
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func Withdraw(n int) bool {
	// save 会进行锁的抢占，修改 balance 后再释放
	// 有时机会导致 balance 降到零下
	Save(-n)
	if Balance() < 0 {
		Save(n)
		return false
	}
	return true
}

func Withdraw2(n int) bool {
	// save 会进行锁的抢占，修改 balance 后再释放
	// 有时机会导致 balance 降到零下
	mu.Lock() // 互斥锁是不能重入的
	defer mu.Unlock()
	Save(-n)
	if Balance() < 0 {
		Save(n)
		return false
	}
	return true
}
