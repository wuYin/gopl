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
				workList <- crawl2(url) // 将子页面内的所有链接重新发送到 channel 中重新 range 处理
			}(url)
		}
	}
}

// 利用缓冲通道来实现信号量
// 类似于卡槽的意思，获得卡槽才能执行函数，从而控制并发的数量
var tokens = make(chan struct{}, 20)

func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}            // 获取令牌
	list, err := links.Extract(url) // 获得信号量，开始执行函数
	<-tokens                        // 释放令牌
	if err != nil {
		fmt.Println(err)
	}
	return list
}
