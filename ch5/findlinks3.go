package main

import (
	"fmt"
)

func main() {
	bfsTraverse(crawl, []string{"https://golang.google.cn"})
}

// BFS 遍历借助哈希表实现
func bfsTraverse(f func(url string) []string, urls []string) {
	if f == nil {
		return
	}

	accessed := make(map[string]bool)
	for len(urls) > 0 {
		subURLs := urls
		urls = nil // 避免底层数组篡改

		for _, subURL := range subURLs {
			if accessed[subURL] {
				continue
			}
			accessed[subURL] = true
			urls = append(urls, f(subURL)...) // 将子页面的所有 URL 抓取到 urls 中准备下次任务
		}
	}
}

// 爬取 url 页面的所有链接并返回
func crawl(url string) (subURLs []string) {
	fmt.Println(url)
	subURLs, err := Extract(url)
	if err != nil {
		fmt.Println(err)
	}
	return
}
