package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{
		items: map[string]dollars{"shoe": 50, "socks": 5},
		lock:  sync.RWMutex{},
	}

	// // 参数检查
	// go func(db database) {
	// 	for {
	// 		db.lock.RLock()
	// 		for k, v := range db.items {
	// 			fmt.Printf("%s: %s\n", k, v)
	// 		}
	// 		db.lock.RUnlock()
	// 		time.Sleep(5 * time.Second)
	// 	}
	// }(db)

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/add", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("%.2f", d)
}

type database struct {
	items map[string]dollars
	lock  sync.RWMutex // 考虑并发处理
}

// 查
func (db database) list(w http.ResponseWriter, req *http.Request) {
	db.lock.RLock()
	defer db.lock.RUnlock()

	for k, v := range db.items {
		fmt.Fprintf(w, "%s: $%s\n", k, v)
	}
}

// 查
func (db database) price(w http.ResponseWriter, req *http.Request) {
	db.lock.RLock()
	defer db.lock.RUnlock()

	item := req.URL.Query().Get("item")
	if price, ok := db.items[item]; ok {
		fmt.Fprintf(w, "%s: $%s\n", item, price)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "%s not found\n", item)
}

// curl http://localhost:8080/update?item=mouse&price=1000
func (db database) update(w http.ResponseWriter, req *http.Request) {
	db.lock.Lock()
	defer db.lock.Unlock()

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	newPrice, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "price arg invalid: %s\n", price)
		return
	}

	if _, ok := db.items[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s not found\n", item)
		return
	}

	db.items[item] = dollars(newPrice)
}

// curl http://localhost:8080/delete?item=mouse&price=1000
func (db database) delete(w http.ResponseWriter, req *http.Request) {
	db.lock.Lock()
	defer db.lock.Unlock()

	item := req.URL.Query().Get("item")
	if _, ok := db.items[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s not found\n", item)
		return

	}
	delete(db.items, item)
}

// curl http://localhost:8080/add?item=mouse&price=100
func (db database) add(w http.ResponseWriter, req *http.Request) {
	db.lock.Lock()
	defer db.lock.Unlock()

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	newPrice, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "price arg invalid: %s\n", price)
		return
	}

	db.items[item] = dollars(newPrice)
}
