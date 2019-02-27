package main

import (
	"fmt"
	"sync"
	"time"
)

var urls = []string{
	"https://cn.bing.com",
	"https://www.baidu.com",
	"https://cn.bing.com",
}

func main() {
	m := NewMemo(httpGetBody)
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		// 此时可能不会再命中缓存
		go func(url string) {
			start := time.Now()
			v, err := m.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%s\t%s\t%d\n", url, time.Since(start), len(v.([]byte)))
			wg.Done()
		}(url)
	}

	wg.Wait()
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]Item
}

type Item struct {
	v   interface{}
	err error
}

// 缓存值获取函数
type Func func(k string) (interface{}, error)

func NewMemo(f Func) *Memo {
	return &Memo{
		f:     f,
		mu:    sync.Mutex{},
		cache: make(map[string]Item),
	}
}

func (m *Memo) Get(k string) (interface{}, error) {
	m.mu.Lock()
	item, ok := m.cache[k] // 查询锁
	m.mu.Unlock()
	if ok {
		return item.v, nil
	}

	item.v, item.err = m.f(k) // 多个 goroutine 依旧可能同时执行慢操作 f 最终导致同样的 URL 处理多次

	m.mu.Lock()
	m.cache[k] = item // 数据锁
	m.mu.Unlock()
	return item.v, item.err
}
