package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// 简单的按行并发 echo
func handleConn(conn net.Conn) {
	// defer conn.Close() // 要求某个连接上无数据传输时关闭该连接

	var wg sync.WaitGroup
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		wg.Add(1)
		go func(c net.Conn, s string, d time.Duration) {
			fmt.Fprintln(c, "\t", strings.ToUpper(s))
			time.Sleep(d)
			fmt.Fprintln(c, "\t", s)
			time.Sleep(d)
			fmt.Fprintln(c, "\t", strings.ToLower(s))
		}(conn, scanner.Text(), 1*time.Second)
	}

	go func() {
		wg.Wait()
		if conn, ok := conn.(*net.TCPConn); ok {
			conn.CloseWrite() // 关闭写半边
		}
	}()
}
