package main

import (
	"io"
	"log"
	"net"
	"time"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// 建立与客户端的连接后，每隔 1s 给客户端发送当前时间
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		if _, err := io.WriteString(c, time.Now().Format("15:04:05\n")); err != nil {
			return // 发送响应发生错误则断开连接
		}
		time.Sleep(1 * time.Second)
	}
}
