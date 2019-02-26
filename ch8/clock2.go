package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var port = flag.String("port", "8080", "-port 8080")

func main() {
	flag.Parse()
	host := fmt.Sprintf("localhost:%s", *port)
	l, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("clock2 listened:", *port)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 换成并发处理请求
		go handleConn(conn)
	}
}
