package main

import (
	"fmt"
	"gopl/ch8/links"
	"os"
)

func main() {
	worklist := make(chan []string)
	unseenlinks := make(chan string)

	go func() {
		worklist <- os.Args[1:]
	}()

	// 设定 20 个 goroutine 的并发协程池子
	for i := 0; i < 20; i++ {
		go func() {
			for url := range unseenlinks {
				subLinks := crawl3(url) // 抓取子页面链接
				go func() {
					worklist <- subLinks // 让池子里的协程不阻塞，就再开一个子协程将数据发送过去
				}()
			}
		}()
	}

	// 主 goroutine 负责去重处理
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if seen[link] {
				continue
			}
			seen[link] = true
			unseenlinks <- link
		}
	}
}

// 单一职责的 crawler
func crawl3(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		fmt.Println(err)
	}
	return list
}
