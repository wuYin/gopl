package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listened :8080")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// 沙雕的非实时 FTP
func handleConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		cmds := strings.Fields(scanner.Text())
		if len(cmds) < 2 {
			fmt.Println("cmds: ", cmds)
			continue
		}
		name, args := cmds[0], cmds[1:]
		cmd := exec.Command(name, args...)

		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			return
		}
		if err := cmd.Wait(); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("exec succ: %s %v\n", name, args)
	}
}
