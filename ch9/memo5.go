package main

import (
	"fmt"
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
			start := time.Now()
			v, err := m.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%s\t%s\t%d\n", url, time.Since(start), len(v.([]byte))) // 这里看到的两次必应请求其实只有一次
			wg.Done()
		}(url)
	}

	wg.Wait()
}

type Memo struct {
	requests chan request
}

func NewMemo(f Update) *Memo {
	memo := &Memo{
		requests: make(chan request),
	}
	go memo.Serve(f)
	return memo
}

func (m *Memo) Get(k string) (interface{}, error) {
	resp := make(chan Item)
	m.requests <- request{k: k, response: resp,}
	item := <-resp // 立刻阻塞等地数据返回
	return item.v, item.err
}

func (m *Memo) Close() {
	close(m.requests)
}

func (m *Memo) Serve(f Update) {
	cache := make(map[string]*Entry) // 将共享变量限制在单一的 goroutine 中读写
	for req := range m.requests {
		e := cache[req.k]
		if e == nil {
			// 未命中
			e = &Entry{ready: make(chan struct{})}
			cache[req.k] = e
			// 执行缓存更新函数，但是不要阻塞 monitor goroutine 继续处理其他请求，开子协程去处理
			go e.exec(f, req.k)
		}
		// 命中或未命中都要将数据返回
		go e.wait(req.response)
	}
}

func (e *Entry) wait(response chan<- Item) {
	<-e.ready
	response <- e.item
}

func (e *Entry) exec(f Update, k string) {
	fmt.Println("[request]:", k) // good
	e.item.v, e.item.err = f(k)
	close(e.ready)
}

type Update func(k string) (interface{}, error) // 缓存更新函数

type request struct {
	k        string
	response chan<- Item
}

type Entry struct {
	item  Item          // 存储数据
	ready chan struct{} // 存放数据信号
}

type Item struct {
	v   interface{}
	err error
}
