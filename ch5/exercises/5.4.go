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

	allVisit(nil, doc)
}

func allVisit(links []string, node *html.Node) {
	if node.Type == html.ElementNode {
		switch node.Data {
		case "a":
			for _, a := range node.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					links = append(links, a.Val)
				}
			}
		case "img", "script", "style":
			for _, a := range node.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}

	// 获取图片的链接
	if node.FirstChild != nil {
		allVisit(links, node.FirstChild)
	}
	if node.NextSibling != nil {
		allVisit(links, node.NextSibling)
	}
}
