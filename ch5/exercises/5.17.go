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
		log.Fatalf("read file failed: %v", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("html parse fail: %v", err)
	}

	names := []string{"img"}
	nodes := ElementByTagName(doc, names...)
	if nodes == nil {
		fmt.Printf("names %q element not found\n", names)
		return
	}

	fmt.Printf("%v\n[count]: %d", names, len(nodes)) // 1 张图片
}

// 实现类似 ElementID 的查找
func ElementByTagName(n *html.Node, names ...string) []*html.Node {
	tags := make(map[string]bool)
	for _, name := range names {
		tags[name] = true
	}

	var nodes []*html.Node // 放心地放到闭包函数内部
	pre := func(n *html.Node) (stop bool) {
		if n == nil {
			return false
		}
		if n.Type != html.ElementNode {
			return false
		}

		if _, ok := tags[n.Data]; ok { // 要找的节点
			nodes = append(nodes, n) // bingo
			return true
		}
		return false
	}

	foreach(n, pre, nil) // 在 pre 内部尽情地修改 nodes 反正对 nodes 的修改是可见的
	return nodes
}

func foreach(node *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	q := []*html.Node{node}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if pre != nil && pre(cur) { // got one
			return cur
		}

		for next := cur.FirstChild; next != nil; next = next.NextSibling {
			q = append(q, next)
		}

		if post != nil && post(cur) { // ok
			return cur
		}
	}

	return nil
}
