package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

// 可使用 https://golang.google.cn 代替 https://golang.org 连接超时
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Printf("html parse fail:%v", err)
	}

	// 从 doc 节点开始向下解析
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" { // 链接节点
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}

	// 递归处理本页面链接指向的页面数据
	for cur := node.FirstChild; cur != nil; cur = cur.NextSibling {
		links = visit(links, cur) // 递归处理子节点
	}
	return links
}
