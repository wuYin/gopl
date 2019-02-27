package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var urls = []string{
	"https://cn.bing.com",
	"https://www.baidu.com",
	"https://weibo.com",
	"https://bilibili.com",
	"https://cn.bing.com",
}

func main() {
	m := NewMemo(httpGetBody)
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			v, err := m.Get(url)
			if err != nil {
				fmt.Println(url, err)
			}
			fmt.Printf("[URL]: %s\t[cost]: %s\t[size]: %d\n", url, time.Since(start), len(v.([]byte)))
		}(url)
	}
	wg.Wait()
}

type Item struct {
	v   interface{}
	err error
}

// 封装了广播机制的 item
type Entry struct {
	item  Item
	ready chan struct{} // 广播 channel
}

type request struct {
	k        string
	response chan<- Item
}

type Memo struct {
	requests chan request
	done     chan struct{}
}

func NewMemo(update UpdateFunc) *Memo {
	m := &Memo{requests: make(chan request)}
	go m.Serve(update)
	return m
}

func (m *Memo) Get(k string) (interface{}, error) {
	resp := make(chan Item)
	m.requests <- request{k: k, response: resp}
	item := <-resp
	return item.v, item.err
}

func (m *Memo) Serve(update UpdateFunc) {
	cache := make(map[string]*Entry)
	for req := range m.requests {
		entry := cache[req.k]
		if entry == nil {
			// miss
			entry = &Entry{ready: make(chan struct{})}
			cache[req.k ] = entry        // step1
			go entry.exec(update, req.k) // step2
		}
		go entry.wait(req.response)
	}
}

func (m *Memo) Close() {
	close(m.requests)
}

func (e *Entry) wait(response chan<- Item) {
	<-e.ready
	response <- e.item
}

func (e *Entry) exec(update UpdateFunc, k string) {
	e.item.v, e.item.err = update(k)
	close(e.ready) // 发送广播
}

type UpdateFunc func(k string) (interface{}, error)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
