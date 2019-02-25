package http1

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8080", db)) // database 实现了 http.Handler 接口，可以处理请求并响应
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("%.2f", d)
}

// 模拟数据库存储
type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for k, v := range db {
		fmt.Fprintf(w, "%s: $%s\n", k, v) // 将响应数据写入到 writer 数据流中
	}
}
