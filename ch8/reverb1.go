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
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		txt := scanner.Text()
		echo(conn, txt, 1*time.Second)
	}
}

func echo(c net.Conn, s string, d time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(s))
	time.Sleep(d)
	fmt.Fprintln(c, "\t", s)
	time.Sleep(d)
	fmt.Fprintln(c, "\t", strings.ToLower(s))
}
