package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.gopl.io")
	if err != nil {
		log.Fatalf("read filed failed: %v", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("html parse fail: %v", err)
	}

	foreachNode(doc, startElem, endElem)
}

var depth int

func startElem(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", 2*depth, "", n.Data)
		depth++
	}
}

func endElem(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", 2*depth, "", n.Data) //
		depth--                                       // 遇到结束标签自然深度 -1
	}
}

// pre 在遍历当前层之前调用，post 在遍历后调用
func foreachNode(node *html.Node, pre, post func(node *html.Node)) {
	if pre != nil {
		pre(node)
	}

	for cur := node.FirstChild; cur != nil; cur = cur.NextSibling {
		foreachNode(cur, pre, post)
	}

	if post != nil {
		post(node)
	}
}
