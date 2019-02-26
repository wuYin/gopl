package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	go broadcast()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	cliMsgCh := make(chan string) // 当前连接响应的 channel
	go func() {
		for msg := range cliMsgCh { // 接收广播消息
			fmt.Fprint(conn, msg)
		}
	}()

	name := conn.RemoteAddr().String()
	cliMsgCh <- "you are " + name + "\n" // 返回给客户端
	messages <- name + " has joined\n"   //
	entrying <- cliMsgCh

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- name + " : " + input.Text() + "\n"
	}

	leaving <- cliMsgCh
	messages <- name + " has left\n"
}

type clientMsgCh chan<- string // 对客户端发送消息的 channel

var (
	entrying = make(chan clientMsgCh) // 这里的设计十分巧妙 // channel 中的值可以是任意值，甚至是引用 channel
	leaving  = make(chan clientMsgCh)
	messages = make(chan string)
)

func broadcast() {
	clients := make(map[clientMsgCh]bool) // channel 是引用，是可比较的
	for {
		select {
		case newClient := <-entrying:
			clients[newClient] = true
		case leftClient := <-leaving:
			delete(clients, leftClient)
			close(leftClient)
		case msg := <-messages:
			// 广播发送消息
			for cli := range clients {
				cli <- msg
			}
		}
	}
}
