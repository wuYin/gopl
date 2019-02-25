package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoe": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list": // 既然是 path，/ 就是必不可少的，就是这么精确
		for k, v := range db {
			fmt.Fprintf(w, "%s: $%s\n", k, v)
		}
	case "/price":
		item := r.URL.Query().Get("item")
		if price, ok := db[item]; ok {
			fmt.Fprintf(w, "%s: $%s\n", item, price)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item %s not found", item)

		// 这里必须在向 w 写入数据之前就将 header 设置好，否则是不生效的（情理之中）
		// w.WriteHeader(http.StatusNotFound) // http: multiple response.WriteHeader calls
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "url not found")
	}
}
