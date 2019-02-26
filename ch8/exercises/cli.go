package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		_, err := conn.Write([]byte(line))
		if err != nil {
			log.Fatal(err)
		}
	}
}
