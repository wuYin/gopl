package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://golang.google.cn")
	if err != nil {
		log.Fatalf("read file failed: %v", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("html parse fail: %v", err)
	}

	id := "about"
	node := ElementByID(doc, id)
	if node == nil {
		fmt.Printf("id %q element not found\n", id)
		return
	}
	fmt.Println(node.Data) // div // ok
}

// 实现类似 ElementID 的查找
func ElementByID(n *html.Node, id string) *html.Node {
	pre := func(n *html.Node) (stop bool) {
		if n == nil {
			return false
		}
		if n.Type != html.ElementNode {
			return false
		}

		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
		return false
	}

	return foreachNode(n, pre, nil)
}

// 找到第一个元素就停止查找，使用 BFS 即可
func foreachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	q := []*html.Node{n}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if pre != nil && pre(cur) { // cur 就是要找的 id node
			return cur
		}

		for next := cur.FirstChild; next != nil; next = next.NextSibling {
			q = append(q, next)
		}

		if post != nil && post(cur) { // 找到了
			return cur
		}
	}

	return nil
}
