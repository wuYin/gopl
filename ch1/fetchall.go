// fetchall 并发获取 URL 内容并报告它们的时间和大小
package main

import (
	"time"
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // 启动一个 goroutine
	}
	// 阻塞等待所有 goroutine 执行完毕
	for range os.Args[1:] {
		fmt.Println(<-ch) // 输出 ch 中的值
	}
	fmt.Printf("%.2fs spent\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// ch <- fmt.Sprintf("get error: %v", err)
		ch <- err.Error()
		return
	}
	n, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("copy error: %v", err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", secs, n, url)
}
