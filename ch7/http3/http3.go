package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoe": 50, "socks": 5}
	mux := http.NewServeMux() // 路由复用器
	// mux.Handle("/list", http.HandlerFunc(db.list))
	mux.HandleFunc("list", db.list) // HandlerFunc 类型转换
	// mux.Handle("/price", http.HandlerFunc(db.price))
	mux.HandleFunc("price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8080", mux)) // handler 拿给 mux 统一处理和分发 url
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
