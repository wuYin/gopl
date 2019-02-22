package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func main() {
	fmt.Println(title("http://www.gopl.io/"))
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()                                                   // 勿忘
		return fmt.Errorf("%s has content type %v, not text/html", url, ct) // 第一遍 Close()
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close() // 第二遍
	if err != nil {
		return err
	}

	visit := func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "title" {
			if node.FirstChild != nil {
				fmt.Println(node.FirstChild.Data) // The Go Programming Language
			}
		}
	}
	var foreachNode func(node *html.Node, pre, post func(node *html.Node))
	foreachNode = func(node *html.Node, pre, post func(node *html.Node)) {
		if pre != nil {
			pre(node)
		}

		for cur := node.FirstChild; cur != nil; cur = cur.NextSibling {
			foreachNode(cur, pre, post) // 提前 var 申明
		}

		if post != nil {
			post(node)
		}
	}
	foreachNode(doc, visit, nil)
	return nil
}
