// server1 是一个迷你回声服务器
package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe(":2333", nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
