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
		log.Fatalf("html parse fail: %v", err)
	}

	visit(nil, doc)
}

func visit(links []string, node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				fmt.Println(attr.Val)
				links = append(links, attr.Val)
			}
		}
	}

	// 直接将循环遍历分解即可
	if node.FirstChild != nil {
		visit(links, node.FirstChild)
	}
	if node.NextSibling != nil {
		visit(links, node.NextSibling)
	}
}
