package main

import (
	"fmt"
	"log"
	"net"
)

//  curl http://localhost:8080
func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 这里的设计是一次只能处理一个连接请求
		handleConn(c)
	}
}
