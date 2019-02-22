package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("parse fail: %v", err)
	}
	outline(nil, doc)
}

// 输出树的结构
func outline(stack []string, node *html.Node) {
	if node.Type == html.ElementNode { // 节点类型
		stack := append(stack, node.Data) // 节点的值是标签名
		fmt.Println(stack)
	}
	for cur := node.FirstChild; cur != nil; cur = cur.NextSibling {
		outline(stack, cur) // 递归处理子节点
	}
}
