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

	go handleMessage()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleMessage() {
	for {
		select {
		case newClient := <-join:
			allClients[newClient.name] = newClient
			go func() {
				message <- "Remain Clients: \n"
				for _, cli := range allClients {
					message <- cli.name + "\n"
				}
			}()
		case leaveClient := <-leave:
			delete(allClients, leaveClient.name)
			close(leaveClient.msgChan)
		case msg := <-message:
			for _, cliCh := range allClients {
				cliCh.msgChan <- msg
			}
		}
	}
}

var allClients = make(map[string]*Client)

type Client struct {
	msgChan chan<- string // 只能向客户端发送消息的 channel
	name    string        // socket
}

var (
	join    = make(chan *Client)
	leave   = make(chan *Client)
	message = make(chan string)
)

// 当有 client 到来时显示目前还存在的用户
func handleConn(conn net.Conn) {
	defer conn.Close()
	cliCh := make(chan string)
	cli := &Client{
		msgChan: cliCh,                      // 进行类型转换，其实还是要双向读写的
		name:    conn.RemoteAddr().String(), //
	}
	go func() {
		for msg := range cliCh {
			fmt.Fprint(conn, msg)
		}
	}()

	cli.msgChan <- "you are" + cli.name + "\n"
	join <- cli
	message <- "[join] " + cli.name + "\n"
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message <- "[message] " + cli.name + " : " + scanner.Text()
	}

	leave <- cli
	message <- "[leave]: " + cli.name + "\n"
}
