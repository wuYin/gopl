package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func main() {
	f, err := os.Open("web.html")
	if err != nil {
		log.Fatalf("read file failed: %v", err)
	}

	doc, err := html.Parse(f)
	if err != nil {
		log.Fatalf("html parse failed: %v", err)
	}
	var count int
	traverse(&count, doc)
	fmt.Println("node count: ", count) // 100 // outline 验证
}

// 统计文档树内所有元素的数量
func traverse(count *int, node *html.Node) {
	if node.Type == html.ElementNode {
		*count++
	}

	if node.FirstChild != nil {
		traverse(count, node.FirstChild)
	}

	if node.NextSibling != nil {
		traverse(count, node.NextSibling)
	}
}
