package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks2(url)
		if err != nil {
			fmt.Printf("find lilks2 failed: %v %s\n", err, url)
			continue
		}

		for _, l := range links {
			fmt.Println(l)
		}
	}
}

// 手动处理错误
func findLinks2(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return visit2(nil, doc), nil
}

func visit2(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for cur := n.FirstChild; cur != nil; cur = cur.NextSibling {
		links = visit2(links, cur) // 注意对 links 重新赋值，已经在内部 append
	}
	return links
}
