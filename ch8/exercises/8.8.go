package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listen on localhost:8080")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	idleCh := make(chan bool)
	go detect(conn, idleCh)
	for input.Scan() {
		idleCh <- true // ok // 只要是无缓冲的同步channel，都要提防可能卡住的死锁问题
		go echo(conn, input.Text(), 1*time.Second)
	}
}

func detect(conn net.Conn, idleCh chan bool) {
	maxIdle := 10
	idleSeconds := 0
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			idleSeconds++
			if idleSeconds > maxIdle { // 释放连接
				msg := conn.RemoteAddr().String() + " 10s has no request, kicked out"
				fmt.Println(msg)
				fmt.Fprint(conn, msg)
				conn.Close()
				close(idleCh)
				ticker.Stop()
				return
			}
		case <-idleCh:
			idleSeconds = 0 // 有数据来就重置计数器即可
		}
	}
}

func echo(c net.Conn, s string, d time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(s))
	time.Sleep(d)
	fmt.Fprintln(c, "\t", s)
	time.Sleep(d)
	fmt.Fprintln(c, "\t", strings.ToLower(s))
}
