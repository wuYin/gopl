package main

import (
	"log"
	"net"
	"os"
)

// 同时 run 2 个  gr util.go netcat1.go
// clock1.go 只能处理一个请求，另一个阻塞
func main() {
	// 连接到指定的服务器
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn) // 注意 io.Copy 不是一次性操作，会一直读取直到 io.EOF 错误
}
