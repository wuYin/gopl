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

		go func(url string) {
			start := time.Now()
			v, err := m.Get(url) // cache map 存在数据竞态
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
		cache: make(map[string]Item),
	}
}

func (m *Memo) Get(k string) (interface{}, error) {
	item, ok := m.cache[k]
	if ok {
		return item.v, nil
	}

	item.v, item.err = m.f(k)
	m.cache[k] = item
	return item.v, item.err
}
