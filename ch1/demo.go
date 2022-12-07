package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"sync"
)

func main() {
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		m[i] = i
	}
	// map 中迭代顺序是随机，有意设计
	for k, v := range m {
		fmt.Print(k, v, " ")
	}

	// map 为引用拷贝
	demo1(m)              // 1:100
	fmt.Printf("%v\n", m) // 1:100

	x := 1                 // 值拷贝
	fmt.Printf("%v\n", &x) // 0xc4200161e0
	demo2(x)               // 0xc4200161e8

	// verb 可使用 %s 直接格式化 []byte
	b := []byte("demo")
	fmt.Printf("%s\n", b)

	cli := http.Client{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cli.Get("http://0.0.0.0:2333/")
		}()
	}

	wg.Wait()                                       // 把 mu.Lock() 和 mu.Unlock() 注释掉  // go run server2.go
	resp, _ := cli.Get("http://0.0.0.0:2333/count") // 循环请求 1000 次
	if (err == nil) {
	  counts, _ := ioutil.ReadAll(resp.Body) // 计数只有 140 次
	  fmt.Printf("counts: %s", counts) // counts 未加锁的情况下，更新共享变量极为不安全的    }  
    	}
}

func demo1(m map[int]int) {
	m[1] = 100
	fmt.Printf("%v\n", &m)
}

func demo2(i int) {
	fmt.Printf("%v\n", &i)
}
