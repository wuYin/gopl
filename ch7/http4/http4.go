package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoe": 50, "socks": 5}
	http.HandleFunc("/list", db.list) // 将 db.list 注册到全局的 DefaultServeMux 上进行 URL 的分发处理
	http.HandleFunc("/price", db.price)
	// log.Fatal(http.ListenAndServe("localhost:8080", nil))                  // DefaultServeMux 是整个项目的主处理程序，无需显式传递
	log.Fatal(http.ListenAndServe("localhost:8080", http.DefaultServeMux)) // 你要显式地写也不拦着你
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for k, v := range db {
		fmt.Fprintf(w, "%s: $%s\n", k, v)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s: $%s\n", item, price)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "%s not found\n", item)
}
