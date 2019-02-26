package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		fmt.Println("done")
		done <- struct{}{} // 子 goroutine 和主 goroutine 通过 done channel 来通信，只不过主 goroutine 对接收的值不予处理而已
	}()

	mustCopy(conn, os.Stdin)
	conn.Close() // 手动关闭连接
	<-done       // 无缓冲通道阻塞等待
}
