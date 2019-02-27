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
	f     Func
	mu    sync.Mutex
	cache map[string]*Entry // 这里缓存的是指针，方便取值操作
}

func NewMemo(f Func) *Memo {
	return &Memo{
		f:     f,
		mu:    sync.Mutex{},
		cache: make(map[string]*Entry),
	}
}

// 使用广播机制实现了重复抑制的多 goroutine 执行函数
func (m *Memo) Get(k string) (interface{}, error) {
	m.mu.Lock() // 先上锁
	entry := m.cache[k]
	if entry == nil { // 没有取到值就更新缓存
		entry = &Entry{ready: make(chan struct{})}
		m.cache[k] = entry
		m.mu.Unlock() // 手动释放锁

		entry.item.v, entry.item.err = m.f(k)
		fmt.Println("[request]:", k) // 可以看到对必应只请求了一次 // so cool
		close(entry.ready)         // 将数据准备就绪的消息广播给所有几乎同时执行 Get(k) 的其他 goroutine
	} else {
		// 解除锁并等待数据
		m.mu.Unlock()

		// 等待第一个执行 Get(k) 的 goroutine 来广播消息说它已经执行完毕，将数据存储到 v 中
		// 当 v.ready channel 被关闭时不再阻塞等待
		<-entry.ready
	}
	return entry.item.v, entry.item.err
}

type Entry struct {
	item  Item
	ready chan struct{} // 信号通道
}

type Item struct {
	v   interface{}
	err error
}

type Func func(k string) (interface{}, error)
