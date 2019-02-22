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

	var depth int
	startElem := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", 2*depth, "", n.Data)
			depth++
		}
	}

	endElem := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s</%s>\n", 2*depth, "", n.Data) //
			depth--                                       // 遇到结束标签自然深度 -1
		}
	}

	var foreachNode func(node *html.Node, pre, post func(node *html.Node))

	// 匿名函数内部对外部变量的修改是可见的
	foreachNode = func(node *html.Node, pre, post func(node *html.Node)) {
		if pre != nil {
			pre(node)
		}

		for cur := node.FirstChild; cur != nil; cur = cur.NextSibling {
			foreachNode(cur, pre, post) // 注意这里要使用本 main 函数的 foreachNode，就必须提前声明，不然它认为是其他 main 文件的
		}

		if post != nil {
			post(node)
		}
	}

	foreachNode(doc, startElem, endElem)
}
