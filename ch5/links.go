package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
)

func main() {
	links, err := Extract("https://golang.google.cn")
	if err != nil {
		log.Fatalf("extract failed: %v", err)
	}
	fmt.Println(strings.Join(links, "\n")) // bingo
}

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid %s status: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	count := 0
	visit := func(n *html.Node) {
		if n == nil {
			return
		}
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}

				link, err := resp.Request.URL.Parse(a.Val) // 在指定上下文内补全链接为绝对路径 URL
				if err != nil {
					continue
				}
				links = append(links, link.String()) // 匿名函数内部对变量的操作对外部是有影响的
				count++                              // count 会自增，变量没有以形参的形式传入函数，函数本身就会保存这个变量的状态（值变化）
			}
		}
	}
	foreach(doc, visit, nil)    // 处理 links
	fmt.Println("count", count) // 17

	return links, nil
}

func foreach(node *html.Node, pre, post func(node *html.Node)) {
	if pre != nil {
		pre(node)
	}

	for cur := node.FirstChild; cur != nil; cur = cur.NextSibling {
		foreach(cur, pre, post)
	}

	if post != nil {
		post(node)
	}
}
