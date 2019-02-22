package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := countWordsAndImages(url)
		if err != nil {
			return
		}
		fmt.Printf("%s\n [words]: %d  [images]: %d\n", url, words, images)
	}
}

// 实现 url 页面的图片和文字计数
func countWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("http get failed: %v", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("html parse failed: %v", err)
	}

	words, images = count(doc)
	return
}

func count(node *html.Node) (words, images int) {
	if node == nil {
		return
	}

	// 二叉树的层序遍历
	q := []*html.Node{node}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		switch cur.Type {
		case html.TextNode:
			txt := strings.TrimSpace(cur.Data)
			words += countWords(txt) // ok
		case html.ElementNode:
			if cur.Data == "img" {
				images++
			}
		}

		for cur := cur.FirstChild; cur != nil; cur = cur.NextSibling {
			q = append(q, cur)
		}
	}

	return
}

func countWords(sentence string) int {
	return len(strings.Fields(sentence))
}
