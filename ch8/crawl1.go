package main

import (
	"fmt"
	"gopl/ch8/links"
	"os"
)

// 并发地无止境地抓取数据
func main() {
	workList := make(chan []string)

	go func() {
		workList <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for list := range workList {
		for _, url := range list {
			if seen[url] {
				continue
			}
			seen[url] = true
			go func(url string) {
				workList <- crawl1(url) // 将子页面内的所有链接重新发送到 channel 中重新 range 处理
			}(url)
		}
	}
}

func crawl1(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		fmt.Println(err)
	}
	return list
}
