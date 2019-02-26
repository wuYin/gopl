package main

import (
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 两句话实现数据的双向流动，io.Copy 实在是太方便了
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}
