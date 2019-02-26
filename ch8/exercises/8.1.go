package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}

	loc2hosts := make(map[string]string)
	for _, arg := range args {
		pair := strings.Split(arg, "=")
		if len(pair) < 2 {
			continue
		}
		loc, host := pair[0], pair[1]
		if !strings.Contains(host, ":") {
			continue
		}
		loc2hosts[loc] = host
	}

	if len(loc2hosts) == 0 {
		return
	}

	conns := make(map[string]net.Conn)
	for loc, host := range loc2hosts {
		conn, err := net.Dial("tcp", host)
		if err != nil {
			fmt.Println(err)
			continue
		}
		conns[loc] = conn
	}
	if len(conns) == 0 {
		return
	}

	for {
		for loc, conn := range conns {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err == io.EOF || err != nil { // 连接关闭
				conn.Close()
				return
			}
			if n > 0 {
				os.Stdout.WriteString(fmt.Sprintf("%s: %s\n", loc, buf))
			}
		}
	}
}
