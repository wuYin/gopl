package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("web.html")
	if err != nil {
		log.Fatalf("read file failed: %v", err)
	}

	doc, err := html.Parse(f)
	if err != nil {
		log.Fatalf("html parse fail: %v", err)
	}

	getText(doc)
}

// 获取文档中所有文本
func getText(node *html.Node) {
	if node.Type == html.TextNode && node.Data != "script" && node.Data != "style" {
		if len(strings.TrimSpace(node.Data)) > 0 {
			fmt.Printf("%q\n", strings.TrimSpace(node.Data))
		}
	}

	if node.FirstChild != nil {
		getText(node.FirstChild)
	}
	if node.NextSibling != nil {
		getText(node.NextSibling)
	}
}
