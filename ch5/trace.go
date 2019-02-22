package main

import (
	"log"
	"time"
)

func main() {
	slow()
}

func slow() {
	defer trace("slow ...")() // () 这个厉害了，相当于模拟了 slow 的构造函数 和 析构函数
	// ...
	time.Sleep(5 * time.Second) // 模拟耗时操作
}

func trace(msg string) func() {
	now := time.Now()
	log.Printf("[start]: %s", msg)
	return func() {
		log.Printf("[exit]: %s\t[cost]: %dms",
			msg,
			time.Since(now).Nanoseconds()/int64(time.Millisecond))
	}
}
