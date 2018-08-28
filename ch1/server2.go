// server2 是一个迷你回声和计数服务器
package main

import (
	"sync"
	"net/http"
	"fmt"
	"log"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", echo)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":2333", nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	mu.Unlock()
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count: %d\n", count)
	mu.Unlock()
}
